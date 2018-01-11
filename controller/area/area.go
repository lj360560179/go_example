package area

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_server/model"

	"go_server/controller/common"
)

func FindAllArea(c *gin.Context){
	var animal model.Animal
	if err := model.DB.First(&animal, "1").Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	c.JSON(http.StatusOK, animal)
	return
}