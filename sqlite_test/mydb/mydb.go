package mydb

import (
	"log"
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
		log.Fatalf("failed to auto migrate: %v", err)
	}
	return db, nil
}
