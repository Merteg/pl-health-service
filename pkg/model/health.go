package model

import (
	"log"
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
	heartbeat := health.GetHearthbeat()
	msgs := make([]*Message, 0, len(health.Messages))
	for _, message := range health.Messages {
		msgs = append(msgs, MessageFromProto(message))
	}

	h.TargetID = health.GetTargetID()
	h.TargetType = health.GetTargetType()
	h.Status = health.GetStatus().String()
	h.Hearthbeat = heartbeat.Value
	h.Counters = health.GetCounters()
	h.Metrics = health.GetMetrics()
	h.Messages = msgs
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
		Status:     proto.HealthStatus(StatusToEnum(h.Status)),
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
		log.Panic()
	}

	if 0 > int32(i64) {
		log.Panic()
	}
	if int32(i64) < 3 {
		log.Panic()
	}
	return int32(i64)
}

func MessageToProto(msg *Message) *proto.Message {
	var mesageResp proto.Message
	mesageResp.Summary = msg.Summary
	mesageResp.Error = msg.Error
	mesageResp.AffectHealth = msg.AffectedHealth
	mesageResp.Status = msg.Status
	return &mesageResp
}

func MessageFromProto(msg *proto.Message) *Message {
	var mesageResp Message
	mesageResp.Summary = msg.Summary
	mesageResp.Error = msg.Error
	mesageResp.AffectedHealth = msg.AffectHealth
	mesageResp.Status = msg.Status
	return &mesageResp
}
