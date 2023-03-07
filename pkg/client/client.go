package client

import (
	"math/rand"
	"time"

	"github.com/Merteg/pl-health-service/config"
	service "github.com/Merteg/pl-health-service/pkg/service"
	"github.com/Merteg/pl-health-service/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/rs/zerolog/log"
)

var mongoconfig = config.GetConfig().Mongo

func Testclient() {
	// Conecta a MongoDB
	client, ctx := service.MongodbClient()
	defer client.Disconnect(ctx)

	// Selecciona la colección de salud
	healthCollection := client.Database(mongoconfig["dbname"]).Collection(mongoconfig["healthcollname"])

	// Crear una lista de diccionarios con los valores predefinidos para cada request preparado
	preparedRequests := []map[string]interface{}{
		{"TargetID": "id1", "TargetType": "type1", "Metrics": map[string]float64{"metric1": 1.23, "metric2": 4.56}},
		{"TargetID": "id2", "TargetType": "type2", "Metrics": map[string]float64{"metric1": 2.34, "metric2": 5.67}},
		{"TargetID": "id3", "TargetType": "type3", "Metrics": map[string]float64{"metric1": 3.45, "metric2": 6.78}},

		{"TargetID": "id4", "TargetType": "type4", "Metrics": map[string]float64{"metric1": rand.Float64()*9 + 1, "metric2": rand.Float64()*9 + 1}},
		{"TargetID": "id5", "TargetType": "type5", "Metrics": map[string]float64{"metric1": rand.Float64()*9 + 1, "metric2": rand.Float64()*9 + 1}},
		{"TargetID": "id6", "TargetType": "type6", "Metrics": map[string]float64{"metric1": rand.Float64()*9 + 1, "metric2": rand.Float64()*9 + 1}},
	}

	// Enviar 3 request preparados usando los valores de la lista de diccionarios
	for _, req := range preparedRequests {
		affectHealth := rand.Intn(2) == 1
		var status proto.HealthStatus
		switch rand.Intn(4) {
		case 0:
			status = proto.HealthStatus_UNKNOWN
		case 1:
			status = proto.HealthStatus_HEALTHY
		case 2:
			status = proto.HealthStatus_DEGRADE
		case 3:
			status = proto.HealthStatus_UNHEALTHY
		}

		message := &proto.Message{
			Summary:      "summary",
			Error:        "error",
			AffectHealth: affectHealth,
			Status:       status,
		}

		health := &proto.Health{
			TargetID:   req["TargetID"].(string),
			TargetType: req["TargetType"].(string),
			Status:     message.Status,
			Counters:   map[string]int32{"counter1": 10, "counter2": 20},
			Metrics:    req["Metrics"].(map[string]float64),
			Hearthbeat: &wrapperspb.BoolValue{Value: true},
			Messages:   []*proto.Message{message},
			Timestamp:  time.Now().Unix(),
		}

		_, err := healthCollection.InsertOne(ctx, health)
		if err != nil {
			log.Fatal().Err(err)
		}
	}

}
