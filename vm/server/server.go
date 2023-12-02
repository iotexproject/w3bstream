package server

import (
	"context"
	"log/slog"

	msg "github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/vm/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Instance struct {
	conn *grpc.ClientConn
	resp *proto.CreateResponse
}

func NewInstance(ctx context.Context, endpoint string, projectID uint64, executeBinary string, expParam string) (*Instance, error) {
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
	resp, err := cli.Create(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vm instance")
	}
	return &Instance{conn: conn, resp: resp}, nil
}

func (i *Instance) Execute(ctx context.Context, msgs []*msg.Message) ([]byte, error) {
	datas := []string{}
	for _, m := range msgs {
		datas = append(datas, m.Data)
	}
	req := &proto.ExecuteRequest{
		ProjectID: msgs[0].ProjectID,
		Datas:     datas,
	}
	cli := proto.NewVmRuntimeClient(i.conn)
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
