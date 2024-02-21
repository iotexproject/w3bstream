package testutil

import (
	. "github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/clients"
)

func ClientsCreateSession(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		clients.CreateSession,
		func(string, string) error {
			return err
		},
	)
}
