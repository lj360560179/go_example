package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func GetRedis(c *gin.Context){
    result := getString(c.Query("key"))
	common.SendResponse(result,c)
	return
}

func SetRedis(c *gin.Context){
	key := c.Query("key")
	value := c.Query("value")
	common.SendResponse(setStringTime(key,value,10),c)
	return
}

/**
	缓存String
 */
func setString(key,value string) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := RedisConn.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
	}
	return result == "OK"
}
/**
	缓存String 设置过期时间
 */
func setStringTime(key,value string, timeOutSeconds int ) bool{
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := RedisConn.Do("SETEX",key, timeOutSeconds, value)
	if err != nil {
		fmt.Println(err)
	}
	return result == "OK"
}

/**
     从缓存中获取String
 */
func getString(key string) string {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result, err := redis.String(RedisConn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
