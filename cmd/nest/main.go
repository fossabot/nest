package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	consulapi "github.com/hashicorp/consul/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vx-labs/nest/nest"
	"github.com/vx-labs/nest/nest/api"
	"github.com/vx-labs/nest/nest/async"
	"github.com/vx-labs/nest/nest/fsm"
	"github.com/vx-labs/wasp/cluster"
	"github.com/vx-labs/wasp/cluster/raft"
	"github.com/vx-labs/wasp/rpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func localPrivateHost() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, v := range ifaces {
		if v.Flags&net.FlagLoopback != net.FlagLoopback && v.Flags&net.FlagUp == net.FlagUp {
			h := v.HardwareAddr.String()
			if len(h) == 0 {
				continue
			} else {
				addresses, _ := v.Addrs()
				if len(addresses) > 0 {
					ip := addresses[0]
					if ipnet, ok := ip.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							return ipnet.IP.String()
						}
					}
				}
			}
		}
	}
	panic("could not find a valid network interface")
}
func findPeers(name, tag string, minimumCount int) ([]string, error) {
	config := consulapi.DefaultConfig()
	config.HttpClient = http.DefaultClient
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	var idx uint64
	for {
		services, meta, err := client.Catalog().Service(name, tag, &consulapi.QueryOptions{
			WaitIndex: idx,
			WaitTime:  10 * time.Second,
		})
		if err != nil {
			return nil, err
		}
		idx = meta.LastIndex
		if len(services) < minimumCount {
			continue
		}
		out := make([]string, len(services))
		for idx := range services {
			out[idx] = fmt.Sprintf("%s:%d", services[idx].ServiceAddress, services[idx].ServicePort)
		}
		return out, nil
	}
}

func main() {
	config := viper.New()
	config.SetEnvPrefix("NEST")
	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	config.AutomaticEnv()
	cmd := cobra.Command{
		Use: "nest",
		PreRun: func(cmd *cobra.Command, _ []string) {
			config.BindPFlags(cmd.Flags())
			if !cmd.Flags().Changed("serf-advertized-port") {
				config.Set("serf-advertized-port", config.Get("serf-port"))
			}
			if !cmd.Flags().Changed("raft-advertized-port") {
				config.Set("raft-advertized-port", config.Get("raft-port"))
			}

		},
		Run: func(cmd *cobra.Command, _ []string) {
			ctx, cancel := context.WithCancel(context.Background())
			ctx = nest.StoreLogger(ctx, getLogger(config))
			err := os.MkdirAll(config.GetString("data-dir"), 0700)
			if err != nil {
				nest.L(ctx).Fatal("failed to create data directory", zap.Error(err))
			}
			id, err := loadID(config.GetString("data-dir"))
			if err != nil {
				nest.L(ctx).Fatal("failed to get node ID", zap.Error(err))
			}
			ctx = nest.AddFields(ctx, zap.String("hex_node_id", fmt.Sprintf("%x", id)))
			if config.GetBool("pprof") {
				address := fmt.Sprintf("%s:%d", config.GetString("pprof-address"), config.GetInt("pprof-port"))
				go func() {
					mux := http.NewServeMux()
					mux.HandleFunc("/debug/pprof/", pprof.Index)
					mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
					mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
					mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
					mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
					panic(http.ListenAndServe(address, mux))
				}()
				nest.L(ctx).Info("started pprof", zap.String("pprof_url", fmt.Sprintf("http://%s/", address)))
			}
			healthServer := health.NewServer()
			healthServer.Resume()
			wg := sync.WaitGroup{}
			cancelCh := make(chan struct{})
			commandsCh := make(chan raft.Command)
			recordsCh := make(chan *api.Record, 20)
			if config.GetString("rpc-tls-certificate-file") == "" || config.GetString("rpc-tls-private-key-file") == "" {
				nest.L(ctx).Warn("TLS certificate or private key not provided. GRPC transport security will use a self-signed generated certificate.")
			}
			server := rpc.Server(rpc.ServerConfig{
				VerifyClientCert:            config.GetBool("mtls"),
				TLSCertificateAuthorityPath: config.GetString("rpc-tls-certificate-authority-file"),
				TLSCertificatePath:          config.GetString("rpc-tls-certificate-file"),
				TLSPrivateKeyPath:           config.GetString("rpc-tls-private-key-file"),
			})
			healthpb.RegisterHealthServer(server, healthServer)
			rpcDialer := rpc.GRPCDialer(rpc.ClientConfig{
				InsecureSkipVerify:          config.GetBool("insecure"),
				TLSCertificatePath:          config.GetString("rpc-tls-certificate-file"),
				TLSPrivateKeyPath:           config.GetString("rpc-tls-private-key-file"),
				TLSCertificateAuthorityPath: config.GetString("rpc-tls-certificate-authority-file"),
			})
			clusterListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.GetInt("raft-port")))
			if err != nil {
				nest.L(ctx).Fatal("cluster listener failed to start", zap.Error(err))
			}

			joinList := config.GetStringSlice("join-node")
			if config.GetBool("consul-join") {
				discoveryStarted := time.Now()
				consulJoinList, err := findPeers(
					config.GetString("consul-service-name"), config.GetString("consul-service-tag"),
					config.GetInt("raft-bootstrap-expect"))
				if err != nil {
					nest.L(ctx).Fatal("failed to find other peers on Consul", zap.Error(err))
				}
				nest.L(ctx).Debug("discovered nodes using Consul",
					zap.Duration("consul_discovery_duration", time.Since(discoveryStarted)), zap.Int("node_count", len(consulJoinList)))
				joinList = append(joinList, consulJoinList...)
			}
			messageLog, err := nest.NewMessageLog(ctx, id, config.GetString("data-dir"))
			if err != nil {
				nest.L(ctx).Fatal("failed to load message log", zap.Error(err))
			}

			stateMachine := fsm.NewFSM(id, messageLog, commandsCh)
			clusterNode := cluster.NewNode(cluster.NodeConfig{
				ID:            id,
				ServiceName:   "wasp",
				DataDirectory: config.GetString("data-dir"),
				GossipConfig: cluster.GossipConfig{
					JoinList: joinList,
					Network: cluster.NetworkConfig{
						AdvertizedHost: config.GetString("serf-advertized-address"),
						AdvertizedPort: config.GetInt("serf-advertized-port"),
						ListeningPort:  config.GetInt("serf-port"),
					},
				},
				RaftConfig: cluster.RaftConfig{
					ExpectedNodeCount: config.GetInt("raft-bootstrap-expect"),
					Network: cluster.NetworkConfig{
						AdvertizedHost: config.GetString("raft-advertized-address"),
						AdvertizedPort: config.GetInt("raft-advertized-port"),
						ListeningPort:  config.GetInt("raft-port"),
					},
					AppliedIndex: messageLog.CurrentStateOffset(),
				},
				GetStateSnapshot: func() ([]byte, error) {
					nest.L(ctx).Debug("snapshoting message log")
					return messageLog.Snapshot()
				},
			},
				rpcDialer, server, nest.L(ctx))
			snapshotter := <-clusterNode.Snapshotter()

			async.Run(ctx, &wg, func(ctx context.Context) {
				defer nest.L(ctx).Debug("cluster node stopped")
				clusterNode.Run(ctx)
			})
			waspReceiver := nest.NewWaspReceiver(stateMachine)
			messagesServer := nest.NewServer(stateMachine, messageLog)
			messagesServer.Serve(server)
			waspReceiver.Serve(server)
			async.Run(ctx, &wg, func(ctx context.Context) {
				defer nest.L(ctx).Info("cluster listener stopped")

				err := server.Serve(clusterListener)
				if err != nil {
					nest.L(ctx).Fatal("cluster listener crashed", zap.Error(err))
				}
			})
			async.Run(ctx, &wg, func(ctx context.Context) {
				defer nest.L(ctx).Info("records publisher stopped")
				for {
					select {
					case <-ctx.Done():
						return
					case event := <-recordsCh:
						err := stateMachine.PutRecords(ctx, []*api.Record{event})
						if err != nil {
							nest.L(ctx).Warn("record put failed", zap.Error(err))
						}
					}
				}
			})
			async.Run(ctx, &wg, func(ctx context.Context) {
				defer nest.L(ctx).Info("command publisher stopped")
				for {
					select {
					case <-ctx.Done():
						return
					case event := <-commandsCh:
						err := clusterNode.Apply(event.Ctx, event.Payload)
						select {
						case <-ctx.Done():
						case <-event.Ctx.Done():
						case event.ErrCh <- err:
						}
						close(event.ErrCh)
					}
				}
			})
			async.Run(ctx, &wg, func(ctx context.Context) {
				defer nest.L(ctx).Info("command processor stopped")
				for {
					select {
					case <-ctx.Done():
						return
					case event := <-clusterNode.Commits():
						if event.Payload == nil {
							snapshot, err := snapshotter.Load()
							if err != nil {
								nest.L(ctx).Error("failed to load", zap.Error(err))
							}
							err = messageLog.Restore(ctx, snapshot.Data, clusterNode.Call)
							if err != nil {
								nest.L(ctx).Debug("failed to load state snapshot", zap.Error(err))
							}
						} else {
							stateMachine.Apply(event.Index, event.Payload)
						}
					}
				}
			})
			<-clusterNode.Ready()
			if brokerURL := config.GetString("mqtt-collector-broker-url"); brokerURL != "" {
				mqttCollector, err := nest.MQTTCollector(brokerURL,
					config.GetString("mqtt-collector-broker-username"),
					config.GetString("mqtt-collector-broker-password"))
				if err != nil {
					nest.L(ctx).Error("failed to start mqtt collector", zap.Error(err))
					close(cancelCh)
				} else {
					async.Run(ctx, &wg, func(ctx context.Context) {
						defer nest.L(ctx).Info("mqtt collector stopped")
						err := mqttCollector.Run(ctx, recordsCh)
						if err != nil {
							nest.L(ctx).Error("mqtt collector crashed", zap.Error(err))
						}
					})
				}
			}
			//go stats.ListenAndServe(config.GetInt("metrics-port"))

			healthServer.SetServingStatus("mqtt", healthpb.HealthCheckResponse_SERVING)
			healthServer.SetServingStatus("node", healthpb.HealthCheckResponse_SERVING)
			healthServer.SetServingStatus("rpc", healthpb.HealthCheckResponse_SERVING)
			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			select {
			case <-sigc:
			case <-cancelCh:
			}
			nest.L(ctx).Info("nest shutdown initiated")
			err = stateMachine.Shutdown(ctx)
			if err != nil {
				nest.L(ctx).Error("failed to stop state machine", zap.Error(err))
			} else {
				nest.L(ctx).Debug("state machine stopped")
			}
			err = clusterNode.Shutdown()
			if err != nil {
				nest.L(ctx).Error("failed to leave cluster", zap.Error(err))
			} else {
				nest.L(ctx).Debug("cluster left")
			}
			healthServer.Shutdown()
			nest.L(ctx).Debug("health server stopped")
			go func() {
				<-time.After(1 * time.Second)
				server.Stop()
			}()
			server.GracefulStop()
			nest.L(ctx).Debug("rpc server stopped")
			clusterListener.Close()
			nest.L(ctx).Debug("rpc listener stopped")
			cancel()
			wg.Wait()
			nest.L(ctx).Debug("asynchronous operations stopped")
			messageLog.Close()
			nest.L(ctx).Debug("message log closed")
			nest.L(ctx).Info("nest successfully stopped")
		},
	}
	defaultIP := localPrivateHost()
	cmd.Flags().Bool("pprof", false, "Start pprof endpoint.")
	cmd.Flags().Int("pprof-port", 8080, "Profiling (pprof) port.")
	cmd.Flags().String("pprof-address", "127.0.0.1", "Profiling (pprof) port.")
	cmd.Flags().Bool("debug", false, "Use a fancy logger and increase logging level.")
	cmd.Flags().Bool("mtls", false, "Enforce GRPC service-side TLS certificates validation for client connections.")
	cmd.Flags().Bool("insecure", false, "Disable GRPC client-side TLS validation.")
	cmd.Flags().Bool("consul-join", false, "Use Hashicorp Consul to find other gossip members. Nest won't handle service registration in Consul, you must do it before running Nest.")
	cmd.Flags().String("consul-service-name", "Nest", "Consul auto-join service name.")
	cmd.Flags().String("consul-service-tag", "gossip", "Consul auto-join service tag.")

	cmd.Flags().Int("metrics-port", 0, "Start Prometheus HTTP metrics server on this port.")
	cmd.Flags().Int("serf-port", 2799, "Membership (Serf) port.")
	cmd.Flags().Int("raft-port", 2899, "Clustering (Raft) port.")
	cmd.Flags().String("serf-advertized-address", defaultIP, "Advertize this adress to other gossip members.")
	cmd.Flags().String("raft-advertized-address", defaultIP, "Advertize this adress to other raft nodes.")
	cmd.Flags().Int("serf-advertized-port", 2799, "Advertize this port to other gossip members.")
	cmd.Flags().Int("raft-advertized-port", 2899, "Advertize this port to other raft nodes.")
	cmd.Flags().StringSliceP("join-node", "j", nil, "Join theses nodes to form a cluster.")
	cmd.Flags().StringP("data-dir", "d", "/tmp/nest", "Nest persistent data location.")

	cmd.Flags().IntP("raft-bootstrap-expect", "n", 1, "Nest will wait for this number of nodes to be available before bootstraping a cluster.")

	cmd.Flags().String("rpc-tls-certificate-authority-file", "", "x509 certificate authority used by RPC Server.")
	cmd.Flags().String("rpc-tls-certificate-file", "", "x509 certificate used by RPC Server.")
	cmd.Flags().String("rpc-tls-private-key-file", "", "Private key used by RPC Server.")

	cmd.Flags().String("mqtt-collector-broker-url", "", "Collect messages from this MQTT Broker.")
	cmd.Flags().String("mqtt-collector-broker-username", "vx:psk", "Username to authenticate against MQTT Broker.")
	cmd.Flags().String("mqtt-collector-broker-password", "", "Password to authenticate against MQTT Broker.")

	cmd.AddCommand(TLSHelper(config))
	cmd.Execute()
}
