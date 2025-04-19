package mydb

import (
	"sqlitetest/mydb/model"

	"gorm.io/gorm"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao(db *gorm.DB) *TaskDao {
	return &TaskDao{db: db}
}

func (t *TaskDao) CreateTask(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *TaskDao) CheckExistence(tasks []*model.Task) ([]*model.Task, error) {
	var taskIds []string
	for _, task := range tasks {
		taskIds = append(taskIds, task.TaskId)
	}

	var foundTasks []*model.Task
	if err := t.db.Where("task_id IN ?", taskIds).Find(&foundTasks).Error; err != nil {
		return nil, err
	}

	return foundTasks, nil

}
