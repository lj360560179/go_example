package model

// 地区
type Area struct {
	AreaId      string `gorm:"primary_key" json:"areaId"`
	ParentId  	string  `json:"parentId"`
	Name        string  `gorm:"size:64"`
	Levels      int
	Pcapital    int
	GovCity     int
	MapPoint    string
	AreaFullName    string  `gorm:"size:64"`
	AreaAlias		string  `gorm:"size:64"`
	IsHot       int
}
func (Area) TableName() string {
	return "tsys_area"
}