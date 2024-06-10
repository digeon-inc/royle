package model

import (
	"time"
)

type UserDetail struct {
	UserDetailID string `gorm:"primaryKey;size:30;not null"`
	UserID       string
	Name         string `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
