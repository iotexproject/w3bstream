package vm

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iotexproject/w3bstream/task"
)

type Handler struct {
	vmClients map[uint64]*grpc.ClientConn
}

func (r *Handler) Handle(task *task.Task, vmTypeID uint64, code string, expParam string) ([]byte, error) {
	conn, ok := r.vmClients[vmTypeID]
	if !ok {
		return nil, errors.Errorf("unsupported vm type id %d", vmTypeID)
	}

	if err := create(context.Background(), conn, task.ProjectID, []byte(code), []byte(expParam)); err != nil {
		return nil, errors.Wrap(err, "failed to create vm instance")
	}
	slog.Debug("create vm instance success", "vm_type_id", vmTypeID)

	res, err := execute(context.Background(), conn, task)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(vmEndpoints map[uint64]string) (*Handler, error) {
	clients := map[uint64]*grpc.ClientConn{}
	for t, e := range vmEndpoints {
		conn, err := grpc.NewClient(e, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, errors.Wrap(err, "failed to new grpc client")
		}
		clients[t] = conn
	}
	return &Handler{
		vmClients: clients,
	}, nil
}
