package model

import (
	"time"
)

type TProduct struct {
	Guid string `gorm:"column:guid" json:"guid"`
	Sellerid string `gorm:"column:sellerId" json:"sellerid"`
	Psnapshotid string `gorm:"column:pSnapshotId" json:"psnapshotid"`
	Pdttypeid string `gorm:"column:pdtTypeId" json:"pdttypeid"`
	Pdttypename string `gorm:"column:pdtTypeName" json:"pdttypename"`
	Otherpropvalueids string `gorm:"column:otherPropValueIds" json:"otherpropvalueids"`
	Title string `gorm:"column:title" json:"title"`
	Sellingpoint string `gorm:"column:sellingPoint" json:"sellingpoint"`
	Pdtquality string `gorm:"column:pdtQuality" json:"pdtquality"`
	Sourcecode string `gorm:"column:sourceCode" json:"sourcecode"`
	Sourceplace string `gorm:"column:sourcePlace" json:"sourceplace"`
	Purchasetime time.Time `gorm:"column:purchaseTime" json:"purchasetime"`
	Packaged string `gorm:"column:packaged" json:"packaged"`
	Pdtformat string `gorm:"column:pdtFormat" json:"pdtformat"`
	Spec string `gorm:"column:spec" json:"spec"`
	Minnum float64 `gorm:"column:minNum" json:"minnum"`
	Price float64 `gorm:"column:price" json:"price"`
	Shelvetime time.Time `gorm:"column:shelveTime" json:"shelvetime"`
	Offsaletime time.Time `gorm:"column:offSaleTime" json:"offsaletime"`
	Slideimgs string `gorm:"column:slideImgs" json:"slideimgs"`
	Inspectstatus string `gorm:"column:inspectStatus" json:"inspectstatus"`
	Status string `gorm:"column:status" json:"status"`
	Checkstatus string `gorm:"column:checkStatus" json:"checkstatus"`
	Saleamount float64 `gorm:"column:saleAmount" json:"saleamount"`
	Createby string `gorm:"column:createBy" json:"createby"`
	Createtime time.Time `gorm:"column:createTime" json:"createtime"`
	Updateby string `gorm:"column:updateBy" json:"updateby"`
	Updatetime time.Time `gorm:"column:updateTime" json:"updatetime"`
	Isdelete string `gorm:"column:isdelete" json:"isdelete"`
}

func (TProduct) TableName() string {
	return "t_product"
}