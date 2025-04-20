package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId string `gorm:"unique"`
	Gender string
}
