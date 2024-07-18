package dispatcher

import (
	"sync/atomic"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/project"
)

func TestNewLocal(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToNewPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))

		_, err := NewLocal(&mockPersistence{}, nil, nil, "", "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToGetProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyMethodReturn(pm, "Project", nil, errors.New(t.Name()))
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToAddPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))
		p.ApplyMethodReturn(pm, "Project", nil, nil)

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", "", []byte(""), 0)
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

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", "", []byte(""), 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pm := &mockProjectManager{}
		w := &window{}
		a := atomic.Uint64{}
		p.ApplyFuncReturn(p2p.NewPubSubs, &p2p.PubSubs{}, nil)
		p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p.ApplyFuncReturn(newProjectDispatcher, &projectDispatcher{window: w, requiredProverAmount: &a}, nil)
		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{0, 0})
		p.ApplyMethodReturn(pm, "Project", &project.Project{}, nil)
		p.ApplyPrivateMethod(w, "setSize", func(uint64) {})

		_, err := NewLocal(&mockPersistence{}, nil, pm, "", "", "", "", "", []byte(""), 0)
		r.NoError(err)
	})
}
