package middleware

import (
	"Blog/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "need login",
			})
			c.Abort()
			return
		}

		uc, err := jwt.AnalyzeToken(auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", uc.Identity)
		c.Next()
	}
}
