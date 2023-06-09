// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMinerBeneficiary = "miner_beneficiaries"

// MinerBeneficiary mapped from table <miner_beneficiaries>
type MinerBeneficiary struct {
	Height                int64    `gorm:"column:height;type:bigint;primaryKey" json:"height"`
	StateRoot             string   `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`
	MinerID               string   `gorm:"column:miner_id;type:text;primaryKey" json:"miner_id"`
	Beneficiary           string   `gorm:"column:beneficiary;type:text;not null" json:"beneficiary"`
	Quota                 float64  `gorm:"column:quota;type:numeric;not null" json:"quota"`
	UsedQuota             float64  `gorm:"column:used_quota;type:numeric;not null" json:"used_quota"`
	Expiration            int64    `gorm:"column:expiration;type:bigint;not null" json:"expiration"`
	NewBeneficiary        *string  `gorm:"column:new_beneficiary;type:text" json:"new_beneficiary"`
	NewQuota              *float64 `gorm:"column:new_quota;type:numeric" json:"new_quota"`
	NewExpiration         *int64   `gorm:"column:new_expiration;type:bigint" json:"new_expiration"`
	ApprovedByBeneficiary *bool    `gorm:"column:approved_by_beneficiary;type:boolean" json:"approved_by_beneficiary"`
	ApprovedByNominee     *bool    `gorm:"column:approved_by_nominee;type:boolean" json:"approved_by_nominee"`
}

// TableName MinerBeneficiary's table name
func (*MinerBeneficiary) TableName() string {
	return TableNameMinerBeneficiary
}
