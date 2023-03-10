package dao

import (
	"Blog/models"
	"gorm.io/gorm"
)

func CheckUserExist(name string) bool {

	err := db.Model(&models.User{}).Where("user_name", name).First(new(string)).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

func UserCreate(user *models.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(name, password string) (*models.User, error) {
	user := new(models.User)

	err := db.Where("user_name = ? AND password = ?", name, password).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
