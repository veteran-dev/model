package model

import "time"

type AccoutLoginLog struct {
	ID           uint `gorm:"primaryKey"`
	StatusCode   int
	StatusReason string
	AccountID    int
	IPAddress    string
	Browser      string
	OS           string
	CountryName  string
	CountryID    int
	CreatedAt    time.Time
}
