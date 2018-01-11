package common

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"net/http"
)

func SendErrorMsg(msg string,c *gin.Context){
	var errNo = model.ErrorCode.ERROR
	c.JSON(http.StatusOK, gin.H{
		"error": errNo,
		"msg":msg,
	})
}