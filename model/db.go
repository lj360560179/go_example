package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
	"time"
	"go_server/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接
var DB *gorm.DB

var MongoDB *mgo.Session

// RedisPool Redis连接池
var RedisPool *redis.Pool


//func initDB() {
//	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(-1)
//	}
//	db.LogMode(true)
//	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
//	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
//	DB = db
//}

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func initMongo() {
	session, err := mgo.Dial(config.MongoConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	MongoDB = session
	// Optional. Switch the session to a monotonic behavior.
	MongoDB.SetMode(mgo.Monotonic, true)
}



func init() {
	initMongo()
	//initDB()
	initRedis()

}