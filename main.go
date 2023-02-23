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

const (
    port = "localhost:8080"
    mongoURI = "mongodb+srv://admin:admin@pl-health-service.s25udti.mongodb.net/test"
    dbName = "pl-health-service"
    targetsCollName = "targets"
    healthCollName = "health"
)

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

	targetCollection, err := client.Database(healthdb).ListCollectionNames(ctx, bson.M{"name": coll1})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(targetCollection) == 0 {
		err = client.Database(healthdb).CreateCollection(ctx, coll1)
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + coll1)
	} else {
		log.Info().Msg("collection exist:" + coll1)
	}
	healthCollection, err := client.Database(healthdb).ListCollectionNames(ctx, bson.M{"name": coll2})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(healthCollection) == 0 {
		err = client.Database(healthdb).CreateCollection(ctx, coll2)
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + coll2)
	} else {
		log.Info().Msg("Collection exist:" + coll2)
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
