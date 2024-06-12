package vm

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/machinefi/sprout/task"
	"github.com/machinefi/sprout/vm/proto"
)

type instance struct {
	conn *grpc.ClientConn
}

func newInstance(ctx context.Context, projectID uint64, endpoint, executeBinary, expParam string) (*instance, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial vm server")
	}
	cli := proto.NewVmRuntimeClient(conn)

	req := &proto.CreateRequest{
		ProjectID: projectID,
		Content:   executeBinary,
		ExpParam:  expParam,
	}
	if _, err := cli.Create(ctx, req); err != nil {
		return nil, errors.Wrap(err, "failed to create vm instance")
	}
	return &instance{conn: conn}, nil
}

func (i *instance) execute(ctx context.Context, task *task.Task) ([]byte, error) {
	ds := []string{}
	for _, d := range task.Data {
		ds = append(ds, string(d))
	}
	req := &proto.ExecuteRequest{
		ProjectID:          task.ProjectID,
		TaskID:             task.ID,
		ClientID:           task.ClientID,
		SequencerSignature: task.Signature,
		Datas:              ds,
	}
	cli := proto.NewVmRuntimeClient(i.conn)
	resp, err := cli.ExecuteOperator(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}
	return resp.Result, nil
}

func (i *instance) release() {
	if err := i.conn.Close(); err != nil {
		slog.Error("failed to close grpc conn", "error", err)
	}
}
