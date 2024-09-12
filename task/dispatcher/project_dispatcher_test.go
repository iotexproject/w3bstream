package dispatcher

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/persistence/contract"
	"github.com/iotexproject/w3bstream/task"
)

type mockDatasource struct{}

func (m *mockDatasource) Retrieve(projectID, nextTaskID uint64) (*task.Task, error) {
	return nil, nil
}

func TestProjectDispatcher_handle(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyPrivateMethod(&window{}, "consume", func() {})

	d := &projectDispatcher{}
	d.handle(nil)
}

func TestProjectDispatcher_run(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	d := &projectDispatcher{
		paused: &atomic.Bool{},
	}
	p.ApplyPrivateMethod(d, "dispatch", func(uint64) (uint64, error) {
		return 0, nil
	})
	p.ApplyFunc(time.Sleep, func(time.Duration) { panic(errors.New(t.Name())) })

	r.Panics(func() { d.run() })
}

func TestProjectDispatcher_dispatch(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToRetrieveTask", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		i := atomic.Bool{}
		d := &projectDispatcher{datasource: ds, idle: &i}
		p.ApplyMethodReturn(ds, "Retrieve", nil, errors.New(t.Name()))

		_, err := d.dispatch(0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("NilTask", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		i := atomic.Bool{}
		d := &projectDispatcher{datasource: ds, idle: &i}
		p.ApplyMethodReturn(ds, "Retrieve", nil, nil)

		id, err := d.dispatch(0)
		r.Equal(id, uint64(0))
		r.NoError(err)
	})
	t.Run("FailedToVerifyTaskSignature", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		i := atomic.Bool{}
		d := &projectDispatcher{datasource: ds, idle: &i}
		tk := &task.Task{}
		p.ApplyMethodReturn(ds, "Retrieve", tk, nil)
		p.ApplyMethodReturn(tk, "VerifySignature", errors.New(t.Name()))

		_, err := d.dispatch(0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPublishData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		pubSubs := &p2p.PubSub{}
		i := atomic.Bool{}
		d := &projectDispatcher{
			datasource: ds,
			pubSubs:    pubSubs,
			idle:       &i,
		}
		tk := &task.Task{}
		p.ApplyMethodReturn(ds, "Retrieve", tk, nil)
		p.ApplyMethodReturn(tk, "VerifySignature", nil)
		p.ApplyPrivateMethod(&window{}, "produce", func() {})
		p.ApplyMethodReturn(pubSubs, "Publish", errors.New(t.Name()))

		_, err := d.dispatch(0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		pubSubs := &p2p.PubSub{}
		i := atomic.Bool{}
		d := &projectDispatcher{
			datasource: ds,
			pubSubs:    pubSubs,
			idle:       &i,
		}
		tk := &task.Task{}
		p.ApplyMethodReturn(ds, "Retrieve", tk, nil)
		p.ApplyMethodReturn(tk, "VerifySignature", nil)
		p.ApplyPrivateMethod(&window{}, "produce", func() {})
		p.ApplyMethodReturn(pubSubs, "Publish", nil)

		_, err := d.dispatch(0)
		r.NoError(err)
	})
}

func TestNewProjectDispatcher(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToFetchNextTaskID", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), errors.New(t.Name()))

		_, err := newProjectDispatcher(ps, "", nil, &contract.Project{}, nil, nil, nil, nil, 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewTaskRetriever", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
		nd := func(string) (datasource.Datasource, error) { return nil, errors.New(t.Name()) }

		_, err := newProjectDispatcher(ps, "", nd, &contract.Project{}, nil, nil, nil, nil, 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectRequiredProverAmount", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
		nd := func(string) (datasource.Datasource, error) { return nil, nil }

		_, err := newProjectDispatcher(ps, "", nd, &contract.Project{
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmount: []byte("err")},
		}, nil, nil, nil, nil, 0)
		r.ErrorContains(err, "failed to parse project required prover amount")
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
		p.ApplyFuncReturn(newWindow, nil)
		p.ApplyPrivateMethod(&projectDispatcher{}, "run", func() {})
		nd := func(string) (datasource.Datasource, error) { return nil, nil }

		paused := true
		_, err := newProjectDispatcher(ps, "", nd, &contract.Project{
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmount: []byte("1")},
			Paused:     paused,
		}, nil, nil, nil, nil, 0)
		time.Sleep(10 * time.Millisecond)
		r.NoError(err)
	})
}
