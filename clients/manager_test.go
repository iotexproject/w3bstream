package clients_test

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/testutil"
)

func TestClientManager(t *testing.T) {
	var (
		r = require.New(t)
		p = gomonkey.NewPatches()
		m *clients.Manager
	)

	t.Run("FailedToFillMockClients", func(t *testing.T) {
		p = testutil.JsonUnmarshal(p, errors.New(t.Name()))

		defer p.Reset()
		defer func() {
			if err := recover(); err != nil {
				r.Contains(fmt.Sprint(err), t.Name())
			}
		}()

		_ = clients.NewManager()
	})

	t.Run("NewManager", func(t *testing.T) {
		m = clients.NewManager()
		r.NotNil(m)
		r.Equal(m, clients.NewManager())
	})

	t.Run("AddAndGetClient", func(t *testing.T) {
		c, ok := m.GetByClientDID("did:ethr:0x9d9250fb4e08ba7a858fe7196a6ba946c6083ff0")
		r.NotNil(c)
		r.True(ok)
		c, ok = m.GetByClientDID("not_exists")
		r.Nil(c)
		r.False(ok)

		m.AddClient(&clients.Client{
			ClientDID: "unit_test_added",
			Projects:  []uint64{1, 2, 3},
		})
		c, ok = m.GetByClientDID("unit_test_added")
		r.NotNil(c)
		r.True(ok)
		r.NotNil(c.Metadata)
		r.Len(c.Projects, 3)
	})
}
