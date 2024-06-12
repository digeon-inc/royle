package testdata

import "time"

type Users struct {
	ID    int32  `gorm:"primaryKey;autoIncrement;"`
	Name  string `gorm:"size:255;not null;"`
	Email string `gorm:"size:255;not null;unique;"`
}

type UserDetail struct {
	UserDetailID string `gorm:"primaryKey;size:30;not null"`
	UserID       int32  // 外部キー
	User         Users  `gorm:"foreignKey:UserID"` // Usersとの関連
	Name         string `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
