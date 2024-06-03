package dispatcher

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task"
)

type mockPersistence struct{}

func (m *mockPersistence) Create(tl *task.StateLog, t *task.Task) error {
	return nil
}
func (m *mockPersistence) ProcessedTaskID(projectID uint64) (uint64, error) {
	return 0, nil
}
func (m *mockPersistence) UpsertProcessedTask(projectID, taskID uint64) error {
	return nil
}

type mockProjectManager struct{}

func (m *mockProjectManager) ProjectIDs() []uint64 {
	return []uint64{uint64(0)}
}
func (m *mockProjectManager) Project(projectID uint64) (*project.Project, error) {
	return nil, nil
}

func TestDispatcher_handleP2PData(t *testing.T) {
	d := &Dispatcher{projectDispatchers: &sync.Map{}}
	t.Run("TaskStateLogNil", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{}, nil)
	})
	t.Run("TaskStateLogNil", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{}, nil)
	})
	t.Run("ProjectDispatcherNotExist", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{TaskStateLog: &task.StateLog{}}, nil)
	})
	t.Run("Success", func(t *testing.T) {
		pid := uint64(1)
		pd := &projectDispatcher{}
		d.projectDispatchers.Store(pid, pd)

		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(pd, "handle", func(*task.StateLog) {})
		d.handleP2PData(&p2p.Data{TaskStateLog: &task.StateLog{
			ProjectID: pid,
		}}, nil)
	})
}

func TestDispatcher_setWindowSize(t *testing.T) {
	r := require.New(t)
	po := &scheduler.ProjectEpochOffsets{}
	c := &contract.Contract{}
	d := &Dispatcher{
		projectDispatchers: &sync.Map{},
		projectOffsets:     po,
		contract:           c,
	}
	t.Run("ProjectNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(po, "Projects", nil)
		d.setWindowSize(1)
	})
	t.Run("ContractProjectNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(po, "Projects", []*scheduler.ScheduledProject{{ID: 1, ScheduledBlockNumber: 0}})
		p.ApplyMethodReturn(c, "Project", nil)
		d.setWindowSize(1)
	})
	t.Run("ProjectDispatcherNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(po, "Projects", []*scheduler.ScheduledProject{{ID: 1, ScheduledBlockNumber: 0}})
		p.ApplyMethodReturn(c, "Project", &contract.Project{ID: 1})
		d.setWindowSize(1)
	})
	t.Run("FailedToParseProjectRequiredProverAmount", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(po, "Projects", []*scheduler.ScheduledProject{{ID: 1, ScheduledBlockNumber: 0}})
		p.ApplyMethodReturn(c, "Project", &contract.Project{
			ID:         1,
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmountHash: []byte("err")},
		})
		d.projectDispatchers.Store(uint64(1), &projectDispatcher{})
		d.setWindowSize(1)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(po, "Projects", []*scheduler.ScheduledProject{{ID: 1, ScheduledBlockNumber: 0}})
		p.ApplyMethodReturn(c, "Project", &contract.Project{
			ID:         1,
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmountHash: []byte("2")},
		})
		size := atomic.Uint64{}
		d.projectDispatchers.Store(uint64(1), &projectDispatcher{
			window: &window{size: &size},
		})
		d.setWindowSize(1)
		r.Equal(size.Load(), uint64(2))
	})
}

func TestNew(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToNewPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))

		_, err := New(&mockPersistence{}, nil, nil, "", "", "", "", []byte(""), 0, nil, nil, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
		p.ApplyFuncReturn(newTaskStateHandler, nil)

		_, err := New(&mockPersistence{}, nil, nil, "", "", "", "", []byte(""), 0, nil, nil, nil, nil)
		r.NoError(err)
	})
}

func TestNewLocal(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToNewPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))

		_, err := NewLocal(&mockPersistence{}, nil, nil, "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToGetProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyMethodReturn(pm, "Project", nil, errors.New(t.Name()))
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToAddPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))
		p.ApplyMethodReturn(pm, "Project", nil, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProjectDispatch", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(newProjectDispatcher, nil, errors.New(t.Name()))
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(newProjectDispatcher, &projectDispatcher{}, nil)
		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{0, 0})
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", []byte(""), 0)
		r.NoError(err)
	})
}

func TestDispatcher_setProjectDispatcher(t *testing.T) {
	paused := true
	cp := &contract.Project{
		ID:     1,
		Uri:    "http://test.com",
		Paused: paused,
	}
	pc := &contract.Contract{}
	mp := &mockProjectManager{}
	ps := &p2p.PubSubs{}
	t.Run("ProjectExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(pc, "LatestProject", cp)

		d := &Dispatcher{
			contract:           pc,
			projectDispatchers: &sync.Map{},
		}
		projectDispatcher := &projectDispatcher{
			paused: &atomic.Bool{},
		}
		d.projectDispatchers.Store(uint64(1), projectDispatcher)
		d.setProjectDispatcher(1)
	})
	t.Run("ProjectURIIsEmpty", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		ncp := *cp
		ncp.Uri = ""

		p.ApplyMethodReturn(pc, "LatestProject", &ncp)

		d.setProjectDispatcher(1)
	})
	t.Run("FailedToGetContractProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		p.ApplyMethodReturn(pc, "LatestProject", nil)

		d.setProjectDispatcher(1)
	})
	t.Run("FailedToGetProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		p.ApplyMethodReturn(pc, "LatestProject", &contract.Project{})
		p.ApplyMethodReturn(mp, "Project", nil, errors.New(t.Name()))

		d.setProjectDispatcher(1)
	})
	t.Run("FailedToAddPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			pubSubs:            ps,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		p.ApplyMethodReturn(pc, "LatestProject", &contract.Project{})
		p.ApplyMethodReturn(mp, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))

		d.setProjectDispatcher(1)
	})
	t.Run("FailedToNewProjectDispatcher", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			pubSubs:            ps,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		p.ApplyMethodReturn(pc, "LatestProject", &contract.Project{})
		p.ApplyMethodReturn(mp, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(newProjectDispatcher, nil, errors.New(t.Name()))

		d.setProjectDispatcher(1)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		d := &Dispatcher{
			contract:           pc,
			pubSubs:            ps,
			projectManager:     mp,
			projectDispatchers: &sync.Map{},
		}
		p.ApplyMethodReturn(pc, "LatestProject", &contract.Project{})
		p.ApplyMethodReturn(mp, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(newProjectDispatcher, nil, nil)

		d.setProjectDispatcher(1)
	})
}

func TestDispatcher_Run(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	d := &Dispatcher{
		local: true,
	}
	d.Run()

	projectNotification := make(chan uint64, 10)
	projectNotification <- 1
	chainHeadNotification := make(chan uint64, 10)
	chainHeadNotification <- 1
	d = &Dispatcher{
		projectNotification:   projectNotification,
		chainHeadNotification: chainHeadNotification,
		contract:              &contract.Contract{},
		projectDispatchers:    &sync.Map{},
	}

	p.ApplyMethodReturn(d.contract, "LatestProjects", []*contract.Project{{}})
	p.ApplyPrivateMethod(d, "setProjectDispatcher", func(*contract.Project) {})
	p.ApplyPrivateMethod(d, "setWindowSize", func(head uint64) {})

	d.Run()
	time.Sleep(10 * time.Millisecond)
}
