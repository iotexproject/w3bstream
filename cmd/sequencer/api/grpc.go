package api

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/sequencer"
)

func NewGrpcServer(seq *sequencer.Sequencer) *GrpcServer {
	s := &GrpcServer{
		seq: seq,
	}

	return s
}

type GrpcServer struct {
	seq *sequencer.Sequencer
	proto.UnimplementedSequencerServiceServer
}

func (s *GrpcServer) Fetch(context.Context, *proto.FetchRequest) (*proto.FetchResponse, error) {
	return nil, nil
}
func (s *GrpcServer) Report(context.Context, *proto.ReportRequest) (*proto.ReportResponse, error) {
	return nil, nil
}

// this func will block caller
func (s *GrpcServer) Run(endpoint string) error {
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.Wrapf(err, "listen %s failed", endpoint)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSequencerServiceServer(grpcServer, s)

	if err := grpcServer.Serve(listen); err != nil {
		return errors.Wrap(err, "start grpc server failed")
	}
	return nil
}
