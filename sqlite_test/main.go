package main

import (
	"fmt"
	"log"
	"os"
	"sqlitetest/mydb"
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
	taskDao := mydb.NewTaskDao(db)

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
			tasks := tasks
			if existence, err := taskDao.CheckAllExistence(tasks); err != nil {
				log.Fatalf("failed to check existence at %dth tria: %v", i, err)
			} else if existence {
				fmt.Printf("tasks already exist at %dth trial\n", i)
			} else if err = taskDao.Create(tasks); err != nil {
				log.Fatalf("failed to create tasks at %dth tria: %v", i, err)
			}
			wait.Done()
		}()
	}

	wait.Wait()
}
