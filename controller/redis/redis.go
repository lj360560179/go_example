package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/controller/common"
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
	fmt.Println(value)

	if  err != nil {
		common.SendErrorMsg("json!!!",c)
	}
	uid,_:= uuid.NewV1()
	result ,err := RedisConn.Do("SETEX",uid, int(60*60*24*time.Second),string(value))
	if err != nil {
		fmt.Println(err)
		common.SendErrorMsg("没存进去obj!!!",c)
	}
	s :=time.Now().Add(time.Duration(10000)).Unix()
	fmt.Println(s)
	if addZet("ZSET",uid.String(),s) {
		common.SendErrorMsg("出错了呢~",c)
	}
	common.SendResponse(result,c)
}

func GetZset(c *gin.Context)  {
	zset :=getSoredSetByRange("ZSET",0,10,true)
	result := make(map[string]int64)
	for _, set := range zset {
		result[set] = zetScore("ZSET",set)
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
	result ,err := execRedisCommand("SET", key, value)
	if err != nil {
		fmt.Println(err)
	}
	return result == "OK"
}
/**
	缓存String 设置过期时间
 */
func setStringTime(key,value string, timeOutSeconds int ) bool{
	result ,err := execRedisCommand("SETEX",key, timeOutSeconds, value)
	if err != nil {
		fmt.Println(err)
	}
	return result == "OK"
}

/**
     从缓存中获取String
 */
func getString(key string) string {
	result, err := redis.String(execRedisCommand("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return result
}

/**
	删除
 */
func remoceString(key string) bool{
	result, err := execRedisCommand("DEL", key)
	if err != nil {
		fmt.Println(err)
	}
	return result =="1"
}

/**
 	存对象
 */
func setObj(key string,obj interface{}) bool {
	value, err := json.Marshal(obj)
	if  err != nil {
		fmt.Println(err)
	}
	result ,err := execRedisCommand("SET", key, value)
	if err != nil {
		fmt.Println(err)
	}

	return result == "OK"
}
/**
	向有序列表存入
 */
func addZet(key ,value string ,score int64) bool {
	result ,err := execRedisCommand("ZADD", key, score,value)
	if err != nil {
		fmt.Println(err)
	}
	return result == 1
}
/**
	删除有序列表元素
 */
func remZet(key string) bool {
	result ,err := execRedisCommand("ZREM", key)
	if err != nil {
		fmt.Println(err)
	}
	return result == 1
}

func getSoredSetByRange(key string,startRange,endRange int ,orderByDesc bool) [] string {
	fing := "ZRANGE"
	if orderByDesc {
		fing = "ZREVRANGE"
	}
	result ,err :=  redis.Strings(execRedisCommand(fing, key,startRange,endRange))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
/**
	获取分数
 */
func zetScore(key,value string) int64  {
	result ,err :=  redis.Int64(execRedisCommand("ZSCORE", key,value))
	if err != nil {
		fmt.Println(err)
	}
	return result
}

/**
	获取列表长度
 */
func countList(key string) int {
	result ,err :=  redis.Int(execRedisCommand("LLEN", key))
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
	result ,err := redis.Strings(execRedisCommand("LRANGE", start,end))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
/**
 	删除LIST数据
 */
func remListValue(key,value string ,count int) bool {
	result ,err := redis.Int(execRedisCommand("LREM", key,count,value))
	if err != nil {
		fmt.Println(err)
	}
	return result > 0
}
// 执行redis命令, 执行完成后连接自动放回连接池
func execRedisCommand(command string, args ...interface{}) (interface{}, error) {
	redis := model.RedisPool.Get()
	defer redis.Close()
	return redis.Do(command, args...)
}