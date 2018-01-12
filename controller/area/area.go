package area

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_server/model"

	"go_server/controller/common"
)

func FindAllArea(c *gin.Context){
	var area []model.Area
	if err := model.DB.Where("levels = ?", "1").Find(&area).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	c.JSON(http.StatusOK, area)
	return
}