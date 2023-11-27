package api

import (
	"context"
	"log/slog"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/sequencer"
)

type GrpcServer struct {
	seq *sequencer.Sequencer
	proto.UnimplementedSequencerServer
}

func NewGrpcServer(seq *sequencer.Sequencer) *GrpcServer {
	s := &GrpcServer{
		seq: seq,
	}
	return s
}

// this func will block caller
func (s *GrpcServer) Run(endpoint string) error {
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.Wrapf(err, "listen %s failed", endpoint)
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	proto.RegisterSequencerServer(grpcServer, s)

	if err := grpcServer.Serve(listen); err != nil {
		return errors.Wrap(err, "start grpc server failed")
	}
	return nil
}

func (s *GrpcServer) Fetch(ctx context.Context, req *proto.FetchRequest) (*proto.FetchResponse, error) {
	m, err := s.seq.Fetch(req.ProjectID)
	if err != nil {
		slog.Error("sequencer fetch failed", "error", err)
		return nil, err
	}
	ms := []*proto.Message{}
	if m != nil {
		ms = append(ms, m)
	}
	return &proto.FetchResponse{
		Messages: ms,
	}, nil
}

func (s *GrpcServer) Report(ctx context.Context, req *proto.ReportRequest) (*proto.ReportResponse, error) {
	if len(req.MessageIDs) == 0 {
		return nil, nil
	}
	if err := s.seq.UpdateMessageState(req.MessageIDs, req.State, req.Comment); err != nil {
		slog.Error("sequencer update message state failed", "error", err)
		return nil, err
	}
	return &proto.ReportResponse{}, nil
}
