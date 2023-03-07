package model

import "gorm.io/gorm"

//账号服务表
type AccountService struct {
	AccountID     uint `gorm:"default:0"`
	PlanID        uint `gorm:"default:0"`
	AppID         int  `gorm:"uniqueIndex,autoIncrement:10000"`
	AppSecret     string
	ServiceCredit int `gorm:"default:0"` //账户服务余额
	Plan          Plan
	gorm.Model
}
