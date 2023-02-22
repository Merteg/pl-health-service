package service

import (
	"context"
	"errors"

	"github.com/Merteg/pl-health-service/proto"
)

type Health struct {
	proto.UnimplementedHealthServiceServer
}

func (s *Health) Push(context.Context, *proto.PushRequest) (*proto.PushResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *Health) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}
