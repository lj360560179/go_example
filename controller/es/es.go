package es

import (
	"github.com/olivere/elastic"
	"github.com/gin-gonic/gin"
	"go_server/controller/common"
	"go_server/model"

	"context"
	"net/http"
	"strconv"
)

func EsIndex(c *gin.Context){
	client, err := elastic.NewClient()
	if err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	var pdt []model.TProduct
	if err := model.DB.Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	for i ,pdtitem := range pdt{
		_, err := client.Index().Index("pdt").Type("doc").Id(strconv.Itoa(i)).BodyJson(pdtitem).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
	}
	_, err = client.Flush().Index("pdt").Do(context.Background())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"state": true,
	})
	return
}

func EsIndexArea(c *gin.Context){
	client, err := elastic.NewClient()
	if err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	var areas []model.Area
	if err := model.DB.Find(&areas).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	for i ,areaitem := range areas{
		_, err := client.Index().Index("area").Type("doc").Id(strconv.Itoa(i)).BodyJson(areaitem).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
	}
	_, err = client.Flush().Index("area").Do(context.Background())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"state": true,
	})
	return
}

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

func SerchArea(c *gin.Context){
	client, err := elastic.NewClient()
	if err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	termQuery := elastic.NewTermQuery("areaName", c.Query("name"))
	searchResult, err := client.Search().Index("area").Query(termQuery).From(0).Size(10).Pretty(true).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	common.SendResponse(searchResult,c)
}