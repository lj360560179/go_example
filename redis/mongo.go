package redis

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"go_server/model"
	"log"
	"fmt"
)

type Msg struct {
	id string
	msg string
}

func GetByMo(c *gin.Context){
	a := model.MongoDB.DB("delay").C("msg")
	result := Msg{}
	err := a.Find(bson.M{"id": "msg"}).One(&result)
	if err != nil {
		log.Fatal(err)
		sendErrorMsg("mmpp",c)
	}
	sendResponse(result,c)
	return
}

func saveMongo(v string)  {
	a := model.MongoDB.DB("test").C("people")
	err:=a.Insert(Msg{
		id:"msg",
		msg:v,
	})
	if err!=nil{
		fmt.Println(err)
	}
}