package testutil

import (
	"context"
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func TopicPublish(p *Patches, err error) *Patches {
	var topic *pubsub.Topic
	return p.ApplyMethodFunc(
		reflect.TypeOf(topic),
		"Publish",
		func(ctx context.Context, data []byte, opts ...pubsub.PubOpt) error {
			return err
		},
	)
}
