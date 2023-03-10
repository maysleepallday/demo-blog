package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identity int64  `gorm:"column:user_id"`
	Name     string `gorm:"column:user_name"`
	Password string
	Token    string
}

func (User) TableName() string {
	return "user"
}
