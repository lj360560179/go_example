package es

import (
	"github.com/olivere/elastic"
	"github.com/gin-gonic/gin"
	"go_server/controller/common"

	"context"
)



func SerchEs(c *gin.Context){
	client, err := elastic.NewClient()
	if err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	termQuery := elastic.NewTermQuery("title", c.Query("name"))
	searchResult, err := client.Search().Index("pdt").Query(termQuery).From(0).Size(10).Pretty(true).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	common.SendResponse(searchResult,c)
}
