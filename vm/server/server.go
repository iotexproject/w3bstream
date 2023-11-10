package server

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/machinefi/w3bstream-mainnet/msg"
)

type Instance struct {
	conn *grpc.ClientConn
	resp *CreateResponse
}

func NewInstance(ctx context.Context, grpcAddr string, msgKey msg.MsgKey, executeBinary []byte, expParam string) (*Instance, error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial risc0 server")
	}
	cli := NewVmRuntimeClient(conn)

	req := &CreateRequest{
		Project:  string(msgKey),
		Content:  executeBinary,
		ExpParam: expParam,
	}
	resp, err := cli.Create(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create risc0 instance")
	}
	return &Instance{conn: conn, resp: resp}, nil
}

func (i *Instance) Execute(ctx context.Context, msg *msg.Msg) ([]byte, error) {
	req := &ExecuteRequest{
		Project: string(msg.Key()),
		Param:   msg.Data,
	}
	cli := NewVmRuntimeClient(i.conn)
	resp, err := cli.ExecuteOperator(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute risc0 instance")
	}
	return resp.Result, nil
}

func (i *Instance) Release() {
	i.conn.Close()
}
