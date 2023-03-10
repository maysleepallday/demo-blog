package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get("user_id")

	if !ok {
		return 0, errors.New("未登录")
	}

	return uid.(int64), nil
}
