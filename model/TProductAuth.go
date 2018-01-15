package model

import "time"

type TProductAuth struct {
	Guid string `gorm:"column:guid" json:"guid"`
	Sellerid string `gorm:"column:sellerId" json:"sellerId"`
	Pid string `gorm:"column:pId" json:"pId"`
	Snapshotid string `gorm:"column:snapshotId" json:"snapshotId"`
	Checkstatus string `gorm:"column:checkStatus" json:"checkStatus"`
	Checkinfo string `gorm:"column:checkInfo" json:"checkInfo"`
	Checkby string `gorm:"column:checkBy" json:"checkBy"`
	Checktime time.Time `gorm:"column:checkTime" json:"checkTime"`
}

func (TProductAuth) TableName() string {
	return "t_product_auth"
}