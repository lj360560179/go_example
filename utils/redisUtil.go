package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"go_server/model"
)
/**
	缓存String
 */
func SetString(key,value string) bool {
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
func SetStringTime(key,value string, timeOutSeconds int ) bool{
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
func GetString(key string) string {
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
func RemoceString(key string) bool{
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
func SetObj(key string,obj interface{}) bool {
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
func AddZet(key ,value string ,score int32) bool {
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
func RemZet(key string) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := RedisConn.Do("ZREM", key)
	if err != nil {
		fmt.Println(err)
	}
	return result == 1
}

func GetSoredSetByRange(key string,startRange,endRange int ,orderByDesc bool) [] string {
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
func ZetScore(key,value string) int64  {
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
func CountList(key string) int {
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
func InsertList(key string, value... string) bool {
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
func RangeList(key string, start,end int) []string {
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
func RemListValue(key,value string ,count int) bool {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	result ,err := redis.Int(RedisConn.Do("LREM", key,count,value))
	if err != nil {
		fmt.Println(err)
	}
	return result > 0
}
