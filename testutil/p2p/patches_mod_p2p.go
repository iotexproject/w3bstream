package p2p

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/p2p"
)

func P2pNewPubSubs(p *Patches, ps *p2p.PubSubs, err error) *Patches {
	return p.ApplyFunc(
		p2p.NewPubSubs,
		func(handle p2p.HandleSubscriptionMessage, bootNodeMultiaddr string, iotexChainID int) (*p2p.PubSubs, error) {
			return ps, err
		},
	)
}

func P2pPubSubsAdd(p *Patches, err error) *Patches {
	var ps *p2p.PubSubs
	return p.ApplyMethodFunc(
		reflect.TypeOf(ps),
		"Add",
		func(projectID uint64) error {
			return err
		},
	)
}
