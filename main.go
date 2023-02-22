package main

import (
	"context"
	"net"
	"time"

	service "github.com/Merteg/pl-health-service/pkg/service"
	"github.com/Merteg/pl-health-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const port string = "localhost:8080"
const mongoURI string = "mongodb+srv://admin:admin@pl-health-service.s25udti.mongodb.net/test"

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal().Err(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	targetCollection, err := client.Database("pl-health-service").ListCollectionNames(context.TODO(), bson.M{"name": "target"})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(targetCollection) == 0 {
		_ = client.Database("pl-health-service").CreateCollection(context.TODO(), "target")
		log.Info().Msg("target collection created and ready")
	} else {
		log.Info().Msg("target collection ready")
	}
	healthCollection, err := client.Database("pl-health-service").ListCollectionNames(context.TODO(), bson.M{"name": "health"})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(healthCollection) == 0 {
		_ = client.Database("pl-health-service").CreateCollection(context.TODO(), "health")
		log.Info().Msg("health collection created and ready")
	} else {
		log.Info().Msg("health collection ready")
	}
}

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
