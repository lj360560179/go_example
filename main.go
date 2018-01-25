package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"go_server/config"
	"go_server/model"
	"github.com/gin-gonic/gin"
	"go_server/controller/area"
	"go_server/controller/tproduct"
	"go_server/controller/es"
	"go_server/controller/mongo"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/mgo.v2"
)

func init() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	//打印SQL语句
	db.LogMode(true)
	model.DB = db

	session, err := mgo.Dial("192.168.99.100:27017")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	model.MongoDB = session
	// Optional. Switch the session to a monotonic behavior.
	model.MongoDB.SetMode(mgo.Monotonic, true)

}


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
	router.Run(":8000")
}