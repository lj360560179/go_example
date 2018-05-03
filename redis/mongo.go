package redis

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"go_server/model"
	"log"
	"fmt"
)

type Msg struct {
	Id string
	Msg string
}

func GetByMo(c *gin.Context){
	a := model.MongoDB.DB("delay").C("msg")
	result := [] Msg{}
	err := a.Find(bson.M{"id": "msg"}).All(&result)
	if err != nil {
		log.Fatal(err)
		sendErrorMsg("mmpp",c)
	}
	sendResponse(result,c)
	return
}

func saveMongo(v string)  {
	a := model.MongoDB.DB("delay").C("msg")
	err:=a.Insert(&Msg{
		"msg",
		v,
	})
	if err!=nil{
		fmt.Println(err)
	}
}