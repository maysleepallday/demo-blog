package dao

import (
	"Blog/models"
)

func PostDetail(identity int64) (*models.Post, error) {
	post := new(models.Post)

	err := db.Where("post_id", identity).First(post).Error
	if err != nil {
		return nil, err
	}

	return post, nil

}

func PostCreate(p *models.Post) error {
	err := db.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}
