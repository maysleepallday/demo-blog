package models

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	Name string `gorm:column:"community_name"`
}

func (Community) TableName() string {
	return "community"
}
