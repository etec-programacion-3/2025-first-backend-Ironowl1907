package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Title         string `gorm:"not null; unique_index"`
	Author        string
	Isbn          string
	Category      string
	State         string `gorm:"default:Available"` // Available, Unavailable
	Creation_date time.Time
}
