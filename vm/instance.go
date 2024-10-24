package vm

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/vm/proto"
)

func create(ctx context.Context, conn *grpc.ClientConn, projectID uint64, executeBinary, expParam string) error {
	cli := proto.NewVmRuntimeClient(conn)

	req := &proto.CreateRequest{
		ProjectID: projectID,
		Content:   executeBinary,
		ExpParam:  expParam,
	}
	if _, err := cli.Create(ctx, req); err != nil {
		return errors.Wrap(err, "failed to create vm instance")
	}
	return nil
}

func execute(ctx context.Context, conn *grpc.ClientConn, task *task.Task) ([]byte, error) {
	ds := []string{}
	for _, d := range task.Payloads {
		ds = append(ds, string(d))
	}
	req := &proto.ExecuteRequest{
		ProjectID:          task.ProjectID,
		TaskID:             0,   // TODO
		ClientID:           "0", // TODO
		SequencerSignature: "",  // TODO
		Datas:              ds,
	}
	cli := proto.NewVmRuntimeClient(conn)
	resp, err := cli.Execute(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}
	return resp.Result, nil
}
