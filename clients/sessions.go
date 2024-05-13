package clients

import (
	"github.com/pkg/errors"
)

func VerifyProjectPermissionByClientDID(clientID string, projectID uint64) error {
	client := manager.ClientByDID(clientID)
	if client != nil && client.HasProjectPermission(projectID) {
		return nil
	}
	return errors.Errorf("no project permission %s %d", clientID, projectID)
}

func ClientByClientID(clientID string) *Client {
	return manager.ClientByDID(clientID)
}
