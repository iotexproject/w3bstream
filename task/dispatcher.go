package task

import (
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/pkg/errors"
)

type Dispatcher struct {
	ps         *p2p.PubSubs
	pg         *persistence.Postgres
	projectIDs []uint64
}

// will block caller
func (d *Dispatcher) Dispatch() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, projectID := range d.projectIDs {
			t, err := d.pg.Fetch(projectID)
			if err != nil {
				slog.Error("get task failed", "error", err, "projectID", projectID)
				continue
			}
			if t == nil {
				continue
			}
			if err := d.ps.Publish(projectID, &p2p.Data{
				Task: t,
			}); err != nil {
				slog.Error("publish data failed", "error", err, "projectID", projectID)
			}
		}
	}
}

func (d *Dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	l := data.TaskStateLog
	if err := d.pg.UpdateState(l.TaskID, l.State, l.Comment, l.CreatedAt); err != nil {
		slog.Error("update task state failed", "error", err, "taskID", l.TaskID)
	}
}

func NewDispatcher(projectIDs []uint64, pg *persistence.Postgres, bootNodeMultiaddr string, iotexChainID int) (*Dispatcher, error) {
	d := &Dispatcher{
		projectIDs: projectIDs,
		pg:         pg,
	}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.ps = ps

	for _, id := range projectIDs {
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "add project %d pubsub failed", id)
		}
	}
	return d, nil
}
