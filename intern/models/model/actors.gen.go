// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActor = "actors"

// Actor mapped from table <actors>
type Actor struct {
	ID        string `gorm:"column:id;type:text;primaryKey" json:"id"`                 // Actor address.
	Code      string `gorm:"column:code;type:text;not null" json:"code"`               // Human readable identifier for the type of the actor.
	Head      string `gorm:"column:head;type:text;not null" json:"head"`               // CID of the root of the state tree for the actor.
	Nonce     int64  `gorm:"column:nonce;type:bigint;not null" json:"nonce"`           // The next actor nonce that is expected to appear on chain.
	Balance   string `gorm:"column:balance;type:text;not null" json:"balance,string"`  // Actor balance in attoFIL.
	StateRoot string `gorm:"column:state_root;type:text;primaryKey" json:"state_root"` // CID of the state root.
	Height    int64  `gorm:"column:height;type:bigint;primaryKey" json:"height"`       // Epoch when this actor was created or updated.
}

// TableName Actor's table name
func (*Actor) TableName() string {
	return TableNameActor
}
