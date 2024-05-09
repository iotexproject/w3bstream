package dispatcher

import (
	"context"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/task"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestDispatchedTask_handleState(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	h := &taskStateHandler{}
	p.ApplyPrivateMethod(h, "handle", func() bool { return true })

	d := &dispatchedTask{
		cancel:  func() {},
		handler: h,
	}
	d.handleState(nil)
	r.Equal(d.finished.Load(), true)
}

func TestDispatchedTask_runWatchdog(t *testing.T) {
	t.Run("ContextCancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		d := &dispatchedTask{
			task: &task.Task{},
		}
		d.runWatchdog(ctx)
	})
	t.Run("Retry", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		retryChan := make(chan time.Time, 10)
		timeoutChan := make(chan time.Time, 10)
		retryChan <- time.Now()
		pubSubs := &p2p.PubSubs{}
		p.ApplyMethodReturn(pubSubs, "Publish", errors.New(t.Name()))
		p.ApplyFuncSeq(time.After, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{retryChan},
				Times:  1,
			},
			{
				Values: gomonkey.Params{timeoutChan},
				Times:  1,
			},
		})

		ctx, cancel := context.WithCancel(context.Background())
		d := &dispatchedTask{
			task: &task.Task{ProjectID: 1},
		}
		go d.runWatchdog(ctx)
		time.Sleep(1 * time.Millisecond)
		cancel()
		time.Sleep(10 * time.Millisecond)
	})
	t.Run("Timeout", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		retryChan := make(chan time.Time, 10)
		timeoutChan := make(chan time.Time, 10)
		timeoutChan <- time.Now()
		p.ApplyFuncSeq(time.After, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{retryChan},
				Times:  1,
			},
			{
				Values: gomonkey.Params{timeoutChan},
				Times:  1,
			},
		})

		ctx, cancel := context.WithCancel(context.Background())
		d := &dispatchedTask{
			task:    &task.Task{ID: 1},
			timeOut: func(*task.StateLog) {},
		}
		go d.runWatchdog(ctx)
		time.Sleep(1 * time.Millisecond)
		cancel()
		time.Sleep(10 * time.Millisecond)
	})
}

func TestNewDispatchedTask(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyPrivateMethod(&dispatchedTask{}, "runWatchdog", func(context.Context) {})

	newDispatchedTask(nil, nil, nil, nil)
}
