package model

import "gorm.io/gorm"

type Account struct {
	Name           string
	Password       string
	Email          string `gorm:"uniqueIndex:idx_email,sort:desc"`
	Status         int8
	Region         string
	AccountService []AccountService
	AccoutLoginLog []AccoutLoginLog
	gorm.Model
}
