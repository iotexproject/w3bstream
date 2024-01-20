package task

import (
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
)

type Dispatcher struct {
	ps  *p2p.PubSubs
	pg  *persistence.Postgres
	mgr *project.Manager
}

// will block caller
func (d *Dispatcher) Dispatch() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		t, err := d.pg.Fetch()
		if err != nil {
			slog.Error("get task failed", "error", err)
			continue
		}
		if t == nil {
			continue
		}

		projectID := t.Messages[0].ProjectID
		if err := d.ps.Add(projectID); err != nil {
			slog.Error("add project pubsub failed", "error", err, "projectID", projectID)
			continue
		}

		slog.Debug("dispatch project task", "projectID", projectID, "taskID", t.ID)
		if err := d.ps.Publish(projectID, &p2p.Data{
			Task: t,
		}); err != nil {
			slog.Error("publish data failed", "error", err, "projectID", projectID)
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

func NewDispatcher(pg *persistence.Postgres, bootNodeMultiaddr string, iotexChainID int) (*Dispatcher, error) {
	d := &Dispatcher{
		pg: pg,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.ps = ps

	return d, nil
}
