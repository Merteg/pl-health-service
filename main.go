package main

import (
	"net"

	service "github.com/Merteg/pl-health-service/pkg/service"
	"github.com/Merteg/pl-health-service/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal().Msg("unable to listen on localhost:8080")
	} else {
		log.Printf("start gRPC server at localhost:8080", listener)
	}

	serv := grpc.NewServer()
	proto.RegisterHealthServiceServer(serv, &service.Server{})
	if err = serv.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Problem with services")
	} else {
		log.Printf("start gRPC server at %s", listener.Addr().String())
	}
}
