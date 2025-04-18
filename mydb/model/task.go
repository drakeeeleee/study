package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskId      string `gorm:"unique"`
	EndPoint    string
	Description string
}
