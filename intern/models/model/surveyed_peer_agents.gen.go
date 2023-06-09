// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSurveyedPeerAgent = "surveyed_peer_agents"

// SurveyedPeerAgent mapped from table <surveyed_peer_agents>
type SurveyedPeerAgent struct {
	SurveyerPeerID  string    `gorm:"column:surveyer_peer_id;type:text;primaryKey" json:"surveyer_peer_id"`           // Peer ID of the node performing the survey.
	ObservedAt      time.Time `gorm:"column:observed_at;type:timestamp with time zone;primaryKey" json:"observed_at"` // Timestamp of the observation.
	RawAgent        string    `gorm:"column:raw_agent;type:text;primaryKey" json:"raw_agent"`                         // Unprocessed agent string as reported by a peer.
	NormalizedAgent string    `gorm:"column:normalized_agent;type:text;not null" json:"normalized_agent"`             // Agent string normalized to a software name with major and minor version.
	Count           int64     `gorm:"column:count;type:bigint;not null" json:"count"`                                 // Number of peers that reported the same raw agent.
}

// TableName SurveyedPeerAgent's table name
func (*SurveyedPeerAgent) TableName() string {
	return TableNameSurveyedPeerAgent
}
