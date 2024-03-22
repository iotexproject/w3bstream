package persistence

import (
	"sync"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/persistence/znode"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/testutil/znodecontract"
)

func PatchNewZnode(p *gomonkey.Patches, node *Prover, err error) *gomonkey.Patches {
	return p.ApplyFunc(
		NewProver,
		func(_, _ string) (*Prover, error) {
			return node, err
		},
	)
}

func TestNewZnode(t *testing.T) {
	r := require.New(t)

	t.Run("DailEth", func(t *testing.T) {
		t.Run("FailedToDialEthClient", func(t *testing.T) {
			p := gomonkey.NewPatches()
			defer p.Reset()

			p = testutil.EthClientDial(p, nil, errors.New(t.Name()))

			z, err := NewProver("any", "any")
			r.Nil(z)
			r.ErrorContains(err, t.Name())
		})
	})
	t.Run("CreateZnodeContractInstance", func(t *testing.T) {
		t.Run("FailedToCreateZnodeContractInstance", func(t *testing.T) {
			p := gomonkey.NewPatches()
			defer p.Reset()

			p = testutil.EthClientDial(p, nil, nil)
			p = znodecontract.PatchNewZnode(p, nil, errors.New(t.Name()))

			z, err := NewProver("any", "any")
			r.Nil(z)
			r.ErrorContains(err, t.Name())
		})
	})
	t.Run("LoopFetchZnodeFromContractUntilFetchedEmpty", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = testutil.EthClientDial(p, nil, nil)
		p = znodecontract.PatchNewZnode(p, &znode.Znode{
			ZnodeCaller:     znode.ZnodeCaller{},
			ZnodeFilterer:   znode.ZnodeFilterer{},
			ZnodeTransactor: znode.ZnodeTransactor{},
		}, nil)
		p = znodecontract.PatchZnodeCallerZnodesSeq(p,
			struct {
				Did    string
				Paused bool
				Err    error
			}{Err: errors.New(t.Name())},
			struct {
				Did    string
				Paused bool
				Err    error
			}{Did: "any", Err: nil},
			struct {
				Did    string
				Paused bool
				Err    error
			}{Did: "", Err: nil},
		)
		z, err := NewProver("any", "any")
		r.NotNil(z)
		r.Nil(err)
	})
}

func TestZNode_GetAll(t *testing.T) {
	r := require.New(t)

	zn := &Prover{
		mux: sync.Mutex{},
		proverIDs: map[string]bool{
			"any1": true,
			"any2": true,
		},
	}

	nodes := zn.GetAll()
	r.Equal(len(nodes), len(zn.proverIDs))
}
