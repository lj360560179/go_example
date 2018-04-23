package main

import (
	"github.com/gin-gonic/gin"
	"go_server/controller/mongo"
	"go_server/controller/redis"
)

func main(){
	router := gin.Default()
	router.GET("/getString",redis.GetRedis)
	router.GET("/mongo",mongo.GetByMo)
	router.GET("/setString",redis.SetRedis)
	router.GET("/zetTest",redis.ZetTest)
	router.Run(":8000")


}

