package handler

import (
	"Blog/logic"
	"Blog/param"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHandler 用户登录接口
// @Summary      用户登录接口
// @Description  实现用户登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param object query param.UserRegister false "登陆参数"
// @Success      200  {object}  ResponseData
// @Router       /user/register [post]
func RegisterHandler(c *gin.Context) {
	//参数校验
	p := new(param.UserRegister)

	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		c.JSON(http.StatusOK, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	// 业务处理 logic.RegisterLogic
	err := logic.RegisterLogic(p)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "注册成功",
	})
}

func LoginHandler(c *gin.Context) {
	//参数校验
	p := new(param.UserLogin)

	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		c.JSON(http.StatusOK, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	// 业务处理 logic.LoginLogic
	user, err := logic.LoginLogic(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("user_id", user.Identity)

	c.JSON(http.StatusOK, gin.H{
		"token": user.Token,
	})

}
