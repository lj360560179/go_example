package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"
	"github.com/garyburd/redigo/redis"
)

func GetRedis(c *gin.Context){
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()
	RedisConn.Do("SET", "foo", "bar")
	foo, err := redis.String(RedisConn.Do("GET", "foo"))
	if err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(foo,c)
	return
}