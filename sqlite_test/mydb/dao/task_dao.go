package dao

import (
	"sqlitetest/mydb/model"

	"gorm.io/gorm"
)

type TaskDao struct {
}

func NewTaskDao() *TaskDao {
	return &TaskDao{}
}

func (t *TaskDao) Create(db *gorm.DB, tasks []*model.Task) error {
	if err := db.Create(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (t *TaskDao) SelectFirst(db *gorm.DB, taskId string) (*model.Task, error) {
	var task model.Task
	if err := db.First(&task, "task_id = ?", taskId).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskDao) CheckAllExistence(db *gorm.DB, tasks []*model.Task) (bool, error) {
	var taskIds []string
	for _, task := range tasks {
		taskIds = append(taskIds, task.TaskId)
	}

	var foundTasks []*model.Task
	if err := db.Where("task_id IN ?", taskIds).Find(&foundTasks).Error; err != nil {
		return false, err
	}

	return len(foundTasks) == len(tasks), nil

}
