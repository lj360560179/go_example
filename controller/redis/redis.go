package redis

import (
	"github.com/gin-gonic/gin"
	"go_server/utils"
	"go_server/controller/common"

	"strconv"
)



func GetRedis(c *gin.Context){
    result := utils.GetString(c.Query("key"))
	common.SendResponse(result,c)
	return
}

func SetRedis(c *gin.Context){
	k := c.Query("k")
	v := c.Query("v")
	common.SendResponse(utils.SetStringTime(k,v,10),c)
	return
}

func ZetTest(c *gin.Context)  {
	k := c.Query("k")
	v := c.Query("v")
	s,_ := strconv.ParseInt(c.Query("s"),10,32)
	utils.AddZet(k,v,int32(s))
	zset := utils.GetSoredSetByRange(k,0,10,true)

	result := make(map[string]int64)
	for _, set := range zset {
		result[set] = utils.ZetScore(k,set)
	}
	common.SendResponse(result,c)
}

func AddToList()  {
	
}

//todo
func job()  {
	
}
