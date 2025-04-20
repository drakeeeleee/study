package dao

import (
	"os"
	"sqlitetest/mydb"
	"sqlitetest/mydb/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserDao(t *testing.T) {
	as := assert.New(t)

	dbName := "user_dao.db"
	db, err := mydb.NewLocalDb(dbName)
	as.NoError(err)
	defer func() { _ = os.Remove(dbName) }()

	userDao := NewUserDao()

	err = userDao.Create(db, []*model.User{{UserId: "test_user_id_1"}, {UserId: "test_user_id_2"}})
	as.NoError(err)
}
