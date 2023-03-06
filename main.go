package main

import (
	"context"
	"net"
	"time"

	"github.com/Merteg/pl-health-service/config"
	service "github.com/Merteg/pl-health-service/pkg/service"
	"github.com/Merteg/pl-health-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var mongoconfig = config.GetConfig().Mongo

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoconfig["mongouri"]))
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

	targetCollection, err := client.Database(mongoconfig["dbname"]).ListCollectionNames(ctx, bson.M{"name": mongoconfig["targetscollname"]})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(targetCollection) == 0 {
		err = client.Database(mongoconfig["dbname"]).CreateCollection(ctx, mongoconfig["targetscollname"])
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + mongoconfig["targetscollname"])
	} else {
		log.Info().Msg("collection exist:" + mongoconfig["targetscollname"])
	}

	healthCollection, err := client.Database(mongoconfig["dbname"]).ListCollectionNames(ctx, bson.M{"name": mongoconfig["healthcollname"]})
	if err != nil {
		log.Fatal().Err(err)
	}
	if len(healthCollection) == 0 {
		err = client.Database(mongoconfig["dbname"]).CreateCollection(ctx, mongoconfig["healthcollname"])
		if err != nil {
			log.Fatal().Err(err)
		}
		log.Info().Msg("Collection created:" + mongoconfig["healthcollname"])
	} else {
		log.Info().Msg("Collection exist:" + mongoconfig["healthcollname"])
	}
}

func main() {
	listener, err := net.Listen("tcp", mongoconfig["port"])
	if err != nil {
		resp := "unable to listen on" + mongoconfig["port"]
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
