package clients

import (
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
)

type Client struct {
	jwk *ioconnect.JWK
}

func (c *Client) HasProjectPermission(projectID uint64) bool {
	// TODO
	return true
}

func (c *Client) KeyAgreementKID() string {
	return c.jwk.KeyAgreementKID()
}

func (c *Client) DID() string {
	return c.jwk.DID()
}
