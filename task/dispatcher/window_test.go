package dispatcher

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/task"
)

func TestWindow(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	tk := &task.Task{ID: 1}
	dt := &dispatchedTask{
		task: tk,
	}
	dt.finished.Store(true)
	ps := &mockPersistence{}
	p.ApplyPrivateMethod(dt, "handleState", func(s *task.StateLog) {})
	p.ApplyFuncReturn(newDispatchedTask, dt)
	p.ApplyMethodReturn(ps, "UpsertProcessedTask", errors.New(t.Name()))
	w := newWindow(10, nil, nil, ps)
	r.True(w.isEmpty())
	r.False(w.isFull())

	w.produce(tk)
	r.False(w.isEmpty())

	w.consume(&task.StateLog{TaskID: 1})
	r.True(w.isEmpty())

	dt.finished.Store(false)

	w.produce(tk)
	r.False(w.isEmpty())

	w.consume(&task.StateLog{TaskID: 0})
	r.False(w.isEmpty())
	w.consume(&task.StateLog{TaskID: 1})
	r.False(w.isEmpty())
}
