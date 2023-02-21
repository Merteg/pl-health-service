package service

import (
	"context"
	"errors"

	"github.com/Merteg/pl-health-service/proto"
)

type Server struct {
	proto.UnimplementedHealthServiceServer
}

func (s *Server) Push(context.Context, *proto.PushRequest) (*proto.PushResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *Server) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}
