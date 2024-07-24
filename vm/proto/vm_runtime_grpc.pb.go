// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VmRuntimeClient is the client API for VmRuntime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VmRuntimeClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type vmRuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewVmRuntimeClient(cc grpc.ClientConnInterface) VmRuntimeClient {
	return &vmRuntimeClient{cc}
}

func (c *vmRuntimeClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/vm_runtime.VmRuntime/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vmRuntimeClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/vm_runtime.VmRuntime/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VmRuntimeServer is the server API for VmRuntime service.
// All implementations must embed UnimplementedVmRuntimeServer
// for forward compatibility
type VmRuntimeServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	mustEmbedUnimplementedVmRuntimeServer()
}

// UnimplementedVmRuntimeServer must be embedded to have forward compatible implementations.
type UnimplementedVmRuntimeServer struct {
}

func (UnimplementedVmRuntimeServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVmRuntimeServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedVmRuntimeServer) mustEmbedUnimplementedVmRuntimeServer() {}

// UnsafeVmRuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VmRuntimeServer will
// result in compilation errors.
type UnsafeVmRuntimeServer interface {
	mustEmbedUnimplementedVmRuntimeServer()
}

func RegisterVmRuntimeServer(s grpc.ServiceRegistrar, srv VmRuntimeServer) {
	s.RegisterService(&VmRuntime_ServiceDesc, srv)
}

func _VmRuntime_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VmRuntimeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vm_runtime.VmRuntime/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VmRuntimeServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VmRuntime_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VmRuntimeServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vm_runtime.VmRuntime/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VmRuntimeServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VmRuntime_ServiceDesc is the grpc.ServiceDesc for VmRuntime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VmRuntime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vm_runtime.VmRuntime",
	HandlerType: (*VmRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VmRuntime_Create_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _VmRuntime_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vm_runtime.proto",
}
