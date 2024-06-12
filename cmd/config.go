package cmd

import (
	"fmt"
)

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		database,
	) + "?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True"
}

func INFORMATION_SCHEMA_DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		"information_schema",
	) + "?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True"
}

func DatabaseName() string {
	return database
}

func Title() string {
	return title
}
