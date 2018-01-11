package model

type Animal struct {
	ID int64
	Name string `gorm:"default:'galeone'"`
	Age int64
}
