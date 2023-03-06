package model

import "gorm.io/gorm"

// 账号和服务的关系是一对多的关系
type App struct {
	Name   string
	Credit uint
	Status bool
	gorm.Model
}
