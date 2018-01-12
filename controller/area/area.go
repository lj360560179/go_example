package area

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"
)

func FindAllArea(c *gin.Context){
	var area []model.Area
	if err := model.DB.Where("areaId = ?", c.Param("id")).Find(&area).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(area,c)
	return
}