package main

import (
	"github.com/gin-gonic/gin"
	"go_server/controller/area"
	"go_server/controller/tproduct"
	"go_server/controller/es"
	"go_server/controller/mongo"
	"go_server/controller/redis"

)




func main(){
	router := gin.Default()
	router.GET("/area/:id",area.FindAllArea)
	router.GET("/prodcut/:id",tproduct.FindPdtById)
	router.GET("/hotprodcuts",tproduct.FindHotPdtList)
	router.GET("/prodcuts",tproduct.FindBySellerId)
	router.GET("/noprodcuts",tproduct.FindNpassBySellerId)
	router.GET("/esindex",es.EsIndex)
	router.GET("/esserch",es.SerchEs)
	router.GET("/mongo",mongo.GetByMo)
	router.GET("/redis",redis.GetRedis)
	router.Run(":8000")
}