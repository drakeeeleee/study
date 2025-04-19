package mydb

import (
	"os"
	"sqlitetest/mydb/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TaskDao(t *testing.T) {
	as := assert.New(t)

	dbName := "task_dao.db"
	db, err := NewLocalDb(dbName)
	as.NoError(err)
	defer func() { _ = os.Remove(dbName) }()

	taskDao := NewTaskDao(db)

	err = taskDao.Create([]*model.Task{{TaskId: "test_eval_id_1"}, {TaskId: "test_eval_id_2"}})
	as.NoError(err)

	task, err := taskDao.SelectFirst("test_eval_id_1")
	as.NoError(err)
	as.NotNil(task)
}
