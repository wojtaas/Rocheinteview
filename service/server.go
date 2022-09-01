package service

import (
	"context"
	"google.golang.org/grpc"
	. "rocheinteview/grpc/proto"
)

type Server struct {
	service Service
}

func NewServer(sv Service) *Server {
	return &Server{service: sv}
}

func (s *Server) Register(server *grpc.Server) {
	RegisterGRPCPingServer(server, s)
}

func (s *Server) Ping(ctx context.Context, request *PingRequest) (*PingResponse, error) {
	pingOutput := s.service.Ping(request.Message)
	return &PingResponse{
		Echo:      pingOutput.Echo,
		Timestamp: pingOutput.Timestamp,
		Env:       pingOutput.Env,
		Version:   pingOutput.Version,
	}, nil
}
