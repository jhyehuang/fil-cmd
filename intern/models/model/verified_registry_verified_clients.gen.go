// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVerifiedRegistryVerifiedClient = "verified_registry_verified_clients"

// VerifiedRegistryVerifiedClient mapped from table <verified_registry_verified_clients>
type VerifiedRegistryVerifiedClient struct {
	Height    int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"`                   // Epoch at which this verified client state changed.
	StateRoot string  `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`             // CID of the parent state root at this epoch.
	Address   string  `gorm:"column:address;type:text;primaryKey" json:"address"`                   // Address of verified client this state change applies to.
	DataCap   float64 `gorm:"column:data_cap;type:numeric;not null" json:"data_cap"`                // DataCap of verified client at this state change.
	Event     string  `gorm:"column:event;type:verified_registry_event_type;not null" json:"event"` // Name of the event that occurred.
}

// TableName VerifiedRegistryVerifiedClient's table name
func (*VerifiedRegistryVerifiedClient) TableName() string {
	return TableNameVerifiedRegistryVerifiedClient
}
