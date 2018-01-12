package model

import (
	"time"
)

type TProduct struct {
	Guid string `gorm:"column:guid" json:"guid"`
	Sellerid string `gorm:"column:sellerId" json:"sellerId"`
	Psnapshotid string `gorm:"column:pSnapshotId" json:"pSnapshotId"`
	Pdttypeid string `gorm:"column:pdtTypeId" json:"pdtTypeId"`
	Pdttypename string `gorm:"column:pdtTypeName" json:"pdtTypeName"`
	Otherpropvalueids string `gorm:"column:otherPropValueIds" json:"otherPropValueIds"`
	Title string `gorm:"column:title" json:"title"`
	Sellingpoint string `gorm:"column:sellingPoint" json:"sellingPoint"`
	Pdtquality string `gorm:"column:pdtQuality" json:"pdtQuality"`
	Sourcecode string `gorm:"column:sourceCode" json:"sourceCode"`
	Sourceplace string `gorm:"column:sourcePlace" json:"sourcePlace"`
	Purchasetime time.Time `gorm:"column:purchaseTime" json:"purchaseTime"`
	Packaged string `gorm:"column:packaged" json:"packaged"`
	Pdtformat string `gorm:"column:pdtFormat" json:"pdtFormat"`
	Spec string `gorm:"column:spec" json:"spec"`
	Minnum float64 `gorm:"column:minNum" json:"minNum"`
	Price float64 `gorm:"column:price" json:"price"`
	Shelvetime time.Time `gorm:"column:shelveTime" json:"shelveTime"`
	Offsaletime time.Time `gorm:"column:offSaleTime" json:"offSaleTime"`
	Slideimgs string `gorm:"column:slideImgs" json:"slideImgs"`
	Inspectstatus string `gorm:"column:inspectStatus" json:"inspectStatus"`
	Status string `gorm:"column:status" json:"status"`
	Checkstatus string `gorm:"column:checkStatus" json:"checkStatus"`
	Saleamount float64 `gorm:"column:saleAmount" json:"saleAmount"`
	Createby string `gorm:"column:createBy" json:"createBy"`
	Createtime time.Time `gorm:"column:createTime" json:"createTime"`
	Updateby string `gorm:"column:updateBy" json:"updateBy"`
	Updatetime time.Time `gorm:"column:updateTime" json:"updateTime"`
	Isdelete string `gorm:"column:isdelete" json:"isdelete"`
}

func (TProduct) TableName() string {
	return "t_product"
}