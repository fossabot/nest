package fsm

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/vx-labs/wasp/cluster/raft"
)

type State interface {
	Append(uint64, []uint64, [][]byte) error
	SetCurrentStateOffset(v uint64)
}

func decode(payload []byte) ([]*StateTransition, error) {
	format := StateTransitionSet{}
	err := proto.Unmarshal(payload, &format)
	if err != nil {
		return nil, err
	}
	return format.Events, nil
}
func encode(events ...*StateTransition) ([]byte, error) {
	format := StateTransitionSet{
		Events: events,
	}
	return proto.Marshal(&format)
}

func NewFSM(id uint64, state State, commandsCh chan raft.Command) *FSM {
	return &FSM{id: id, state: state, commandsCh: commandsCh}
}

type FSM struct {
	id         uint64
	state      State
	commandsCh chan raft.Command
}

func (f *FSM) commit(ctx context.Context, payload []byte) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	out := make(chan error)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case f.commandsCh <- raft.Command{Ctx: ctx, ErrCh: out, Payload: payload}:
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-out:
			return err
		}
	}
}
func (f *FSM) Shutdown(ctx context.Context) error {
	payload, err := encode(&StateTransition{Event: &StateTransition_PeerLost{
		PeerLost: &PeerLost{
			Peer: f.id,
		},
	}})
	if err != nil {
		return err
	}
	return f.commit(ctx, payload)
}
func (f *FSM) PutRecords(ctx context.Context, records [][]byte) error {
	stamps := make([]uint64, len(records))
	for idx, record := range records {
		stamps[idx] = uint64(time.Now().UnixNano())
		if len(record) > 20*1000*1000 {
			return errors.New("payload size exceeded")
		}
	}
	payload, err := encode(&StateTransition{Event: &StateTransition_RecordsPut{
		RecordsPut: &RecordsPut{
			Records:    records,
			Timestamps: stamps,
		},
	}})
	if err != nil {
		return err
	}
	return f.commit(ctx, payload)
}

func (f *FSM) Apply(index uint64, b []byte) error {
	events, err := decode(b)
	if err != nil {
		return err
	}
	for _, event := range events {
		switch event := event.GetEvent().(type) {
		case *StateTransition_RecordsPut:
			err = f.state.Append(index, event.RecordsPut.Timestamps, event.RecordsPut.Records)
		default:
			f.state.SetCurrentStateOffset(index)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
