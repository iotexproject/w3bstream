package task

import (
	"context"
	"fmt"
	. "github.com/agiledragon/gomonkey/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
	"reflect"
	"time"
)

func p2pNewPubSubs(p *Patches, ps *p2p.PubSubs, err error) *Patches {
	return p.ApplyFunc(
		p2p.NewPubSubs,
		func(handle p2p.HandleSubscriptionMessage, bootNodeMultiaddr string, iotexChainID int) (*p2p.PubSubs, error) {
			return ps, err
		},
	)
}

func p2pPubSubsAdd(p *Patches, err error) *Patches {
	var ps *p2p.PubSubs
	return p.ApplyMethodFunc(
		reflect.TypeOf(ps),
		"Add",
		func(projectID uint64) error {
			return err
		},
	)
}

func projectManagerGetAllProjectID(p *Patches, ids []uint64) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"GetAllProjectID",
		func() []uint64 {
			return ids
		},
	)
}

func projectManagerGetNotify(p *Patches, c <-chan uint64) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"GetNotify",
		func() <-chan uint64 {
			return c
		},
	)
}

func projectManagerGet(p *Patches, err error) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"Get",
		func(projectID uint64, version string) (*project.Config, error) {
			return nil, err
		},
	)
}

func persistencePostgresUpdateState(p *Patches, err error) *Patches {
	//var pg *persistence.Postgres
	pg := &persistence.Postgres{}
	return p.ApplyMethodFunc(
		reflect.TypeOf(pg),
		"UpdateState",
		func(postgres *persistence.Postgres, taskID string, state types.TaskState, comment string, createdAt time.Time) error {
			fmt.Println("up")
			return err
		},
	)
}

func persistencePostgresFetchByID(p *Patches, task *types.Task, err error) *Patches {
	var pg *persistence.Postgres
	return p.ApplyMethodFunc(
		reflect.TypeOf(pg),
		"FetchByID",
		func(taskID string) (*types.Task, error) {
			return task, err
		},
	)
}

func projectConfigGetOutput(p *Patches, err error) *Patches {
	var config *project.Config
	return p.ApplyMethodFunc(
		reflect.TypeOf(config),
		"GetOutput",
		func(privateKeyECDSA, privateKeyED25519 string) (output.Output, error) {
			return nil, err
		},
	)
}

func processorReportSuccess(p *Patches) *Patches {
	var pro *Processor
	return ApplyPrivateMethod(pro, "reportSuccess", func(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {})
}

func processorReportFail(p *Patches) *Patches {
	var pro *Processor
	return ApplyPrivateMethod(pro, "reportFail", func(taskID string, err error, topic *pubsub.Topic) {})
}

func vmHandlerHandle(p *Patches, err error) *Patches {
	var hander *vm.Handler
	return p.ApplyMethodFunc(
		reflect.TypeOf(hander),
		"Handle",
		func(msgs []*types.Message, vmtype types.VM, code string, expParam string) ([]byte, error) {
			return nil, err
		},
	)
}

func topicPublish(p *Patches, err error) *Patches {
	var topic *pubsub.Topic
	return p.ApplyMethodFunc(
		reflect.TypeOf(topic),
		"Publish",
		func(ctx context.Context, data []byte, opts ...pubsub.PubOpt) error {
			return err
		},
	)
}
