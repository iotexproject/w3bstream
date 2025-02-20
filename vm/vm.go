package vm

import (
	"context"
	"encoding/hex"
	"log/slog"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/util/filefetcher"
	"github.com/iotexproject/w3bstream/vm/proto"
)

type Handler struct {
	vmClients map[uint64]*grpc.ClientConn
}

func (r *Handler) Handle(tasks []*task.Task, projectConfig *project.Config) ([]byte, error) {
	task := tasks[0]
	// TODO: load binary before being stored in db
	now := time.Now()
	bi, err := decodeBinary(projectConfig.Code)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode code")
	}
	slog.Info("fetch circuit code success", "time duration", time.Since(now).String())
	now = time.Now()
	metadata, err := decodeBinary(projectConfig.Metadata)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode metadata")
	}
	slog.Info("fetch circuit metadata success", "time duration", time.Since(now).String())
	taskPayload, err := loadPayload(tasks, projectConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load payload")
	}
	conn, ok := r.vmClients[projectConfig.VMTypeID]
	if !ok {
		return nil, errors.Errorf("unsupported vm type id %d", projectConfig.VMTypeID)
	}
	cli := proto.NewVMClient(conn)
	if _, err := cli.NewProject(context.Background(), &proto.NewProjectRequest{
		ProjectID:      task.ProjectID.String(),
		ProjectVersion: task.ProjectVersion,
		Binary:         bi,
		Metadata:       metadata,
	}); err != nil {
		slog.Error("failed to new project", "project_id", tasks[0].ProjectID, "err", err)
		return nil, errors.Wrap(err, "failed to create vm instance")
	}

	resp, err := cli.ExecuteTask(context.Background(), &proto.ExecuteTaskRequest{
		ProjectID:      task.ProjectID.String(),
		ProjectVersion: task.ProjectVersion,
		TaskID:         task.ID[:],
		Payloads:       [][]byte{taskPayload},
	})
	if err != nil {
		slog.Error("failed to execute task", "project_id", task.ProjectID, "vm_type", projectConfig.VMTypeID,
			"task_id", task.ID, "binary", projectConfig.Code, "payloads", task.Payload, "err", err)
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}
	slog.Info("proof", "proof", hexutil.Encode(resp.Result))

	return resp.Result, nil
}

// TODO: validate fetched file with hash
func decodeBinary(b string) ([]byte, error) {
	if strings.Contains(b, "http") ||
		strings.Contains(b, "ipfs") {
		fd := filefetcher.Filedescriptor{Uri: b}
		skipHash := true
		return fd.FetchFile(skipHash)
	}
	return hex.DecodeString(b)
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
