package main

import (
	"github.com/gin-gonic/gin"

	"go_server/redis"
)

func main(){
	router := gin.Default()
	router.GET("/getString",redis.GetRedis)
	router.GET("/mongo",redis.GetByMo)
	router.GET("/setString",redis.SetRedis)
	router.GET("/addZset",redis.AddRedisMq)
	router.GET("/zetTest",redis.GetZset)
	router.GET("/list",redis.GetList)
	router.Run(":8000")


}

