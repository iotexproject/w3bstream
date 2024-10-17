package vm

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/vm/proto"
)

func create(ctx context.Context, conn *grpc.ClientConn, projectID uint64, executeBinary, expParam []byte) error {
	cli := proto.NewVMClient(conn)

	req := &proto.NewProjectRequest{
		ProjectID: projectID,
		Binary:    executeBinary,
		Metadata:  expParam,
	}
	if _, err := cli.NewProject(ctx, req); err != nil {
		return errors.Wrap(err, "failed to create vm instance")
	}
	return nil
}

func execute(ctx context.Context, conn *grpc.ClientConn, task *task.Task) ([]byte, error) {
	req := &proto.ExecuteTaskRequest{
		ProjectID: task.ProjectID,
		TaskID:    task.ID.Bytes(),
		Payloads:  task.Payloads,
	}
	cli := proto.NewVMClient(conn)
	resp, err := cli.ExecuteTask(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}
	return resp.Result, nil
}
