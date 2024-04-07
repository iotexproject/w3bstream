package project

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/project/contracts"
)

func TestNewManager(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToDialChain", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)

	t.Run("FailedToNewContracts", func(t *testing.T) {
		p = p.ApplyFuncReturn(contracts.NewContracts, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(contracts.NewContracts, nil, nil)

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Manager{}, "watchProjectContract", func() error { return nil })

		_, err := NewManager("", "", "", "")
		r.NoError(err)
	})
}
