// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGopgMigration = "gopg_migrations"

// GopgMigration mapped from table <gopg_migrations>
type GopgMigration struct {
	ID        int32      `gorm:"column:id;type:integer;not null" json:"id"`
	Version   *int64     `gorm:"column:version;type:bigint" json:"version"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp with time zone" json:"created_at"`
}

// TableName GopgMigration's table name
func (*GopgMigration) TableName() string {
	return TableNameGopgMigration
}