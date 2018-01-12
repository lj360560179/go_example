package model


// 地区
type Area struct {
	AreaId      string `gorm:"column:areaId" json:"areaId" `
	ParentId  	string `gorm:"column:parentId" json:"parentId" `
	AreaName        string `gorm:"column:areaName" json:"name" `
	Levels      int `json:"levels"`
	Pcapital    int `json:"pcapital"`
	GovCity     int `gorm:"column:govcity" json:"govCity" `
	MapPoint    string `gorm:"column:column:mapPoint" json:"mapPoint" `
	AreaFullName    string `gorm:"column:areafullname" json:"areaFullName" `
	AreaAlias		string `gorm:"column:areaAlias" json:"areaAlias" `
	IsHot       int `gorm:"column:column:isHot" json:"isHot"`
}
func (Area) TableName() string {
	return "tsys_area"
}