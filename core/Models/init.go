package models

import (
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db = initGorm()

func GetDB() *gorm.DB {
	return _db
}
func initGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open("whk:123456@tcp(192.168.118.132:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Printf("gorm new db Error:%v", err)
		return nil
	}
	return db
}
func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
