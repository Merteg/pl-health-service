package main

import (
	"net"

	service "github.com/Merteg/pl-health-service/pkg/service"
	"github.com/Merteg/pl-health-service/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const port string = "localhost:8080"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		resp := "unable to listen on" + port
		log.Fatal().Msg(resp)
	}

	serv := grpc.NewServer()
	proto.RegisterHealthServiceServer(serv, &service.Health{})
	if err = serv.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Problem with services")
	} else {
		str := "start gRPC server at " + listener.Addr().String()
		log.Info().Msg(str)
	}
}
