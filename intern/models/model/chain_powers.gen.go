// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameChainPower = "chain_powers"

// ChainPower mapped from table <chain_powers>
type ChainPower struct {
	StateRoot                  string  `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`                                        // CID of the parent state root.
	TotalRawBytesPower         float64 `gorm:"column:total_raw_bytes_power;type:numeric;not null" json:"total_raw_bytes_power"`                 // Total storage power in bytes in the network. Raw byte power is the size of a sector in bytes.
	TotalRawBytesCommitted     float64 `gorm:"column:total_raw_bytes_committed;type:numeric;not null" json:"total_raw_bytes_committed"`         // Total provably committed storage power in bytes. Raw byte power is the size of a sector in bytes.
	TotalQaBytesPower          float64 `gorm:"column:total_qa_bytes_power;type:numeric;not null" json:"total_qa_bytes_power"`                   // Total quality adjusted storage power in bytes in the network. Quality adjusted power is a weighted average of the quality of its space and it is based on the size, duration and quality of its deals.
	TotalQaBytesCommitted      float64 `gorm:"column:total_qa_bytes_committed;type:numeric;not null" json:"total_qa_bytes_committed"`           // Total provably committed, quality adjusted storage power in bytes. Quality adjusted power is a weighted average of the quality of its space and it is based on the size, duration and quality of its deals.
	TotalPledgeCollateral      float64 `gorm:"column:total_pledge_collateral;type:numeric;not null" json:"total_pledge_collateral"`             // Total locked FIL (attoFIL) miners have pledged as collateral in order to participate in the economy.
	QaSmoothedPositionEstimate float64 `gorm:"column:qa_smoothed_position_estimate;type:numeric;not null" json:"qa_smoothed_position_estimate"` // Total power smoothed position estimate - Alpha Beta Filter "position" (value) estimate in Q.128 format.
	QaSmoothedVelocityEstimate float64 `gorm:"column:qa_smoothed_velocity_estimate;type:numeric;not null" json:"qa_smoothed_velocity_estimate"` // Total power smoothed velocity estimate - Alpha Beta Filter "velocity" (rate of change of value) estimate in Q.128 format.
	MinerCount                 *int64  `gorm:"column:miner_count;type:bigint" json:"miner_count"`                                               // Total number of miners.
	ParticipatingMinerCount    *int64  `gorm:"column:participating_miner_count;type:bigint" json:"participating_miner_count"`                   // Total number of miners with power above the minimum miner threshold.
	Height                     int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"`                                              // Epoch this power summary applies to.
}

// TableName ChainPower's table name
func (*ChainPower) TableName() string {
	return TableNameChainPower
}
