package aggregator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/project"
	apidb "github.com/iotexproject/w3bstream/service/apinode/db"
	"github.com/iotexproject/w3bstream/service/sequencer/api"
	"github.com/pkg/errors"
)

func Run(projectManager *project.Manager, db *apidb.DB, sequencerAddr string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ts, err := db.FetchAllTask()
		if err != nil {
			slog.Error("failed to fetch all tasks", "error", err)
			continue
		}
		if len(ts) == 0 {
			continue
		}

		taskMap := make(map[string][]*apidb.Task)
		for i := range ts {
			k := ts[i].ProjectID + "_" + ts[i].DevicePubKey
			taskMap[k] = append(taskMap[k], ts[i])
		}

		for taskMapKey, tasks := range taskMap {
			pid, ok := new(big.Int).SetString(tasks[0].ProjectID, 10)
			if !ok {
				slog.Error("failed to decode project id string", "project_string", tasks[0].ProjectID)
				continue
			}
			p, err := projectManager.Project(pid)
			if err != nil {
				slog.Error("failed to get project", "error", err, "project_id", pid.String())
				continue
			}
			// TODO support project config
			cfg, err := p.DefaultConfig()
			if err != nil {
				slog.Error("failed to get project config", "error", err, "project_id", pid.String())
				continue
			}
			if cfg.ProofType == "movement" {
				now := time.Now()
				hoursAgo24 := now.Add(-24 * time.Hour) // TODO move to project config
				count, err := db.CountSettledTask(tasks[0].ProjectID, tasks[0].DevicePubKey, hoursAgo24, now)
				if err != nil {
					slog.Error("failed to count settled task", "error", err, "project_id", pid.String())
					continue
				}
				if count >= 5 { // TODO move to project config
					delete(taskMap, taskMapKey)
				} else {
					prevTaskID := tasks[0].TaskID
					tasks[len(tasks)-1].PrevTaskID = prevTaskID
				}
			}
		}

		if err := dumpTasks(db, ts); err != nil {
			slog.Error("failed to dump tasks", "error", err)
			continue
		}

		for _, tasks := range taskMap {
			lastTask := tasks[len(tasks)-1]
			if err := notify(sequencerAddr, common.HexToHash(lastTask.TaskID)); err != nil {
				slog.Error("failed to notify sequencer", "error", err)
				continue
			}
		}
	}
}

func dumpTasks(db *apidb.DB, ts []*apidb.Task) error {
	// add tasks to remote
	if err := db.CreateTasks(ts); err != nil {
		slog.Error("failed to create tasks", "error", err)
		return err
	}
	// remove tasks from local
	if err := db.DeleteTasks(ts); err != nil {
		slog.Error("failed to delete tasks at local", "error", err)
		return err
	}
	return nil
}

func notify(sequencerAddr string, taskID common.Hash) error {
	reqSequencer := &api.CreateTaskReq{TaskID: taskID}
	reqSequencerJ, err := json.Marshal(reqSequencer)
	if err != nil {
		return errors.Wrap(err, "failed to marshal sequencer request")
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/task", sequencerAddr), "application/json", bytes.NewBuffer(reqSequencerJ))
	if err != nil {
		return errors.Wrap(err, "failed to call sequencer service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			err = errors.New(string(body))
		}
		return errors.Wrap(err, "failed to call sequencer service")
	}
	return nil
}
