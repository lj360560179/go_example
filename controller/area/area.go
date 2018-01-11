package area

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_server/model"
	"go_server/controller/common"
)

func FindAllArea(c *gin.Context){
	var animal model.Animal
	if err := model.DB.First(&animal, "4").Error; err != nil {
		common.SendErrJSON("错误的分类id", c)
		return
	}
	c.JSON(http.StatusOK, animal)
	return
}