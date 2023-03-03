package model

import (
	"strconv"

	"github.com/Merteg/pl-health-service/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Health struct {
	Id         string             `json:"_id,omitempty", bson: "_id, omitempty"`
	TargetID   string             `json:"targetid,omitempty" validate:"required", bson:"targetid, omitempty" validate:"required"`
	TargetType string             `json:"targettype,omitempty" validate:"required", bson:"targettype, omitempty" validate:"required"`
	Status     string             `json:"healthstatus,omitempty",  bson:"healthstatus, omitempty"`
	Counters   map[string]int32   `json:"counters,omitempty",  bson:"counters, omitempty"`
	Metrics    map[string]float64 `json:"metrics,omitempty",  bson:"metrics"`
	Hearthbeat bool               `json:"heartbeat,omitempty",  bson:"heartbeat, omitempty"`
	Messages   []*Message         `json:"messages,omitempty", bson:"messages, omitempty"`
	Timestamp  int64              `json:"timestamp,omitempty" validate:"required", bson:"timestamp, omitempty" validate:"required"`
}
type Message struct {
	Summary        string             `json:"Summary"`
	Error          string             `json:"Error"`
	AffectedHealth bool               `json:"AffectedHealth"`
	Status         proto.HealthStatus `json:"Status"`
}

func (h *Health) FromProto(health *proto.Health) {
	h.TargetID = health.GetTargetID()
	h.TargetType = health.GetTargetType()
	h.Status = health.GetStatus().String()
	h.Hearthbeat = health.GetHearthbeat()
	h.Counters = health.GetCounters()
	h.Metrics = health.GetMetrics()
	h.Messages = health.GetMessages()
	h.Timestamp = health.GetTimestamp()
}

func (h *Health) ToProto() *proto.Health {
	msgs := make([]*proto.Message, 0, len(h.Messages))
	for _, message := range h.Messages {
		msgs = append(msgs, MessageToProto(message))
	}
	return &proto.Health{
		TargetID:   h.TargetID,
		TargetType: h.TargetType,
		Status:     StatusToEnum(h.Status),
		Hearthbeat: &wrapperspb.BoolValue{Value: h.Hearthbeat},
		Counters:   h.Counters,
		Metrics:    h.Metrics,
		Messages:   msgs,
		Timestamp:  h.Timestamp,
	}
}

func StatusToEnum(status string) int32 {
	i64, err := strconv.ParseInt(status, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i64)
}

func MessageToProto(msg *Message) *proto.Message {

}
