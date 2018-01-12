package tproduct

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"
)

func FindPdtById(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Where("guid = ?", c.Param("id")).Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(pdt,c)
	return
}

func FindHotPdtList(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Limit(5).Where(map[string]interface{}{"isDelete": "FALSE", "sellerId": c.Query("id"),"status":"ON_SALE","checkStatus":"CHECKSUCCESS"}).Order("updateTime desc").Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(pdt,c)
	return
}