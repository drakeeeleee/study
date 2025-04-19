package main

import (
	"log"
	"os"
	"sqlitetest/mydb"
	"sqlitetest/mydb/model"
)

func main() {
	dbName := "test.db"
	db, err := mydb.NewLocalDb(dbName)
	defer func() { _ = os.Remove(dbName) }()
	if err != nil {
		log.Fatalf("failed to make local db: %v", err)
	}

	db.Create(&model.Task{TaskId: "task_id_1"})
}
