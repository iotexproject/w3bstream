package task

import (
	"sync"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	internaldispatcher "github.com/machinefi/sprout/task/internal/dispatcher"
	"github.com/machinefi/sprout/task/internal/handler"
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

func TestRunDispatcher(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToNewPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))

		err := RunDispatcher(&mockPersistence{}, nil, nil,
			"", "", "", []byte(""), 0,
			nil, nil, nil)

		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
		p.ApplyFuncReturn(handler.NewTaskStateHandler, nil)
		p.ApplyFuncReturn(dispatch, nil)

		err := RunDispatcher(&mockPersistence{}, nil, nil,
			"", "", "", []byte(""), 0,
			nil, nil, nil)
		r.NoError(err)
	})
}

func TestRunLocalDispatcher(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToNewPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))

		err := RunLocalDispatcher(&mockPersistence{}, nil, nil, nil,
			"", "", "", []byte(""), 0)

		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToGetProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)

		err := RunLocalDispatcher(&mockPersistence{}, nil, func() []uint64 {
			return []uint64{uint64(0)}
		}, func(projectID uint64) (*project.Project, error) {
			return nil, errors.New(t.Name())
		},
			"", "", "", []byte(""), 0)

		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToAddPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))

		err := RunLocalDispatcher(&mockPersistence{}, nil, func() []uint64 {
			return []uint64{uint64(0)}
		}, func(projectID uint64) (*project.Project, error) {
			return nil, nil
		},
			"", "", "", []byte(""), 0)

		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToNewProjectDispatch", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(internaldispatcher.NewProjectDispatcher, nil, errors.New(t.Name()))

		err := RunLocalDispatcher(&mockPersistence{}, nil, func() []uint64 {
			return []uint64{uint64(0)}
		}, func(projectID uint64) (*project.Project, error) {
			return &project.Project{}, nil
		},
			"", "", "", []byte(""), 0)

		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(internaldispatcher.NewProjectDispatcher, &internaldispatcher.ProjectDispatcher{}, nil)

		err := RunLocalDispatcher(&mockPersistence{}, nil, func() []uint64 {
			return []uint64{uint64(0)}
		}, func(projectID uint64) (*project.Project, error) {
			return &project.Project{}, nil
		},
			"", "", "", []byte(""), 0)

		r.NoError(err)
	})
}
