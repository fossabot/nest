package nest

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/vx-labs/nest/nest/api"
	"go.uber.org/zap"
)

type Collector interface {
	Run(context.Context, chan *api.Record) error
}

type mqttCollector struct {
	client MQTT.Client
}

func (m *mqttCollector) Run(ctx context.Context, ch chan *api.Record) error {
	token := m.client.Subscribe("#", 2, func(client MQTT.Client, msg MQTT.Message) {
		if msg.Retained() {
			return
		}
		L(ctx).Debug("mqtt message collected",
			zap.String("mqtt_topic", msg.Topic()), zap.String("mqtt_payload", string(msg.Payload())))
		ch <- &api.Record{
			Timestamp: time.Now().UnixNano(),
			Payload:   msg.Payload(),
			Topic:     []byte(msg.Topic()),
		}
	})
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	select {
	case <-ctx.Done():
		m.client.Disconnect(500)
		return nil
	}
}

func MQTTCollector(broker, username, password string) (Collector, error) {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.Username = username
	opts.Password = password
	brokerURL, err := url.Parse(broker)
	if err != nil {
		return nil, err
	}
	if brokerURL.Scheme == "tls" {
		host, _, _ := net.SplitHostPort(brokerURL.Host)
		opts.TLSConfig = &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			},
			ServerName: host,
		}
	}
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttCollector{
		client: c,
	}, nil
}
