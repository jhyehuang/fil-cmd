// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMinerPreCommitInfo = "miner_pre_commit_infos"

// MinerPreCommitInfo mapped from table <miner_pre_commit_infos>
type MinerPreCommitInfo struct {
	MinerID                string  `gorm:"column:miner_id;type:text;primaryKey" json:"miner_id"`                          // Address of the miner who owns the sector.
	SectorID               int64   `gorm:"column:sector_id;type:bigint;primaryKey" json:"sector_id"`                      // Numeric identifier for the sector.
	StateRoot              string  `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`                      // CID of the parent state root at this epoch.
	SealedCid              string  `gorm:"column:sealed_cid;type:text;not null" json:"sealed_cid"`                        // CID of the sealed sector.
	SealRandEpoch          *int64  `gorm:"column:seal_rand_epoch;type:bigint" json:"seal_rand_epoch"`                     // Seal challenge epoch. Epoch at which randomness should be drawn to tie Proof-of-Replication to a chain.
	ExpirationEpoch        *int64  `gorm:"column:expiration_epoch;type:bigint" json:"expiration_epoch"`                   // Epoch this sector expires.
	PreCommitDeposit       float64 `gorm:"column:pre_commit_deposit;type:numeric;not null" json:"pre_commit_deposit"`     // Amount of FIL (in attoFIL) used as a PreCommit deposit. If the Sector is not ProveCommitted on time, this deposit is removed and burned.
	PreCommitEpoch         *int64  `gorm:"column:pre_commit_epoch;type:bigint" json:"pre_commit_epoch"`                   // Epoch this PreCommit was created.
	DealWeight             float64 `gorm:"column:deal_weight;type:numeric;not null" json:"deal_weight"`                   // Total space*time of submitted deals.
	VerifiedDealWeight     float64 `gorm:"column:verified_deal_weight;type:numeric;not null" json:"verified_deal_weight"` // Total space*time of submitted verified deals.
	IsReplaceCapacity      *bool   `gorm:"column:is_replace_capacity;type:boolean" json:"is_replace_capacity"`            // Whether to replace a "committed capacity" no-deal sector (requires non-empty DealIDs).
	ReplaceSectorDeadline  *int64  `gorm:"column:replace_sector_deadline;type:bigint" json:"replace_sector_deadline"`     // The deadline location of the sector to replace.
	ReplaceSectorPartition *int64  `gorm:"column:replace_sector_partition;type:bigint" json:"replace_sector_partition"`   // The partition location of the sector to replace.
	ReplaceSectorNumber    *int64  `gorm:"column:replace_sector_number;type:bigint" json:"replace_sector_number"`         // ID of the committed capacity sector to replace.
	Height                 int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"`                            // Epoch this PreCommit information was added/changed.
}

// TableName MinerPreCommitInfo's table name
func (*MinerPreCommitInfo) TableName() string {
	return TableNameMinerPreCommitInfo
}
