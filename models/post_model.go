package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Identity    int64  `gorm:"column:post_id"`
	Title       string `gorm:"column:title"`
	Vote        int64  `gorm:"column:vote"`
	UserID      int64
	CommunityID int64
}

func (Post) TableName() string {
	return "post"
}
