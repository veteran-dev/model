package model

import "gorm.io/gorm"

//	一个服务可以有多个套餐
//	面向商家
//	套餐计划定价额度表
type Plan struct {
	Name   string
	Status bool
	Credit int
	AppID  App
	gorm.Model
}
