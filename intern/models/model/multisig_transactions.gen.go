// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMultisigTransaction = "multisig_transactions"

// MultisigTransaction mapped from table <multisig_transactions>
type MultisigTransaction struct {
	Height        int64    `gorm:"column:height;type:bigint;primaryKey" json:"height"`                 // Epoch at which this transaction was executed.
	MultisigID    string   `gorm:"column:multisig_id;type:text;primaryKey" json:"multisig_id"`         // Address of the multisig actor involved in the transaction.
	StateRoot     string   `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`           // CID of the parent state root at this epoch.
	TransactionID int64    `gorm:"column:transaction_id;type:bigint;primaryKey" json:"transaction_id"` // Number identifier for the transaction - unique per multisig.
	To            string   `gorm:"column:to;type:text;not null" json:"to"`                             // Address of the recipient who will be sent a message if the proposal is approved.
	Value         string   `gorm:"column:value;type:text;not null" json:"value"`                       // Amount of FIL (in attoFIL) that will be transferred if the proposal is approved.
	Method        int64    `gorm:"column:method;type:bigint;not null" json:"method"`                   // The method number to invoke on the recipient if the proposal is approved. Only unique to the actor the method is being invoked on. A method number of 0 is a plain token transfer - no method exectution.
	Params        *[]uint8 `gorm:"column:params;type:bytea" json:"params"`                             // CBOR encoded bytes of parameters to send to the method that will be invoked if the proposal is approved.
	Approved      string   `gorm:"column:approved;type:jsonb;not null" json:"approved"`                // Addresses of signers who have approved the transaction. 0th entry is the proposer.
}

// TableName MultisigTransaction's table name
func (*MultisigTransaction) TableName() string {
	return TableNameMultisigTransaction
}
