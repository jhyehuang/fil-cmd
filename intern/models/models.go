package models

import (
	"github.com/jhyehuang/fil-cmd/intern/models/model"
	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) (*gorm.DB, error) {

	if err := migrateSchemas(db); err != nil {
		return nil, err
	}
	return db, nil
}

func migrateSchemas(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.ChainReward{},
		&model.BlockHeader{},
		&model.DerivedGasOutput{},
		&model.MinerLockedFund{},
		&model.MinerInfo{},
		&model.Message{},
		&model.MessageParam{},
	); err != nil {
		return err
	}
	return nil
}
