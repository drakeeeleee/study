package mydb

import (
	"sqlitetest/mydb/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewLocalDb(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Task{}, &model.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
