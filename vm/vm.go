package vm

import (
	"context"
	"encoding/hex"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/vm/proto"
)

type Handler struct {
	vmClients map[uint64]*grpc.ClientConn
}

func (r *Handler) Handle(task *task.Task, vmTypeID uint64, code string, expParam string) ([]byte, error) {
	conn, ok := r.vmClients[vmTypeID]
	if !ok {
		return nil, errors.Errorf("unsupported vm type id %d", vmTypeID)
	}
	cli := proto.NewVMClient(conn)

	bi, err := hex.DecodeString(code)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode code")
	}
	metadata, err := hex.DecodeString(expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode metadata")
	}
	if _, err := cli.NewProject(context.Background(), &proto.NewProjectRequest{
		ProjectID:      task.ProjectID,
		ProjectVersion: task.ProjectVersion,
		Binary:         bi,
		Metadata:       metadata,
	}); err != nil {
		slog.Error("failed to new project", "project_id", task.ProjectID, "err", err)
		return nil, errors.Wrap(err, "failed to create vm instance")
	}

	resp, err := cli.ExecuteTask(context.Background(), &proto.ExecuteTaskRequest{
		ProjectID: task.ProjectID,
		TaskID:    task.ID[:],
		Payloads:  task.Payloads,
	})
	if err != nil {
		slog.Error("failed to execute task", "project_id", task.ProjectID, "vm_type", vmTypeID,
			"task_id", task.ID, "binary", code, "payloads", task.Payloads, "err", err)
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}

	return resp.Result, nil
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
