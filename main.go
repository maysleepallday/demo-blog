package main

import (
	"Blog/config"
	"Blog/dao"
	"Blog/router"
	"Blog/utils/snowflake"
	"fmt"
)

// @title blog项目接口文档
// @version 0.1
// @description 基于gin的个人博客项目

// @host localhost:8080
// @BasePath /
func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := dao.InitMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	if err := dao.InitRedis(config.Cfg.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	if err := snowflake.Init("2020-02-03", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetupRouter()
	err := r.Run()
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
