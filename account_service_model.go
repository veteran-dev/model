package model

import "gorm.io/gorm"

type AccountService struct {
	AccountID uint
	ServiceID uint
	AppID     string `gorm:"uniqueIndex"`
	AppSecret string
	Credit    int
	Service   App
	gorm.Model
}
