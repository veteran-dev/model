package model

import "gorm.io/gorm"

type App struct {
	Name   string `gorm:"default:'app'"`
	Status bool   `gorm:"default:true"`
	gorm.Model
}
