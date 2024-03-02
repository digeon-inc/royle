package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"gitlab.com/digeon-inc/templates/open-mysql/filter/consumer"
	"gitlab.com/digeon-inc/templates/open-mysql/filter/producer"
	"gitlab.com/digeon-inc/templates/open-mysql/filter/transformer"
)

const (
	HTML     = "html"
	MARKDOWN = "md"
	STDOUT   = "stdout"
)

var (
	outputFileFormat string
	outputFileName   string
	dbUser           string
	dbPassword       string
	dbHost           string
	dbPort           string
	dbName           string
)

var rootCmd = &cobra.Command{
	Use:   "open-mysql",
	Short: "Generates documentation for the MySQL tables.",
	Long:  "This is a command-line application written in Go that connects to a MySQL database, extracts table information, and generates a file documenting the database tables.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("mysql", INFORMATION_SCHEMA_DSN())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		source, err := producer.FetchColumnMetadata(db, DBName())
		if err != nil {
			log.Fatal(err)
		}

		tables := transformer.ConvertColumnMetadataToTableMetaData(source)

		file, err := os.Create(OutputFileName())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		switch OutputFileFormat() {
		case HTML:
			err = consumer.ExportToHTML(file, tables)
		case MARKDOWN:
			err = consumer.ExportToMarkdown(file, tables)
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputFileFormat, "format", "f", MARKDOWN, "output file format. Choose either md or html.")
	if outputFileFormat != HTML && outputFileFormat != MARKDOWN {
		log.Fatalf("%s is unavailable", outputFileFormat)
	}
	rootCmd.Flags().StringVarP(&outputFileName, "filename", "o", "output", "output file name")

	rootCmd.Flags().StringVarP(&dbUser, "user", "u", "", "mysql user")
	if err := rootCmd.MarkFlagRequired("user"); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().StringVarP(&dbPassword, "password", "p", "", "mysql password")
	if err := rootCmd.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().StringVarP(&dbHost, "host", "s", "", "mysql host")
	if err := rootCmd.MarkFlagRequired("host"); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().StringVarP(&dbPort, "port", "r", "", "mysql port")
	if err := rootCmd.MarkFlagRequired("port"); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().StringVarP(&dbName, "dbname", "n", "", "Database name you want to know about the table ")
	if err := rootCmd.MarkFlagRequired("dbname"); err != nil {
		log.Fatal(err)
	}

}
