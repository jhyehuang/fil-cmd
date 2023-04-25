// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameInternalMessage = "internal_messages"

// InternalMessage mapped from table <internal_messages>
type InternalMessage struct {
	Height        int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"`         // Epoch this message was executed at.
	Cid           string  `gorm:"column:cid;type:text;primaryKey" json:"cid"`                 // CID of the message.
	StateRoot     string  `gorm:"column:state_root;type:text;not null" json:"state_root"`     // CID of the parent state root at which this message was executed.
	SourceMessage *string `gorm:"column:source_message;type:text" json:"source_message"`      // CID of the message that caused this message to be sent.
	From          string  `gorm:"column:from;type:text;not null" json:"from"`                 // Address of the actor that sent the message.
	To            string  `gorm:"column:to;type:text;not null" json:"to"`                     // Address of the actor that received the message.
	Value         float64 `gorm:"column:value;type:numeric;not null" json:"value"`            // Amount of FIL (in attoFIL) transferred by this message.
	Method        int64   `gorm:"column:method;type:bigint;not null" json:"method"`           // The method number invoked on the recipient actor. Only unique to the actor the method is being invoked on. A method number of 0 is a plain token transfer - no method exectution.
	ActorName     string  `gorm:"column:actor_name;type:text;not null" json:"actor_name"`     // The full versioned name of the actor that received the message (for example fil/3/storagepower).
	ActorFamily   string  `gorm:"column:actor_family;type:text;not null" json:"actor_family"` // The short unversioned name of the actor that received the message (for example storagepower).
	ExitCode      int64   `gorm:"column:exit_code;type:bigint;not null" json:"exit_code"`     // The exit code that was returned as a result of executing the message. Exit code 0 indicates success. Codes 0-15 are reserved for use by the runtime. Codes 16-31 are common codes shared by different actors. Codes 32+ are actor specific.
	GasUsed       int64   `gorm:"column:gas_used;type:bigint;not null" json:"gas_used"`       // A measure of the amount of resources (or units of gas) consumed, in order to execute a message.
}

// TableName InternalMessage's table name
func (*InternalMessage) TableName() string {
	return TableNameInternalMessage
}