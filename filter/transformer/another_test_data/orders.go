package model

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	OrderID     string `gorm:"primaryKey;size:30;not null"`
	TotalAmount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
