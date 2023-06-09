// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSurveyedMinerProtocol = "surveyed_miner_protocols"

// SurveyedMinerProtocol mapped from table <surveyed_miner_protocols>
type SurveyedMinerProtocol struct {
	ObservedAt time.Time `gorm:"column:observed_at;type:timestamp with time zone;primaryKey" json:"observed_at"` // Timestamp of the observation.
	MinerID    string    `gorm:"column:miner_id;type:text;primaryKey" json:"miner_id"`                           // Address (ActorID) of the miner.
	PeerID     *string   `gorm:"column:peer_id;type:text" json:"peer_id"`                                        // PeerID of the miner advertised in on-chain MinerInfo structure.
	Agent      *string   `gorm:"column:agent;type:text" json:"agent"`                                            // Agent string as reported by the peer.
	Protocols  *string   `gorm:"column:protocols;type:jsonb" json:"protocols"`                                   // List of supported protocol strings supported by the peer.
}

// TableName SurveyedMinerProtocol's table name
func (*SurveyedMinerProtocol) TableName() string {
	return TableNameSurveyedMinerProtocol
}
