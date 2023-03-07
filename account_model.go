package model

import "gorm.io/gorm"

type Account struct {
	Name           string `gorm:"default:'account'"`
	Password       string
	Email          string `gorm:"uniqueIndex:idx_email,sort:desc"`
	Status         int8   `gorm:"default:0"`
	Region         string `gorm:"default:'未知'"`
	AccountService []AccountService
	AccoutLoginLog []AccoutLoginLog
	gorm.Model
}
