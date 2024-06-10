package anothertestdata

import testdata "github.com/digeon-inc/royle/integration_test/test_data"

type Orders struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:'Primary Key'"`
	ProductName string         `gorm:"size:255;not null;comment:'Name of the product'"`
	UserID      uint           `gorm:"index;comment:'Foreign key to the user'"`
	Quantity    int            `gorm:"default:1;comment:'Quantity of the product being ordered, defaults to 1'"`
	User        testdata.Users `gorm:"foreignKey:UserID"`
}
