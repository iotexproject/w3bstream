package task

import (
	"sync"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/p2p"
	internaldispatcher "github.com/machinefi/sprout/task/internal/dispatcher"
	"github.com/machinefi/sprout/types"
)

type mockPersistence struct{}

func (m *mockPersistence) Create(tl *types.TaskStateLog, t *types.Task) error {
	return nil
}

func (m *mockPersistence) ProcessedTaskID(projectID uint64) (uint64, error) {
	return 0, nil
}

func (m *mockPersistence) UpsertProcessedTask(projectID, taskID uint64) error {
	return nil
}

func TestDispatcher_handleP2PData(t *testing.T) {
	d := &dispatcher{projectDispatchers: &sync.Map{}}
	t.Run("TaskStateLogNil", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{}, nil)
	})
	t.Run("TaskStateLogNil", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{}, nil)
	})
	t.Run("ProjectDispatcherNotExist", func(t *testing.T) {
		d.handleP2PData(&p2p.Data{TaskStateLog: &types.TaskStateLog{}}, nil)
	})
	t.Run("Success", func(t *testing.T) {
		pid := uint64(1)
		pd := &internaldispatcher.ProjectDispatcher{}
		d.projectDispatchers.Store(pid, pd)

		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(pd, "Handle")
		d.handleP2PData(&p2p.Data{TaskStateLog: &types.TaskStateLog{
			ProjectID: pid,
		}}, nil)
	})
}

// func TestRunDispatcher(t *testing.T) {
// 	r := require.New(t)
// 	p := gomonkey.NewPatches()
// 	defer p.Reset()

// 	p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
// 	p.ApplyFuncReturn(handler.NewTaskStateHandler, nil)
// 	p.ApplyFuncReturn(dispatch, nil)
// 	p.ApplyFuncReturn(dummyDispatch, nil)

// 	err := RunDispatcher(&mockPersistence{}, nil, nil, nil, "", "", "", "", "", "", 0)
// 	r.NoError(err)

// 	err = RunDispatcher(&mockPersistence{}, nil, nil, nil, "", "", "", "", "", "/test", 0)
// 	r.NoError(err)
// }

// func TestDummyDispatch(t *testing.T) {
// 	r := require.New(t)

// 	t.Run("FailedToGetProject", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		pm := &project.Manager{}
// 		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{1})
// 		p.ApplyMethodReturn(pm, "Project", nil, errors.New(t.Name()))

// 		err := dummyDispatch(&mockPersistence{}, nil, pm.ProjectIDs, pm.Project, &sync.Map{}, nil, nil)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		pm := &project.Manager{}
// 		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{1, 1})
// 		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
// 		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
// 		p.ApplyFuncReturn(internaldispatcher.NewProjectDispatcher, &internaldispatcher.ProjectDispatcher{}, nil)

// 		err := dummyDispatch(&mockPersistence{}, nil, pm.ProjectIDs, pm.Project, &sync.Map{}, &p2p.PubSubs{}, nil)
// 		r.NoError(err)
// 	})
// }

// func TestDispatch(t *testing.T) {
// 	r := require.New(t)

// 	t.Run("FailedToListAndWatchProject", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		pm := &project.Manager{}
// 		p.ApplyFuncReturn(contract.ListAndWatchProject, nil, errors.New(t.Name()))

// 		err := dispatch(&mockPersistence{}, nil, pm.Project, &sync.Map{}, nil, nil, "", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		ch := make(chan *contract.BlockProject, 10)
// 		ch <- &contract.BlockProject{
// 			Projects: map[uint64]*contract.Project{
// 				1: {
// 					ID: 1,
// 				},
// 			},
// 		}
// 		m := &sync.Map{}
// 		pm := &project.Manager{}
// 		p.ApplyFuncReturn(contract.ListAndWatchProject, ch, nil)
// 		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
// 		p.ApplyFuncReturn(internaldispatcher.NewProjectDispatcher, &internaldispatcher.ProjectDispatcher{}, nil)
// 		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)

// 		err := dispatch(&mockPersistence{}, nil, pm.Project, m, &p2p.PubSubs{}, nil, "", "")
// 		r.NoError(err)
// 		time.Sleep(20 * time.Millisecond)
// 		close(ch)
// 	})
// }
