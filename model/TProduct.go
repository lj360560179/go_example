package model

import (
	"time"
)

type TProduct struct {
	Guid string `gorm:"column:guid" json:"guid"`
	Sellerid string `gorm:"column:sellerid" json:"sellerid"`
	Psnapshotid string `gorm:"column:psnapshotid" json:"psnapshotid"`
	Pdttypeid string `gorm:"column:pdttypeid" json:"pdttypeid"`
	Pdttypename string `gorm:"column:pdttypename" json:"pdttypename"`
	Otherpropvalueids string `gorm:"column:otherpropvalueids" json:"otherpropvalueids"`
	Title string `gorm:"column:title" json:"title"`
	Sellingpoint string `gorm:"column:sellingpoint" json:"sellingpoint"`
	Pdtquality string `gorm:"column:pdtquality" json:"pdtquality"`
	Sourcecode string `gorm:"column:sourcecode" json:"sourcecode"`
	Sourceplace string `gorm:"column:sourceplace" json:"sourceplace"`
	Purchasetime time.Time `gorm:"column:purchasetime" json:"purchasetime"`
	Packaged string `gorm:"column:packaged" json:"packaged"`
	Pdtformat string `gorm:"column:pdtformat" json:"pdtformat"`
	Spec string `gorm:"column:spec" json:"spec"`
	Minnum float64 `gorm:"column:minnum" json:"minnum"`
	Price float64 `gorm:"column:price" json:"price"`
	Shelvetime time.Time `gorm:"column:shelvetime" json:"shelvetime"`
	Offsaletime time.Time `gorm:"column:offsaletime" json:"offsaletime"`
	Slideimgs string `gorm:"column:slideimgs" json:"slideimgs"`
	Inspectstatus string `gorm:"column:inspectstatus" json:"inspectstatus"`
	Status string `gorm:"column:status" json:"status"`
	Checkstatus string `gorm:"column:checkstatus" json:"checkstatus"`
	Saleamount float64 `gorm:"column:saleamount" json:"saleamount"`
	Createby string `gorm:"column:createby" json:"createby"`
	Createtime time.Time `gorm:"column:createtime" json:"createtime"`
	Updateby string `gorm:"column:updateby" json:"updateby"`
	Updatetime time.Time `gorm:"column:updatetime" json:"updatetime"`
	Isdelete string `gorm:"column:isdelete" json:"isdelete"`
}

func (TProduct) TableName() string {
	return "t_product"
}