package p2p

import (
	"github.com/machinefi/sprout/testutil"
	"github.com/pkg/errors"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
)

func TestNewPubSubs(t *testing.T) {
	require := require.New(t)

	var handle HandleSubscriptionMessage = nil
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2
	_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
	require.NoError(err)
}

func TestAdd(t *testing.T) {
	require := require.New(t)

	p, err := newPubSubs()
	require.NoError(err)

	projectID := uint64(0x1)
	t.Run("AddNewProject", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
	})

	t.Run("AddRepeatProject", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
	})
}

func TestDelete(t *testing.T) {
	require := require.New(t)

	p, err := newPubSubs()
	require.NoError(err)

	projectID := uint64(0x1)

	t.Run("DelNilProject", func(t *testing.T) {
		p.Delete(projectID)
	})

	t.Run("DelProject", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
		p.Delete(projectID)
	})
}

func TestPublish(t *testing.T) {
	require := require.New(t)

	p, err := newPubSubs()
	require.NoError(err)
	patches := NewPatches()

	projectID := uint64(0x1)
	d := &Data{
		Task:         nil,
		TaskStateLog: nil,
	}

	t.Run("PublishNotExistProject", func(t *testing.T) {
		err = p.Publish(projectID, d)
		require.EqualError(err, "project 1 topic not exist")
	})

	t.Run("Publish", func(t *testing.T) {
		err = p.Add(projectID)
		require.NoError(err)
		err = p.Publish(projectID, d)
		require.NoError(err)
	})

	t.Run("MarshalFailed", func(t *testing.T) {
		err = p.Add(projectID)
		require.NoError(err)
		patches = testutil.JsonMarshal(patches, []byte("any"), errors.New(t.Name()))
		err = p.Publish(projectID, d)
		require.ErrorContains(err, t.Name())
	})
}

func newPubSubs() (*PubSubs, error) {
	var handle HandleSubscriptionMessage = nil
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2
	return NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
}
