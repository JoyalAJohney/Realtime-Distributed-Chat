package database

import (
	"time"

	"gorm.io/gorm"
)

type DBUser struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(100);not null"`
}

type DBMessage struct {
	gorm.Model
	UserID    string `gorm:"type:varchar(100);not null"`
	RoomID    string `gorm:"type:varchar(100);not null"`
	Message   string `gorm:"type:varchar(100);not null"`
	Timestamp *time.Time
}
