package clients_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/clients"
)

func TestCreateSession(t *testing.T) {
	_ = clients.NewManager()
	r := require.New(t)

	did1 := "did:ethr:0x9d9250fb4e08ba7a858fe7196a6ba946c6083ff0"
	did2 := "unit_test_not_exists"
	t.Run("NotExists", func(t *testing.T) {
		err := clients.CreateSession("any", did2)
		r.NotNil(err)
	})
	t.Run("Success", func(t *testing.T) {
		err := clients.CreateSession("any", did1)
		r.Nil(err)
	})
}

func TestVerifySessionAndProjectPermission(t *testing.T) {
	_ = clients.NewManager()
	r := require.New(t)

	did1 := "did:ethr:0x9d9250fb4e08ba7a858fe7196a6ba946c6083ff0"

	err := clients.CreateSession("any_right", did1)
	r.Nil(err)

	t.Run("NotExists", func(t *testing.T) {
		t.Run("TokenNotExists", func(t *testing.T) {
			_, err = clients.VerifySessionAndProjectPermission("any_wrong", 1)
			r.NotNil(err)
		})
		t.Run("ProjectNotExists", func(t *testing.T) {
			_, err = clients.VerifySessionAndProjectPermission("any_right", 1000000)
			r.NotNil(err)
		})
	})
	t.Run("Success", func(t *testing.T) {
		did, err := clients.VerifySessionAndProjectPermission("any_right", 1)
		r.Nil(err)
		r.Equal(did, did1)
	})
}
