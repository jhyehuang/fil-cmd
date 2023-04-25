// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMarketDealProposal = "market_deal_proposals"

// MarketDealProposal mapped from table <market_deal_proposals>
type MarketDealProposal struct {
	DealID               int64   `gorm:"column:deal_id;type:bigint;primaryKey" json:"deal_id"`                       // Identifier for the deal.
	StateRoot            string  `gorm:"column:state_root;type:text;not null" json:"state_root"`                     // CID of the parent state root for this deal.
	PieceCid             string  `gorm:"column:piece_cid;type:text;not null" json:"piece_cid"`                       // CID of a sector piece. A Piece is an object that represents a whole or part of a File.
	PaddedPieceSize      int64   `gorm:"column:padded_piece_size;type:bigint;not null" json:"padded_piece_size"`     // The piece size in bytes with padding.
	UnpaddedPieceSize    int64   `gorm:"column:unpadded_piece_size;type:bigint;not null" json:"unpadded_piece_size"` // The piece size in bytes without padding.
	IsVerified           bool    `gorm:"column:is_verified;type:boolean;not null" json:"is_verified"`                // Deal is with a verified provider.
	ClientID             string  `gorm:"column:client_id;type:text;not null" json:"client_id"`                       // Address of the actor proposing the deal.
	ProviderID           string  `gorm:"column:provider_id;type:text;not null" json:"provider_id"`                   // Address of the actor providing the services.
	StartEpoch           int64   `gorm:"column:start_epoch;type:bigint;not null" json:"start_epoch"`                 // The epoch at which this deal with begin. Storage deal must appear in a sealed (proven) sector no later than start_epoch, otherwise it is invalid.
	EndEpoch             int64   `gorm:"column:end_epoch;type:bigint;not null" json:"end_epoch"`                     // The epoch at which this deal with end.
	SlashedEpoch         *int64  `gorm:"column:slashed_epoch;type:bigint" json:"slashed_epoch"`
	StoragePricePerEpoch string  `gorm:"column:storage_price_per_epoch;type:text;not null" json:"storage_price_per_epoch"` // The amount of FIL (in attoFIL) that will be transferred from the client to the provider every epoch this deal is active for.
	ProviderCollateral   string  `gorm:"column:provider_collateral;type:text;not null" json:"provider_collateral"`         // The amount of FIL (in attoFIL) the provider has pledged as collateral. The Provider deal collateral is only slashed when a sector is terminated before the deal expires.
	ClientCollateral     string  `gorm:"column:client_collateral;type:text;not null" json:"client_collateral"`             // The amount of FIL (in attoFIL) the client has pledged as collateral.
	Label                *string `gorm:"column:label;type:text" json:"label"`                                              // An arbitrary client chosen label to apply to the deal.
	Height               int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"`                               // Epoch at which this deal proposal was added or changed.
	IsString             *bool   `gorm:"column:is_string;type:boolean" json:"is_string"`                                   // When true Label contains a valid UTF-8 string encoded in base64. When false Label contains raw bytes encoded in base64. Required by FIP: 27
}

// TableName MarketDealProposal's table name
func (*MarketDealProposal) TableName() string {
	return TableNameMarketDealProposal
}