package vm

import (
	"context"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/vm/proto"
)

type MockClient struct{}

func (*MockClient) Create(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption) (*proto.CreateResponse, error) {
	return nil, nil
}

func (*MockClient) Execute(ctx context.Context, in *proto.ExecuteRequest, opts ...grpc.CallOption) (*proto.ExecuteResponse, error) {
	return nil, nil
}

func TestCreateInstance(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToInvokeGRPCCreate", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(grpc.Dial, &grpc.ClientConn{}, nil)
		p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p.ApplyMethodReturn(&MockClient{}, "Create", nil, errors.New(t.Name()))

		err := create(context.Background(), nil, 100, "any", "any")
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(grpc.Dial, &grpc.ClientConn{}, nil)
		p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p.ApplyMethodReturn(&MockClient{}, "Create", &proto.CreateResponse{}, nil)
		p.ApplyMethodReturn(&grpc.ClientConn{}, "Close", nil)

		err := create(context.Background(), nil, 100, "any", "any")
		r.NoError(err, t.Name())
	})
}

func TestExecuteInstance(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToCallGRPCExecute", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p.ApplyMethodReturn(&MockClient{}, "Execute", nil, errors.New(t.Name()))

		_, err := execute(context.Background(), nil, &task.Task{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(proto.NewVmRuntimeClient, &MockClient{})
		p.ApplyMethodReturn(&MockClient{}, "Execute", &proto.ExecuteResponse{Result: []byte("any")}, nil)

		res, err := execute(context.Background(), nil, &task.Task{Payloads: [][]byte{[]byte("data")}})
		r.NoError(err, t.Name())
		r.Equal(res, []byte("any"))
	})
}
