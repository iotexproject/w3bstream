package dispatcher

import (
	"log/slog"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/p2p"
	contractpkg "github.com/iotexproject/w3bstream/persistence/contract"
	"github.com/iotexproject/w3bstream/task"
)

type projectDispatcher struct {
	contract             Contract // optional, nil means local model
	vmTypeID             uint64
	window               *window
	waitInterval         time.Duration
	startTaskID          uint64
	projectID            uint64
	datasource           datasource.Datasource
	pubSubs              *p2p.PubSubs
	datasourcePubKey     []byte
	requiredProverAmount *atomic.Uint64
	paused               *atomic.Bool
	idle                 *atomic.Bool
}

func (d *projectDispatcher) handle(s *task.StateLog) {
	d.window.consume(s)
}

func (d *projectDispatcher) run() {
	nextTaskID := d.startTaskID
	for {
		if d.paused.Load() {
			d.idle.Store(true)
			time.Sleep(d.waitInterval)
			continue
		}
		if d.contract != nil {
			ps := d.contract.LatestProvers()
			isSupported := false
			for _, p := range ps {
				if p.NodeTypes[d.vmTypeID] {
					isSupported = true
					break
				}
			}
			if !isSupported {
				d.idle.Store(true)
				time.Sleep(d.waitInterval)
				continue
			}
		}
		next, err := d.dispatch(nextTaskID)
		if err != nil {
			slog.Error("failed to dispatch task", "error", err, "project_id", d.projectID)
			time.Sleep(d.waitInterval)
			continue
		}
		if nextTaskID == next {
			time.Sleep(d.waitInterval)
		}
		nextTaskID = next
	}
}

func (d *projectDispatcher) dispatch(nextTaskID uint64) (uint64, error) {
	t, err := d.datasource.Retrieve(d.projectID, nextTaskID)
	if err != nil {
		d.idle.Store(true)
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		d.idle.Store(true)
		return nextTaskID, nil
	}
	if err := t.VerifySignature(d.datasourcePubKey); err != nil {
		d.idle.Store(true)
		return 0, errors.Wrap(err, "failed to verify task signature")
	}
	d.idle.Store(false)

	d.window.produce(t)

	metrics.DispatchedTaskNumMtc(d.projectID, t.ProjectVersion)

	if err := d.pubSubs.Publish(t.ProjectID, &p2p.Data{Task: t}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v, task_id %v", t.ProjectID, t.ID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func newProjectDispatcher(persistence Persistence, datasourceURI string, newDatasource NewDatasource, p *contractpkg.Project, pubSubs *p2p.PubSubs, handler *taskStateHandler, datasourcePubKey []byte, contract Contract, vmTypeID uint64) (*projectDispatcher, error) {
	processedTaskID, err := persistence.ProcessedTaskID(p.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch next task_id, project_id %v", p.ID)
	}
	datasource, err := newDatasource(datasourceURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new task retriever")
	}

	proverAmount := &atomic.Uint64{}
	proverAmount.Store(1)
	if v, ok := p.Attributes[contractpkg.RequiredProverAmount]; ok {
		n, err := strconv.ParseUint(string(v), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse project required prover amount, project_id %v", p.ID)
		}
		proverAmount.Store(n)
	}

	window := newWindow(0, pubSubs, handler, persistence)
	paused := atomic.Bool{}
	paused.Store(p.Paused)
	idle := atomic.Bool{}
	idle.Store(true)
	d := &projectDispatcher{
		contract:             contract,
		vmTypeID:             vmTypeID,
		window:               window,
		waitInterval:         3 * time.Second,
		startTaskID:          processedTaskID + 1,
		datasource:           datasource,
		projectID:            p.ID,
		pubSubs:              pubSubs,
		datasourcePubKey:     datasourcePubKey,
		requiredProverAmount: proverAmount,
		paused:               &paused,
		idle:                 &idle,
	}
	go d.run()
	return d, nil
}
