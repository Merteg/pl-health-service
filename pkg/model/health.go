package model

import (
	"github.com/Merteg/pl-health-service/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Health struct {
	TargetID   string             `json:"targetid,omitempty" validate:"required" bson:"targetid, omitempty"`
	TargetType string             `json:"targettype,omitempty" validate:"required" bson:"targettype, omitempty"`
	Status     string             `json:"healthstatus,omitempty"  bson:"healthstatus, omitempty"`
	Counters   map[string]int32   `json:"counters,omitempty"  bson:"counters, omitempty"`
	Metrics    map[string]float64 `json:"metrics,omitempty"  bson:"metrics"`
	Hearthbeat bool               `json:"heartbeat,omitempty"  bson:"heartbeat, omitempty"`
	Messages   []*Message         `json:"messages,omitempty" bson:"messages, omitempty"`
	Timestamp  int64              `json:"timestamp,omitempty" validate:"required" bson:"timestamp, omitempty"`
}

type Message struct {
	Summary        string `json:"Summary" bson:"Summary"`
	Error          string `json:"Error" bson:"Error"`
	AffectedHealth bool   `json:"AffectedHealth" bson:"AffectedHealth"`
	Status         string `json:"Status" bson:"Status"`
}

func (h *Health) FromProto(health *proto.Health) {
	heartbeat := health.GetHearthbeat()
	if heartbeat != nil {
		h.Hearthbeat = heartbeat.Value
	}

	msgs := make([]*Message, 0, len(health.Messages))
	for _, message := range health.Messages {
		msgs = append(msgs, MessageFromProto(message))
	}
	h.Messages = msgs

	h.TargetID = health.GetTargetID()
	h.TargetType = health.GetTargetType()
	h.Status = health.GetStatus().String()
	h.Counters = health.GetCounters()
	h.Metrics = health.GetMetrics()
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
		Status:     proto.HealthStatus(proto.HealthStatus_value[h.Status]),
		Hearthbeat: &wrapperspb.BoolValue{Value: h.Hearthbeat},
		Counters:   h.Counters,
		Metrics:    h.Metrics,
		Messages:   msgs,
		Timestamp:  h.Timestamp,
	}
}

func MessageToProto(msg *Message) *proto.Message {
	var mesageResp proto.Message
	mesageResp.Summary = msg.Summary
	mesageResp.Error = msg.Error
	mesageResp.AffectHealth = msg.AffectedHealth
	mesageResp.Status = proto.HealthStatus(proto.HealthStatus_value[msg.Status])
	return &mesageResp
}

func MessageFromProto(msg *proto.Message) *Message {
	var mesageResp Message
	mesageResp.Summary = msg.Summary
	mesageResp.Error = msg.Error
	mesageResp.AffectedHealth = msg.AffectHealth
	mesageResp.Status = msg.Status.String()
	return &mesageResp
}
