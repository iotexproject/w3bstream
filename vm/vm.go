package vm

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/machinefi/sprout/task"
)

type Type string

const (
	Risc0  Type = "risc0"
	Halo2  Type = "halo2"
	ZKwasm Type = "zkwasm"
	Wasm   Type = "wasm"
)

type Handler struct {
	vmServerClients map[Type]*grpc.ClientConn
}

func (r *Handler) Handle(task *task.Task, vmtype Type, code string, expParam string) ([]byte, error) {
	conn, ok := r.vmServerClients[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	if err := create(context.Background(), conn, task.ProjectID, code, expParam); err != nil {
		return nil, errors.Wrap(err, "failed to create vm instance")
	}
	slog.Debug("create vm instance success", "vm_type", vmtype)

	res, err := execute(context.Background(), conn, task)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(vmServerEndpoints map[Type]string) (*Handler, error) {
	clients := map[Type]*grpc.ClientConn{}
	for t, e := range vmServerEndpoints {
		conn, err := grpc.NewClient(e, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, errors.Wrap(err, "failed to new grpc client")
		}
		clients[t] = conn
	}
	return &Handler{
		vmServerClients: clients,
	}, nil
}
