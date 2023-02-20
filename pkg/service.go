package service

import (
	"context"
	"errors"

	pl_health_service "github.com/Merteg/pl-health-service/proto"
)

type Server struct {
	pl_health_service.UnimplementedHealthServiceServer
}

func (s *Server) Push(context.Context, *pl_health_service.PushRequest) (*pl_health_service.PushResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *Server) Register(context.Context, *pl_health_service.RegisterRequest) (*pl_health_service.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}
