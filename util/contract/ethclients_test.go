package contract_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/machinefi/sprout/util/contract"
)

var endpoint = "https://babel-api.testnet.iotex.io"

func TestNewEthClient(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToDialEndpoint", func(t *testing.T) {
		c, err := NewEthClient("")
		r.Error(err)
		r.Nil(c)
	})

	t.Run("Success", func(t *testing.T) {
		var (
			c1  Client
			c2  Client
			err error
		)
		t.Run("CreateNew", func(t *testing.T) {
			c1, err = NewEthClient(endpoint)
			r.NoError(err)
		})
		t.Run("Exists", func(t *testing.T) {
			c2, err = NewEthClient(endpoint)
			r.NoError(err)
			r.Equal(c2.Endpoint(), c1.Endpoint())
			r.Equal(c2.RefCount(), c1.RefCount())
			r.Equal(c2.RefCount(), int32(2))
		})
	})
}

type MockClient struct {
	Client
}

func (c *MockClient) Endpoint() string {
	return "non-existence"
}

func TestReleaseClient(t *testing.T) {
	r := require.New(t)

	c, err := NewEthClient(endpoint)
	r.NoError(err)
	for c.RefCount() > 0 {
		ReleaseClient(c)
	}

	// not found
	ReleaseClient(c)
}
