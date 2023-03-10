package logic

import (
	"Blog/dao"
	"Blog/models"
	"Blog/param"
	"Blog/utils/snowflake"
	"errors"
	"strconv"
)

func PostCreate(p *param.PostCreate) error {
	identity := snowflake.GenerateID()

	post := &models.Post{
		Identity:    identity,
		Title:       p.Title,
		Vote:        0,
		UserID:      p.UserID,
		CommunityID: p.CommunityID,
	}

	err := dao.PostCreate(post)
	if err != nil {
		return errors.New("帖子创建失败")
	}

	err = dao.CreatePost(p.UserID, p.CommunityID)
	if err != nil {
		return errors.New("帖子创建失败")
	}

	return nil
}

func PostDetailLogic(identity int64) (*models.Post, error) {
	return dao.PostDetail(identity)
	// todo
}

func PostVote(userID int64, p *param.PostVote) error {
	// 判断投票类型
	// 只提供单个点赞，之前已经赞过的，取消点赞voteType=-1
	dao.PostVote(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Value))
	return nil
}
