package service

import (
	"context"
	"time"

	"github.com/Merteg/pl-health-service/config"
	"github.com/Merteg/pl-health-service/pkg/model"
	"github.com/Merteg/pl-health-service/proto"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

type Health struct {
	proto.UnimplementedHealthServiceServer
}

var mongoconfig = config.GetConfig().Mongo

func (s *Health) Push(c context.Context, req *proto.PushRequest) (*proto.PushResponse, error) {
	client := mongodbClient()
	healthcollection := client.Database(mongoconfig["dbname"]).Collection(mongoconfig["healthcollname"])
	var health []*proto.Health = req.GetHealth()

	for _, reqhealth := range health {
		id := reqhealth.TargetID

		var resphealth model.Health

		error := healthcollection.FindOne(context.TODO(), bson.M{"targetid": id}).Decode(&resphealth)
		if error != nil {
			log.Fatal().Err(error)
		}

		if resphealth.TargetID == "" {
			targetcollection := client.Database(mongoconfig["dbname"]).Collection(mongoconfig["targetscollname"])
			var target model.Target

			error := targetcollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&target)
			if error != nil {
				log.Fatal().Err(error)
			}
			resphealth.FromProto(reqhealth)

			_, err := healthcollection.InsertOne(c, resphealth)
			if err != nil {
				log.Fatal().Err(err)
			}
		} else {
			status.Error(codes.AlreadyExists, "This TargetID already exist")
		}
	}
	return &proto.PushResponse{}, nil
}

func (s *Health) Register(c context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	client := mongodbClient()

	collection := client.Database(mongoconfig["dbname"]).Collection(mongoconfig["targetscollname"])
	var target []*proto.Target = req.GetTarget()

	for _, reqtarget := range target {
		id := reqtarget.ID

		var resptarget model.Target

		error := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&resptarget)
		if error != nil {
			log.Fatal().Err(error)
		}

		if resptarget.Id == "" {

			resptarget.FromProto(reqtarget)

			_, err := collection.InsertOne(c, resptarget)
			if err != nil {
				log.Fatal().Err(err)
			}
		} else {
			status.Error(codes.AlreadyExists, "This TargetID already exist")
		}
	}
	return &proto.RegisterResponse{}, nil
}

func mongodbClient() *mongo.Client {
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
	return client
}
