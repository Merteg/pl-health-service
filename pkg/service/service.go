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
func (s Health) Register(c context.Context, reqproto.RegisterRequest) (proto.RegisterResponse, error) {
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal().Err(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal().Err(err)
    }
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal().Err(err)
    }

    collection := client.Database(dbName).Collection(targetsCollName)
    var target []*proto.Target = req.GetTarget()

    for _, reqtarget := range target {
        id := reqtarget.ID

        var resptarget model.Target

        error := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&resptarget)
        if error != nil {
            log.Fatal().Err(err)
        }

        if resptarget.Id == "" {

            resptarget.FromProto(reqtarget)

            err := collection.InsertOne(ctx, resptarget)
            if err != nil {
                log.Fatal().Err(err)
            }
        } else {
            status.Error(codes.AlreadyExists, "This TargetID already exist")
        }
    }
    return &proto.RegisterResponse{}, nil
}
