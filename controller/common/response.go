package common

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"net/http"
)

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("缺少 gin.Context")
	}
	var c *gin.Context
	var errNo = model.ErrorCode.ERROR
	c.JSON(http.StatusBadRequest, gin.H{"error": errNo})
}