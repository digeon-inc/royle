package integration_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	another_testdata "github.com/digeon-inc/royle/integration_test/another_test_data"
	testdata "github.com/digeon-inc/royle/integration_test/test_data"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	informationSchemaDB *sql.DB
)

func init() {

	db, err := gorm.Open(mysql.Open(DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Set("gorm:table_options", "COMMENT='Stores basic information about users'").AutoMigrate(&testdata.Users{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Set("gorm:table_options", "COMMENT='Stores basic information about users details'").AutoMigrate(&testdata.UserDetail{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Set("gorm:table_options", "COMMENT='Stores basic information about orders'").AutoMigrate(&another_testdata.Orders{})
	if err != nil {
		log.Fatal(err)
	}

	informationSchemaDB, err = sql.Open("mysql", INFORMATION_SCHEMA_DSN())
	if err != nil {
		log.Fatal(err)
	}
}

// config

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	) + "?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True&loc=Asia%2FTokyo"
}

func INFORMATION_SCHEMA_DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		"information_schema",
	) + "?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True&loc=Asia%2FTokyo"
}

func DBName() string {
	return os.Getenv("DB_NAME")
}
