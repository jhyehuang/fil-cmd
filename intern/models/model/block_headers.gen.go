// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBlockHeader = "block_headers"

// BlockHeader mapped from table <block_headers>
type BlockHeader struct {
	Cid             string `gorm:"column:cid;type:text;primaryKey" json:"cid"`                           // CID of the block.
	ParentWeight    string `gorm:"column:parent_weight;type:text;not null" json:"parent_weight"`         // Aggregate chain weight of the block's parent set.
	ParentStateRoot string `gorm:"column:parent_state_root;type:text;not null" json:"parent_state_root"` // CID of the block's parent state root.
	Height          int64  `gorm:"column:height;type:bigint;primaryKey" json:"height"`                   // Epoch when this block was mined.
	Miner           string `gorm:"column:miner;type:text;not null" json:"miner"`                         // Address of the miner who mined this block.
	Timestamp       int64  `gorm:"column:timestamp;type:bigint;not null" json:"timestamp"`               // Time the block was mined in Unix time, the number of seconds elapsed since January 1, 1970 UTC.
	WinCount        *int64 `gorm:"column:win_count;type:bigint" json:"win_count"`                        // Number of reward units won in this block.
	ParentBaseFee   string `gorm:"column:parent_base_fee;type:text;not null" json:"parent_base_fee"`     // The base fee after executing the parent tipset.
	ForkSignaling   int64  `gorm:"column:fork_signaling;type:bigint;not null" json:"fork_signaling"`     // Flag used as part of signaling forks.
}

// TableName BlockHeader's table name
func (*BlockHeader) TableName() string {
	return TableNameBlockHeader
}
