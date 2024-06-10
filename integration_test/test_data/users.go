package testdata

type Users struct {
	ID    int   `gorm:"primaryKey;autoIncrement;comment:'Primary Key'"`
	Name  string `gorm:"size:255;not null;comment:'Name of the user'"`
	Email string `gorm:"size:255;not null;unique;comment:'Unique email address'"`
}
