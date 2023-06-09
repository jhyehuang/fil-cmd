// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameInternalParsedMessage = "internal_parsed_messages"

// InternalParsedMessage mapped from table <internal_parsed_messages>
type InternalParsedMessage struct {
	Height int64   `gorm:"column:height;type:bigint;primaryKey" json:"height"` // Epoch this message was executed at.
	Cid    string  `gorm:"column:cid;type:text;primaryKey" json:"cid"`         // CID of the message.
	From   string  `gorm:"column:from;type:text;not null" json:"from"`         // Address of the actor that sent the message.
	To     string  `gorm:"column:to;type:text;not null" json:"to"`             // Address of the actor that received the message.
	Value  float64 `gorm:"column:value;type:numeric;not null" json:"value"`    // Amount of FIL (in attoFIL) transferred by this message.
	Method string  `gorm:"column:method;type:text;not null" json:"method"`     // The method number invoked on the recipient actor. Only unique to the actor the method is being invoked on. A method number of 0 is a plain token transfer - no method exectution.
	Params *string `gorm:"column:params;type:jsonb" json:"params"`             // Method parameters parsed and serialized as a JSON object.
}

// TableName InternalParsedMessage's table name
func (*InternalParsedMessage) TableName() string {
	return TableNameInternalParsedMessage
}
