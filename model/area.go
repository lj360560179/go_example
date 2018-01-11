package model

import "time"

// Category 话题的分类
type Area struct {
	AreaId        uint       `gorm:"primary_key" json:"id"`
	ParentId  uint  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Name      string     `json:"name"`
	Sequence  int        `json:"sequence"` //同级别的分类可根据sequence的值来排序
	ParentID  int        `json:"parentId"` //直接父分类的ID
}
