package dao

import (
	"os"
	"sqlitetest/mydb"
	"sqlitetest/mydb/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TaskDao(t *testing.T) {
	as := assert.New(t)

	dbName := "task_dao.db"
	db, err := mydb.NewLocalDb(dbName)
	as.NoError(err)
	defer func() { _ = os.Remove(dbName) }()

	taskDao := NewTaskDao()

	err = taskDao.Create(db, []*model.Task{{TaskId: "test_eval_id_1"}, {TaskId: "test_eval_id_2"}})
	as.NoError(err)

	task, err := taskDao.SelectFirst(db, "test_eval_id_1")
	as.NoError(err)
	as.NotNil(task)

	existence, err := taskDao.CheckAllExistence(db, []*model.Task{{TaskId: "test_eval_id_1"}, {TaskId: "test_eval_id_2"}})
	as.NoError(err)
	as.True(existence)

	existence, err = taskDao.CheckAllExistence(db, []*model.Task{{TaskId: "test_eval_id_3"}, {TaskId: "test_eval_id_2"}})
	as.NoError(err)
	as.False(existence)

	existence, err = taskDao.CheckAllExistence(db, []*model.Task{{TaskId: "test_eval_id_3"}, {TaskId: "test_eval_id_4"}})
	as.NoError(err)
	as.False(existence)
}
