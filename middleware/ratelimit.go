package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware() func(c *gin.Context) {
	//填充速率，令牌桶容量
	bucket := ratelimit.NewBucket(time.Second*2, 1)

	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) != 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
