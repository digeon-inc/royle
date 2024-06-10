package testdata

type Users struct {
	ID    int32  `gorm:"primaryKey;autoIncrement;"`
	Name  string `gorm:"size:255;not null;"`
	Email string `gorm:"size:255;not null;unique;"`
}
