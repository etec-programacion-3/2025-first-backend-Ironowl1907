package models

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model

	Title         string `gorm:"not null; unique_index"`
	Author        string
	Isbn          string `gorm:"unique_index"`
	Category      string
	State         string `gorm:"default:Available"` // Available, Unavailable
	Creation_date time.Time
}
