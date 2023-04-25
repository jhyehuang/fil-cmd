// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMinerSectorDeal = "miner_sector_deals"

// MinerSectorDeal mapped from table <miner_sector_deals>
type MinerSectorDeal struct {
	MinerID  string `gorm:"column:miner_id;type:text;primaryKey" json:"miner_id"`     // Address of the miner the deal is with.
	SectorID int64  `gorm:"column:sector_id;type:bigint;primaryKey" json:"sector_id"` // Numeric identifier of the sector the deal is for.
	DealID   int64  `gorm:"column:deal_id;type:bigint;primaryKey" json:"deal_id"`     // Numeric identifier for the deal.
	Height   int64  `gorm:"column:height;type:bigint;primaryKey" json:"height"`       // Epoch at which this deal was added/updated.
}

// TableName MinerSectorDeal's table name
func (*MinerSectorDeal) TableName() string {
	return TableNameMinerSectorDeal
}
