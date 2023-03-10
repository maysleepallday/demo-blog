package router

import (
	"Blog/handler"
	"Blog/middleware"
	"github.com/gin-gonic/gin"

	_ "Blog/docs" // 导入swagger文档用的
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.RateLimitMiddleware())

	user := r.Group("/user")
	user.POST("/register", handler.RegisterHandler)
	user.POST("/login", handler.LoginHandler)

	post := r.Group("/post")
	//post.Use(middleware.AuthMiddleware())
	post.POST("/create", handler.PostCreateHandler)
	post.GET("/:id", handler.PostDetailHandler)
	post.GET("/vote", handler.PostVoteHandler)

	return r
}
