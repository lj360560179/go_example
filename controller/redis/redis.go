package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/controller/common"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"go_server/model"
	"time"
	"github.com/satori/go.uuid"
)

type Msg struct {
	id string
	msg string
}

func AddRedisMq(c *gin.Context)  {
	v := c.Query("v")
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	value, err := json.Marshal( Msg{"uuid", v})
	if  err != nil {
		fmt.Println(err)
	}
	result ,err := RedisConn.Do("SETEX",uuid.NamespaceDNS, 60*60*24*time.Second, value)
	if err != nil {
		fmt.Println(err)
	}
	common.SendResponse(result,c)
}

func ZetTest(c *gin.Context)  {
	k := c.Query("k")
	v := c.Query("v")
	s,_ := strconv.ParseInt(c.Query("s"),10,32)
	addZet(k,v,int32(s))
	zset :=getSoredSetByRange(k,0,10,true)

	result := make(map[string]int64)
	for _, set := range zset {
		result[set] = zetScore(k,set)
	}
	common.SendResponse(result,c)
}

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

func getSoredSetByRange(key string,startRange,endRange int ,orderByDesc bool) [] string {
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



/**
	获取分数
 */
func zetScore(key,value string) int64  {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err :=  redis.Int64(RedisConn.Do("ZSCORE", key,value))
	if err != nil {
		fmt.Println(err)
	}
	return result
}

/**
	获取列表长度
 */
func countList(key string) int {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err :=  redis.Int(RedisConn.Do("LLEN", key))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
/**
	添加元素到list（使用右进）
 */
func insertList(key string, value... string) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	err := RedisConn.Send("RPUSH", key,value)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
/**
	获取list
 */
func rangeList(key string, start,end int) []string {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := redis.Strings(RedisConn.Do("LRANGE", start,end))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
/**
 	删除LIST数据
 */
func remListValue(key,value string ,count int) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := redis.Int(RedisConn.Do("LREM", key,count,value))
	if err != nil {
		fmt.Println(err)
	}
	return result > 0
}



