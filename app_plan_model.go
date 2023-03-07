package model

import "gorm.io/gorm"

//	一个服务可以有多个套餐
//	面向商家
//	服务套餐计划定价额度表
type AppPlan struct {
	Name   string `gorm:"default:'plan'"`
	Status bool   `gorm:"default:true"`
	Credit int    `gorm:"default:0"`
	AppID  int    `gorm:"default:0"`
	App    App
	gorm.Model
}
