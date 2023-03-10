package handler

import (
	"Blog/logic"
	"Blog/param"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostCreateHandler(c *gin.Context) {
	p := new(param.PostCreate)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "请求参数错误",
		})
		return
	}
	id, _ := getCurrentUserID(c)
	p.UserID = id

	println("-----------------------")
	println(p.Title)
	println(p.CommunityID)
	err := logic.PostCreate(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "帖子创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "帖子创建成功",
	})
}

func PostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	//
	data, err := logic.PostDetailLogic(pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "查询帖子出错",
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	return
}

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(param.PostVote)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "请求参数错误",
		})
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户未登录",
		})
		return
	}
	// 具体投票的业务逻辑
	if err := logic.PostVote(userID, p); err != nil {

		return
	}

}
