package model

import (
	"github.com/Merteg/pl-health-service/proto"
)

type Target struct {
	Id            string            `json:"_id,omitempty" validate:"required"`
	Metrics       []string          `json:"metrics,omitempty", bson:"metrics, omitempty"  `
	Counters      []string          `json:"counters,omitempty", bson:"counters, omitempty"  `
	TotalCounters []string          `json:"totalcounters,omitempty", bson:"totalcounters, omitempty"  validate:"required"`
	TargetType    string            `json:"targettype,omitempty", bson:"targettype, omitempty"  validate:"required"`
	Metadata      map[string]string `json:"metadata,omitempty", bson:"metadata, omitempty"  validate:"required"`
	Heartbeat     bool              `json:"hearbeat,omitempty", bson:"heartbeat, omitempty" validate:"required"`
}

func (t *Target) ConvertToSchema(target *proto.Target) {
	t.Metrics = target.GetMetrics()
	t.Counters = target.GetCounters()
	t.TotalCounters = target.GetTotalCounters()
	t.TargetType = target.GetTargetType()
	t.Metadata = target.GetMetadata()
	t.Heartbeat = target.GetHeartbeat()
}

func (t *Target) ConvertToProto() *proto.Target {
	return &proto.Target{
		Metrics:       t.Metrics,
		Counters:      t.Counters,
		TotalCounters: t.TotalCounters,
		TargetType:    t.TargetType,
		Metadata:      t.Metadata,
		Heartbeat:     t.Heartbeat,
	}
}
