package integration_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	informationSchemaDB *sql.DB
)

func init() {
	db, err := sql.Open("mysql", DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE
	);
	`

	if _, err = db.Exec(createUsersTableSQL); err != nil {
		log.Fatal(err)
	}

	createOrdersTableSQL := `
	CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		product_name VARCHAR(255) NOT NULL,
		user_id INT,
		quantity INT DEFAULT 1,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	if _, err = db.Exec(createOrdersTableSQL); err != nil {
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
