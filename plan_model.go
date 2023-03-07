package model

import "gorm.io/gorm"

//	一个服务可以有多个套餐
//	面向商家
//	套餐计划定价额度表
type Plan struct {
	Name   string `gorm:"default:'plan'"`
	Status bool   `gorm:"default:true"`
	Credit int    `gorm:"default:0"`
	AppID  int    `gorm:"default:0"`
	App    App
	gorm.Model
}
