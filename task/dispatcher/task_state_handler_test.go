package dispatcher

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/persistence/postgres"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

type mockOutput struct{}

func (m *mockOutput) Output(task *task.Task, proof []byte) (string, error) {
	return "", nil
}

func TestTaskStateHandler_handle(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToCreateTaskStateLog", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		h := &taskStateHandler{persistence: ps}
		p.ApplyMethodReturn(ps, "Create", errors.New(t.Name()))

		r.False(h.handle(time.Now(), &task.StateLog{}, &task.Task{}))
	})
	t.Run("StateFailed", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		h := &taskStateHandler{persistence: ps}
		p.ApplyMethodReturn(ps, "Create", nil)

		r.True(h.handle(time.Now(), &task.StateLog{State: task.StateFailed}, &task.Task{}))
	})
	t.Run("StateDispatched", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		h := &taskStateHandler{persistence: ps}
		p.ApplyMethodReturn(ps, "Create", nil)

		r.False(h.handle(time.Now(), &task.StateLog{State: task.StateDispatched}, &task.Task{}))
	})
	t.Run("FailedToGetProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		pm := &project.Manager{}
		h := &taskStateHandler{
			persistence:    ps,
			projectManager: pm,
		}
		p.ApplyMethodReturn(ps, "Create", nil)
		p.ApplyMethodReturn(pm, "Project", nil, errors.New(t.Name()))

		r.False(h.handle(time.Now(), &task.StateLog{State: task.StateProved}, &task.Task{}))
	})
	t.Run("FailedToGetProjectDefaultConfig", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		pm := &project.Manager{}
		h := &taskStateHandler{
			persistence:    ps,
			projectManager: pm,
		}
		p.ApplyMethodReturn(ps, "Create", nil)
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", nil, errors.New(t.Name()))

		r.False(h.handle(time.Now(), &task.StateLog{State: task.StateProved}, &task.Task{}))
	})
	t.Run("FailedToNewOutput", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		pm := &project.Manager{}
		h := &taskStateHandler{
			persistence:    ps,
			projectManager: pm,
		}
		p.ApplyMethodReturn(ps, "Create", nil)
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)
		p.ApplyFuncReturn(output.New, nil, errors.New(t.Name()))

		r.True(h.handle(time.Now(), &task.StateLog{State: task.StateProved}, &task.Task{}))
	})
	t.Run("FailedToOutput", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		pm := &project.Manager{}
		h := &taskStateHandler{
			persistence:    ps,
			projectManager: pm,
		}
		p.ApplyMethodReturn(ps, "Create", nil)
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)
		p.ApplyFuncReturn(output.New, &mockOutput{}, nil)
		p.ApplyMethodReturn(&mockOutput{}, "Output", "", errors.New(t.Name()))

		r.True(h.handle(time.Now(), &task.StateLog{State: task.StateProved}, &task.Task{}))
	})
	t.Run("FailedToOutput", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &postgres.Postgres{}
		pm := &project.Manager{}
		h := &taskStateHandler{
			persistence:    ps,
			projectManager: pm,
		}
		p.ApplyMethodReturn(ps, "Create", nil)
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)
		p.ApplyFuncReturn(output.New, &mockOutput{}, nil)
		p.ApplyMethodReturn(&mockOutput{}, "Output", "", nil)

		r.True(h.handle(time.Now(), &task.StateLog{State: task.StateProved}, &task.Task{}))
	})
}
