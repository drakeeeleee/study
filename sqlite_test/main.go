package main

import (
	"fmt"
	"log"
	"os"
	"sqlitetest/mydb"
	"sqlitetest/mydb/dao"
	"sqlitetest/mydb/model"
	"sync"
)

func main() {
	dbName := "main.db"
	db, err := mydb.NewLocalDb(dbName)
	if err != nil {
		log.Fatalf("failed to make local db: %v", err)
	}
	defer func() { _ = os.Remove(dbName) }()
	taskDao := dao.NewTaskDao()

	tasks := []*model.Task{
		{TaskId: "task_id_1"},
		{TaskId: "task_id_2"},
		{TaskId: "task_id_3"},
	}

	numGoroutine := 10
	var wait sync.WaitGroup
	wait.Add(numGoroutine)
	for i := range numGoroutine {
		go func() {
			tx := db.Begin()

			tasks := tasks
			if existence, err := taskDao.CheckAllExistence(tx, tasks); err != nil {
				tx.Rollback()
				log.Fatalf("failed to check existence at %dth trial: %v", i, err)
			} else if existence {
				fmt.Printf("tasks already exist at %dth trial\n", i)
			} else if err = taskDao.Create(tx, tasks); err != nil {
				tx.Rollback()
				log.Fatalf("failed to create tasks at %dth trial: %v", i, err)
			}

			tx.Commit()
			wait.Done()
		}()
	}

	wait.Wait()
}
