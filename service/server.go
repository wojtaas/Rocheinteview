package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	. "rocheinteview/grpc/pb-go"
)

type server struct {
	service Service
}

func NewServer(sv Service) *server {
	return &server{service: sv}
}

// Register func is used to register GRPC server.
func (s *server) Register(server *grpc.Server) {
	RegisterGRPCPingServer(server, s)
}

// Ping func is used to recevie GRPC message, and response with additional parameters - works similar like service/ping func.
func (s *server) Ping(ctx context.Context, request *PingRequest) (*PingResponse, error) {
	pingOutput, err := s.service.Ping(request.Message)
	if err != nil {
		return &PingResponse{}, errors.New(fmt.Sprintf("s.service.Ping(): %d", err))
	}

	return &PingResponse{
		Echo:      pingOutput.Echo,
		Timestamp: pingOutput.Timestamp,
		Env:       pingOutput.Env,
		Version:   pingOutput.Version,
	}, nil
}
