package processor

import (
	"crypto/ecdsa"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task"
)

type HandleTask func(task *task.Task, vmTypeID uint64, code string, expParam string) ([]byte, error)
type Project func(projectID uint64) (*project.Project, error)
type RetrieveTask func(projectID uint64, taskID common.Hash) (*task.Task, error)

type DB interface {
	UnprocessedTask() (uint64, common.Hash, error)
	ProcessTask(uint64, common.Hash) error
}

type Processor struct {
	db          DB
	retrieve    RetrieveTask
	handle      HandleTask
	project     Project
	prv         *ecdsa.PrivateKey
	waitingTime time.Duration
}

func (r *Processor) process(projectID uint64, taskID common.Hash) error {
	t, err := r.retrieve(projectID, taskID)
	if err != nil {
		return err
	}
	p, err := r.project(t.ProjectID)
	if err != nil {
		return err
	}
	c, err := p.Config(t.ProjectVersion)
	if err != nil {
		return err
	}
	slog.Debug("get a new task", "project_id", t.ProjectID, "task_id", t.ID)

	proof, err := r.handle(t, c.VMTypeID, c.Code, c.CodeExpParam)
	if err != nil {
		return err
	}
	// TODO write proof to router
}

func (r *Processor) run() {
	for {
		projectID, taskID, err := r.db.UnprocessedTask()
		if err != nil {
			slog.Error("failed to get unprocessed task", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		if projectID == 0 {
			time.Sleep(r.waitingTime)
			continue
		}
		if err := r.process(projectID, taskID); err != nil {
			slog.Error("failed to process task", "error", err)
			continue
		}
		if err := r.db.ProcessTask(projectID, taskID); err != nil {
			slog.Error("failed to process db task", "error", err)
		}
	}
}

// func (r *Processor) reportSuccess(t *task.Task, state task.State, result []byte, signature string, topic *pubsub.Topic) {
// 	d, err := json.Marshal(&p2p.Data{
// 		TaskStateLog: &task.StateLog{
// 			TaskID:    t.ID,
// 			ProjectID: t.ProjectID,
// 			State:     state,
// 			Result:    result,
// 			Signature: signature,
// 			ProverID:  r.proverID,
// 			CreatedAt: time.Now(),
// 		},
// 	})
// 	if err != nil {
// 		slog.Error("failed to marshal p2p task state log data to json", "error", err, "task_id", t.ID)
// 		return
// 	}
// 	if err := topic.Publish(context.Background(), d); err != nil {
// 		slog.Error("failed to publish task state log data to p2p network", "error", err, "task_id", t.ID)
// 	}
// }

func NewProcessor(handle HandleTask, project Project, db DB, retrieve RetrieveTask, prv *ecdsa.PrivateKey) *Processor {
	p := &Processor{
		db:          db,
		retrieve:    retrieve,
		handle:      handle,
		project:     project,
		prv:         prv,
		waitingTime: 3 * time.Second,
	}
	go p.run()
	return p
}
