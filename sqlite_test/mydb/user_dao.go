package mydb

import (
	"sqlitetest/mydb/model"

	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) Create(db *gorm.DB, users []*model.User) error {
	if err := db.Create(users).Error; err != nil {
		return err
	}
	return nil
}
