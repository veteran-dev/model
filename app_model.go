package model

import "gorm.io/gorm"

type App struct {
	Name   string
	Status bool
	gorm.Model
}
