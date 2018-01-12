package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"go_server/config"
	"go_server/model"
	"github.com/gin-gonic/gin"
	"go_server/controller/area"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
}


func main(){
	router := gin.Default()
	router.GET("/area/:id",area.FindAllArea)
	router.Run(":8000")
}