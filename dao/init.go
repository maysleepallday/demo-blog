package dao

import (
	"Blog/config"
	"Blog/models"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var rdb *redis.Client

func InitMySQL() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second)

	db.AutoMigrate(&models.User{}, &models.Community{}, &models.Post{})

	return nil
}

func InitRedis(cfg *config.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
