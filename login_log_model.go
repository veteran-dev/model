package model

import "time"

type AccoutLoginLog struct {
	ID           uint `gorm:"primaryKey"`
	StatusCode   int  `gorm:"default:200"`
	StatusReason string
	AccountID    int `gorm:"default:0"`
	IPAddress    string
	Browser      string
	OS           string
	CountryName  string
	CountryID    int
	CreatedAt    time.Time
}
