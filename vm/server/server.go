package server

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/machinefi/sprout/message"
)

type Instance struct {
	conn *grpc.ClientConn
	resp *CreateResponse
}

func NewInstance(ctx context.Context, endpoint string, projectID uint64, executeBinary string, expParam string) (*Instance, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial vm server")
	}
	cli := NewVmRuntimeClient(conn)

	req := &CreateRequest{
		ProjectID: projectID,
		Content:   executeBinary,
		ExpParam:  expParam,
	}
	resp, err := cli.Create(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vm instance")
	}
	return &Instance{conn: conn, resp: resp}, nil
}

func (i *Instance) Execute(ctx context.Context, msg *message.Message) ([]byte, error) {
	req := &ExecuteRequest{
		ProjectID: msg.ProjectID,
		Param:     msg.Data,
	}
	cli := NewVmRuntimeClient(i.conn)
	resp, err := cli.ExecuteOperator(ctx, req)
	if err != nil {
		slog.Debug("request", "body", req)
		return nil, errors.Wrap(err, "failed to execute vm instance")
	}
	return resp.Result, nil
}

func (i *Instance) Release() {
	i.conn.Close()
}
