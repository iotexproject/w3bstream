package clients

import (
	"sync"

	"github.com/pkg/errors"
)

var sessions sync.Map

func CreateSession(vctoken string, clientdid string) error {
	client, ok := manager.GetByClientDID(clientdid)
	if !ok {
		return errors.Errorf("client did not exists: %s", clientdid)
	}
	sessions.Store(vctoken, client)
	return nil
}

func VerifySessionAndProjectPermission(vctoken string, projectID uint64) error {
	v, exists := sessions.Load(vctoken)
	if !exists || v == nil {
		return errors.Errorf("invalid token or expired")
	}

	if _, exists = v.(*Client).projects[projectID]; !exists {
		return errors.Errorf("project permission denied")
	}
	return nil
}
