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

	targetCollection, err := client.Database(dbName).ListCollectionNames(ctx, bson.M{"name": targetsCollName})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(targetCollection) == 0 {
		err = client.Database(dbName).CreateCollection(ctx, targetsCollName)
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + targetsCollName)
	} else {
		log.Info().Msg("collection exist:" + targetsCollName)
	}

	healthCollection, err := client.Database(dbName).ListCollectionNames(ctx, bson.M{"name": healthCollName})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(healthCollection) == 0 {
		err = client.Database(dbName).CreateCollection(ctx, healthCollName)
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + healthCollName)
	} else {
		log.Info().Msg("Collection exist:" + healthCollName)
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
