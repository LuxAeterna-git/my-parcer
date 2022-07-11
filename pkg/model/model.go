package model

type Good struct {
	Id     uint `gorm:"primary_key"`
	Name   string
	Url    string
	UrlImg string
	Price  string
}
