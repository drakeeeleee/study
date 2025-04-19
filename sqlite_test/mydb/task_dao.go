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

func (t *TaskDao) Create(tasks []*model.Task) error {
	if err := t.db.Create(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (t *TaskDao) SelectFirst(taskId string) (*model.Task, error) {
	var task model.Task
	if err := t.db.First(&task, "task_id = ?", taskId).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskDao) CheckAllExistence(tasks []*model.Task) (bool, error) {
	var taskIds []string
	for _, task := range tasks {
		taskIds = append(taskIds, task.TaskId)
	}

	var foundTasks []*model.Task
	if err := t.db.Where("task_id IN ?", taskIds).Find(&foundTasks).Error; err != nil {
		return false, err
	}

	return len(foundTasks) == len(tasks), nil

}
