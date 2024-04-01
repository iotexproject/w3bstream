package server_test

import (
	"context"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/machinefi/sprout/vm/proto"
	"github.com/machinefi/sprout/vm/server"
)

type MockClient struct{}

func (*MockClient) Create(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption) (*proto.CreateResponse, error) {
	return nil, nil
}

func (*MockClient) ExecuteOperator(ctx context.Context, in *proto.ExecuteRequest, opts ...grpc.CallOption) (*proto.ExecuteResponse, error) {
	return nil, nil
}

func TestNewInstance(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToDialGRPC", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(grpc.Dial, nil, errors.New(t.Name()))
		_, err := server.NewInstance(context.Background(), "any", 100, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToInvokeGRPCCreate", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(grpc.Dial, &grpc.ClientConn{}, nil)
		p = p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p = p.ApplyMethodReturn(&MockClient{}, "Create", nil, errors.New(t.Name()))

		_, err := server.NewInstance(context.Background(), "any", 100, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(grpc.Dial, &grpc.ClientConn{}, nil)
		p = p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p = p.ApplyMethodReturn(&MockClient{}, "Create", &proto.CreateResponse{}, nil)
		p = p.ApplyMethodReturn(&grpc.ClientConn{}, "Close", nil)

		i, err := server.NewInstance(context.Background(), "any", 100, "any", "any")
		r.NoError(err, t.Name())
		r.NotNil(i)
		i.Release()
	})
}

func TestInstance_Execute(t *testing.T) {
	r := require.New(t)
	i := &server.Instance{}

	t.Run("FailedToCallGRPCExecuteOperator", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p = p.ApplyMethodReturn(&MockClient{}, "ExecuteOperator", nil, errors.New(t.Name()))

		_, err := i.Execute(context.Background(), 1, 1, "any", "any", [][]byte{})
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p = p.ApplyMethodReturn(&MockClient{}, "ExecuteOperator", &proto.ExecuteResponse{Result: []byte("any")}, nil)

		res, err := i.Execute(context.Background(), 1, 1, "any", "any", [][]byte{})
		r.NoError(err, t.Name())
		r.Equal(res, []byte("any"))
	})
}
