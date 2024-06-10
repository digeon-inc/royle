package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID         string      `gorm:"primaryKey;size:30;not null"`
	Email          string      `gorm:"size:255;unique;not null;index"`
	EmailToUpdate  string      `gorm:"size:255"`
	HashedPassword string      `gorm:"size:255;not null"`
	UserType       string      `gorm:"size:255;not null"`
	UserDetail     *UserDetail `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt gorm.DeletedAt
}
