package dispatcher

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/task"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
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

func TestProjectDispatcher_dispatch(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToRetrieveTask", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		d := &projectDispatcher{datasource: ds}
		p.ApplyMethodReturn(ds, "Retrieve", nil, errors.New(t.Name()))

		_, err := d.dispatch(0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("NilTask", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		d := &projectDispatcher{datasource: ds}
		p.ApplyMethodReturn(ds, "Retrieve", nil, nil)

		id, err := d.dispatch(0)
		r.Equal(id, uint64(0))
		r.NoError(err)
	})
	t.Run("FailedToVerifyTaskSignature", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ds := &mockDatasource{}
		d := &projectDispatcher{datasource: ds}
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
		pubSubs := &p2p.PubSubs{}
		d := &projectDispatcher{
			datasource: ds,
			pubSubs:    pubSubs,
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
		pubSubs := &p2p.PubSubs{}
		d := &projectDispatcher{
			datasource: ds,
			pubSubs:    pubSubs,
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

		_, err := newProjectDispatcher(ps, "", nil, &contract.Project{}, nil, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewTaskRetriever", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
		nd := func(string) (datasource.Datasource, error) { return nil, errors.New(t.Name()) }

		_, err := newProjectDispatcher(ps, "", nd, &contract.Project{}, nil, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectRequiredProverAmount", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &mockPersistence{}
		p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
		nd := func(string) (datasource.Datasource, error) { return nil, nil }

		_, err := newProjectDispatcher(ps, "", nd, &contract.Project{
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmountHash: []byte("err")},
		}, nil, nil, nil)
		r.ErrorContains(err, "failed to parse project required prover amount")
	})
	// this cannot work in ci, but can work locally
	// t.Run("Success", func(t *testing.T) {
	// 	p := gomonkey.NewPatches()
	// 	defer p.Reset()

	// 	ps := &mockPersistence{}
	// 	p.ApplyMethodReturn(ps, "ProcessedTaskID", uint64(0), nil)
	// 	p.ApplyFuncReturn(newWindow, nil)
	// 	p.ApplyPrivateMethod(&projectDispatcher{}, "run", func() {})
	// 	nd := func(string) (datasource.Datasource, error) { return nil, nil }

	// 	_, err := newProjectDispatcher(ps, "", nd, &contract.Project{
	// 		Attributes: map[common.Hash][]byte{contract.RequiredProverAmountHash: []byte("1")},
	// 	}, nil, nil, nil)
	// 	r.NoError(err)
	// })
}
