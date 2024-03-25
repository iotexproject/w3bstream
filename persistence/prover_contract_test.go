package persistence

import (
	"sync"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/persistence/prover"
	"github.com/machinefi/sprout/testutil"
)

func TestNewProver(t *testing.T) {
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
	t.Run("CreateProverContractInstance", func(t *testing.T) {
		t.Run("FailedToCreateProverContractInstance", func(t *testing.T) {
			p := gomonkey.NewPatches()
			defer p.Reset()

			p = testutil.EthClientDial(p, nil, nil)
			p = p.ApplyFuncReturn(prover.NewProver, nil, errors.New(t.Name()))

			z, err := NewProver("any", "any")
			r.Nil(z)
			r.ErrorContains(err, t.Name())
		})
	})
	t.Run("LoopFetchProverFromContractUntilFetchedEmpty", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = testutil.EthClientDial(p, nil, nil)
		p = p.ApplyFuncReturn(prover.NewProver, &prover.Prover{
			ProverCaller:     prover.ProverCaller{},
			ProverFilterer:   prover.ProverFilterer{},
			ProverTransactor: prover.ProverTransactor{},
		}, nil)

		p = p.ApplyMethodSeq(&prover.ProverCaller{}, "Provers", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{
					struct {
						Id     string
						Paused bool
					}{}, errors.New(t.Name()),
				},
				Times: 1,
			},
			{
				Values: gomonkey.Params{
					struct {
						Id     string
						Paused bool
					}{
						Id: "any",
					}, nil,
				},
				Times: 1,
			},
			{
				Values: gomonkey.Params{
					struct {
						Id     string
						Paused bool
					}{}, nil,
				},
				Times: 1,
			},
		})
		z, err := NewProver("any", "any")
		r.NotNil(z)
		r.Nil(err)
	})
}

func TestProver_GetAll(t *testing.T) {
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
