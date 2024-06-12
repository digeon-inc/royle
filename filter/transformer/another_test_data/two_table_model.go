package model

type Region struct {
	RegionID   uint   `gorm:"primaryKey;autoIncrement;comment:Unique identifier for each geographical region"`
	RegionName string `gorm:"size:50;not null;comment:Name of the geographical region"`
	Country    string `gorm:"size:50;not null;comment:Country where the region is located"`
}

type Shipper struct {
	ShipperID   uint   `gorm:"primaryKey;autoIncrement;comment:Unique identifier for each shipping company"`
	ShipperName string `gorm:"size:100;not null;comment:Name of the shipping company"`
	Phone       string `gorm:"size:20;not null;comment:Contact phone number for the shipping company"`
}
