package model

import (
	"github.com/Merteg/pl-health-service/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Health struct {
	Id         string                `json:"_id,omitempty", bson: "_id, omitempty"`
	TargetID   string                `json:"targetid,omitempty" validate:"required", bson:"targetid, omitempty" validate:"required"`
	TargetType string                `json:"targettype,omitempty" validate:"required", bson:"targettype, omitempty" validate:"required"`
	Status     string                `json:"healthstatus,omitempty",  bson:"healthstatus, omitempty"`
	Counters   map[string]int32      `json:"counters,omitempty",  bson:"counters, omitempty"`
	Metrics    map[string]float64    `json:"metrics,omitempty",  bson:"metrics"`
	Heartbeat  *wrapperspb.BoolValue `json:"heartbeat,omitempty",  bson:"heartbeat, omitempty"`
	Messages   []*proto.Message      `json:"messages,omitempty", bson:"messages, omitempty"`
	Timestamp  int64                 `json:"timestamp,omitempty" validate:"required", bson:"timestamp, omitempty" validate:"required"`
}

func (h *Health) ConvertToSchema(health *proto.Health) {
	h.TargetID = health.GetTargetID()
	h.TargetType = health.GetTargetType()
	h.Status = health.GetStatus().String()
	h.Heartbeat = health.GetHearthbeat()
	h.Counters = health.GetCounters()
	h.Metrics = health.GetMetrics()
	h.Messages = health.GetMessages()
	h.Timestamp = health.GetTimestamp()
}

func (h *Health) ConvertToProto() *proto.Target {
	return &proto.Target{
		TargetID:   h.TargetID,
		TargetType: h.TargetType,
		Status:     h.Status,
		Heartbeat:  h.Heartbeat,
		Counters:   h.Counters,
		Metrics:    h.Metrics,
		Messages:   h.Messages,
		Timestamp:  h.Timestamp,
	}
}
