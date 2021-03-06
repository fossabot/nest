package nest

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/vx-labs/nest/nest/fsm"
	"github.com/vx-labs/wasp/async"
	"github.com/vx-labs/wasp/cluster"
	"github.com/vx-labs/wasp/cluster/raft"
	"go.uber.org/zap"
)

type Shard interface {
	PutRecords(ctx context.Context, records [][]byte) error
	LookupTimestamp(ts uint64) uint64
	Shutdown(ctx context.Context) error
	Consume(f func(r io.ReadSeeker) error) error
	Dump(w io.Writer, fromOffset uint64) error
	Load(r io.Reader) error
	Offset() uint64
	Ready() <-chan struct{}
	Stop() error
}

type shard struct {
	ctx        context.Context
	cancel     context.CancelFunc
	operations async.Operations
	node       cluster.Node
	fsm        *fsm.FSM
	recorder   Recorder
}

func (s *shard) Ready() <-chan struct{} {
	return s.node.Ready()
}
func (s *shard) LookupTimestamp(ts uint64) uint64 {
	return s.recorder.LookupTimestamp(ts)
}
func (s *shard) Stop() error {
	err := s.node.Shutdown()
	if err != nil {
		return err
	}
	s.cancel()
	s.operations.Wait()
	return s.recorder.Close()
}
func (s *shard) Shutdown(ctx context.Context) error {
	return s.fsm.Shutdown(ctx)
}
func (s *shard) PutRecords(ctx context.Context, records [][]byte) error {
	return s.fsm.PutRecords(ctx, records)
}
func (s *shard) Consume(f func(r io.ReadSeeker) error) error {
	return s.recorder.Consume(f)
}
func (s *shard) Dump(w io.Writer, fromOffset uint64) error {
	return s.recorder.Dump(w, fromOffset)
}
func (s *shard) Load(r io.Reader) error {
	return s.recorder.Load(r)
}
func (s *shard) Offset() uint64 {
	return s.recorder.Offset()
}

func newShard(id uint64, stream string, shardID uint64, datadir string, clusterMultiNode cluster.MultiNode, raftConfig cluster.RaftConfig, logger *zap.Logger) (*shard, error) {
	datadir = path.Join(datadir, stream, fmt.Sprintf("%d", shardID))
	commandsCh := make(chan raft.Command)

	recorder, err := NewRecorder(id, stream, shardID, datadir, logger)
	if err != nil {
		return nil, err
	}
	raftConfig.AppliedIndex = recorder.CurrentStateOffset()
	raftConfig.GetStateSnapshot = recorder.Snapshot
	node := clusterMultiNode.Node(fmt.Sprintf("%s-%d", stream, shardID), raftConfig)

	ctx, cancel := context.WithCancel(context.Background())
	snapshotter := <-node.Snapshotter()
	stateMachine := fsm.NewFSM(id, recorder, commandsCh)
	ctx = StoreLogger(ctx, logger)
	operations := async.NewOperations(ctx, logger)
	operations.Run("recorder command processor", func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-node.Commits():
				if event.Payload == nil {
					snapshot, err := snapshotter.Load()
					if err != nil {
						logger.Error("failed to load", zap.Error(err))
					}
					err = recorder.Restore(ctx, snapshot.Data, node.Call)
					if err != nil {
						logger.Debug("failed to load state snapshot", zap.Error(err))
					}
				} else {
					stateMachine.Apply(event.Index, event.Payload)
				}
			}
		}
	})
	operations.Run("command publisher", func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-commandsCh:
				err := node.Apply(event.Ctx, event.Payload)
				select {
				case <-ctx.Done():
				case <-event.Ctx.Done():
				case event.ErrCh <- err:
				}
				close(event.ErrCh)
			}
		}
	})
	operations.Run("cluster node", func(ctx context.Context) {
		node.Run(ctx)
	})
	return &shard{
		ctx:        ctx,
		cancel:     cancel,
		fsm:        stateMachine,
		node:       node,
		recorder:   recorder,
		operations: operations,
	}, nil
}

type Controller interface {
	WaitReady(ctx context.Context)
	Shards() []Shard
	Shutdown(ctx context.Context) error
	Stop() error
}

type controller struct {
	shards []Shard
}

func NewController(ctx context.Context, id uint64, stream string, shardCount int, datadir string, clusterMultiNode cluster.MultiNode, server StreamsServer, raftConfig cluster.RaftConfig, logger *zap.Logger) (Controller, error) {
	outShards := make([]Shard, shardCount)
	var err error
	for idx := range outShards {
		outShards[idx], err = newShard(id, stream, uint64(idx), datadir, clusterMultiNode, raftConfig, logger)
		if err != nil {
			return nil, err
		}
	}
	server.RegisterShards(stream, outShards)
	return &controller{
		shards: outShards,
	}, nil
}

func (c *controller) Shards() []Shard { return c.shards }
func (c *controller) WaitReady(ctx context.Context) {
	for _, shard := range c.shards {
		select {
		case <-shard.Ready():
		case <-ctx.Done():
			return
		}
	}
}
func (c *controller) Shutdown(ctx context.Context) error {
	for _, shard := range c.shards {
		err := shard.Shutdown(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *controller) Stop() error {
	for idx := range c.shards {
		shard := c.shards[idx]
		err := shard.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
