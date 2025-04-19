package mydb

import (
	"sqlitetest/mydb/model"

	"gorm.io/gorm"
)

type TaskDao struct {
	db *gorm.DB
}

func (t *TaskDao) CreateTask(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}
