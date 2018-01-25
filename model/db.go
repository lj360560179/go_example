package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

// DB 数据库连接
var DB *gorm.DB

var MongoDB *mgo.Session