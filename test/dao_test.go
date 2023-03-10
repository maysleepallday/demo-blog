package test

import (
	"Blog/dao"
	"Blog/models"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestDao(t *testing.T) {
	err := dao.Init()
	if err != nil {
		return
	}
	db := dao.GetDB()

	err = db.Model(&models.User{}).Where("name", "zhangsan").First(new(string)).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("用户不存在")
		return
	}
	fmt.Println("用户已存在")
}
