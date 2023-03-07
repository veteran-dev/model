package model

import "gorm.io/gorm"

//账号服务表
type AccountService struct {
	AccountID uint
	PlanID    uint
	AppID     int `gorm:"uniqueIndex,autoIncrement:10000"`
	AppSecret string
	Credit    int
	Plan      Plan
	gorm.Model
}
