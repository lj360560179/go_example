package mongo

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"go_server/model"
	"go_server/controller/common"
)

type Person struct{
	Name string
	Phone string
}

func GetByMo(c *gin.Context){
	a := model.MongoDB.DB("test").C("people")
	result := Person{}
	a.Find(bson.M{"name": "Cla"}).One(&result)
	common.SendResponse(result,c)
	return
}