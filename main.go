package main

import (
	"fmt"
	"net"

	service "github.com/Merteg/pl-health-service/pkg"
	pl_health_service "github.com/Merteg/pl-health-service/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Error().Err(err).Msg("Problem creating the conection")
	} else {
		fmt.Println("start gRPC server at ", listener.Addr().String())
	}

	serv := grpc.NewServer()
	pl_health_service.RegisterHealthServiceServer(serv, &service.Server{})
	if err = serv.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Problem with services")
	} else {
		log.Printf("start gRPC server at %s", listener.Addr().String())
	}
}
