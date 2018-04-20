package main

import (
	"github.com/gin-gonic/gin"
	"go_server/controller/es"
	"go_server/controller/mongo"
	"go_server/controller/redis"

)

func main(){
	router := gin.Default()
	router.GET("/mongo",mongo.GetByMo)
	router.GET("/redis",redis.GetRedis)
	router.Run(":8000")
}