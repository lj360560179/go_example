package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
	"strconv"
)

func GetRedis(c *gin.Context){
    result := getString(c.Query("key"))
	common.SendResponse(result,c)
	return
}

func SetRedis(c *gin.Context){
	k := c.Query("k")
	v := c.Query("v")
	common.SendResponse(setStringTime(k,v,10),c)
	return
}

func ZetTest(c *gin.Context)  {
	k := c.Query("k")
	v := c.Query("v")
	s,_ := strconv.ParseInt(c.Query("s"),10,32)
	addZet(k,v,int32(s))
	common.SendResponse(getSoredSetByRange(k,0,10,true),c)
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

/**
	删除
 */
func remoceString(key string) bool{
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result, err := RedisConn.Do("DEL", key)
	if err != nil {
		fmt.Println(err)
	}
	return result =="1"
}

/**
 	存对象
 */
func setObj(key string,obj interface{}) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	value, err := json.Marshal(obj)
	if  err != nil {
		fmt.Println(err)
	}
	result ,err := RedisConn.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
	}

	return result == "OK"
}
/**
	向有序列表存入
 */
func addZet(key ,value string ,score int32) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := RedisConn.Do("ZADD", key, score,value)
	if err != nil {
		fmt.Println(err)
	}
	return result == 1
}
/**
	删除有序列表元素
 */
func remZet(key string) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := RedisConn.Do("ZREM", key)
	if err != nil {
		fmt.Println(err)
	}
	return result == 1
}

func getSoredSetByRange(key string,startRange,endRange int ,orderByDesc bool) []string {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	fing := "ZRANGE"
	if orderByDesc {
		fing = "ZREVRANGE"
	}
	result ,err :=  redis.Strings(RedisConn.Do(fing, key,startRange,endRange))
	if err != nil {
		fmt.Println(err)
	}
	return result
}