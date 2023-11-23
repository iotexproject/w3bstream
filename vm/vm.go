package vm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/vm/server"
)

type Processor struct {
	endpoints   map[Type]string
	instanceMgr *server.Mgr
}

func (r *Processor) Process(msg *proto.Message, vmtype Type, code string, expParam string) ([]byte, error) {
	endpoint, ok := r.endpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(msg, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	slog.Debug(fmt.Sprintf("acquire %s instance success", vmtype))
	defer r.instanceMgr.Release(msg.ProjectID, ins)

	res, err := ins.Execute(context.Background(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewProcessor(endpoints map[Type]string) *Processor {
	return &Processor{
		endpoints:   endpoints,
		instanceMgr: server.NewMgr(),
	}
}
