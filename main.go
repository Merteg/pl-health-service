package main

import (
	"context"
	"fmt"
	"net"
	"time"

	service "github.com/Merteg/pl-health-service/pkg"
	pl_health_service "github.com/Merteg/pl-health-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@pl-health-service.s25udti.mongodb.net/test"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}

	targetCollection, err := client.Database("pl-health-service").ListCollectionNames(context.TODO(), bson.M{"name": "target"})
	if err != nil {
		fmt.Println(err)
	}
	if len(targetCollection) == 0 {
		_ = client.Database("pl-health-service").CreateCollection(context.TODO(), "target")
		fmt.Println("target collection created and ready")
	} else {
		fmt.Println("target collection ready")
	}
	healthCollection, err := client.Database("pl-health-service").ListCollectionNames(context.TODO(), bson.M{"name": "health"})
	if err != nil {
		fmt.Println(err)
	}
	if len(healthCollection) == 0 {
		_ = client.Database("pl-health-service").CreateCollection(context.TODO(), "health")
		fmt.Println("health collection created and ready")
	} else {
		fmt.Println("health collection ready")
	}
}

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
