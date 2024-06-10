package cmd

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/digeon-inc/royle/filter/consumer"
	"github.com/digeon-inc/royle/filter/producer"
	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	title      string
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	database   string
	dirs       []string
)

var rootCmd = &cobra.Command{
	Use:   "royle",
	Short: "Generates documentation for the MySQL tables.",
	Long:  "This is a command-line application written in Go that connects to a MySQL database, extracts table information, and generates a file documenting the database tables.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("mysql", INFORMATION_SCHEMA_DSN())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var (
			columnSource []pipe.ColumnMetadata
			tableSource  []pipe.TableMetadata
			wg           sync.WaitGroup
		)

		// ゴルーチンの開始
		wg.Add(2)
		go func() {
			defer wg.Done()
			var err error
			columnSource, err = producer.FetchColumnMetadata(db, DatabaseName())
			if err != nil {
				log.Fatal(err)
			}
		}()

		go func() {
			defer wg.Done()
			var err error
			tableSource, err = producer.FetchTableMetadata(db, DatabaseName())
			if err != nil {
				log.Fatal(err)
			}
		}()

		// すべてのゴルーチンの終了を待つ
		wg.Wait()

		tables := transformer.MergeMetadataIntoTables(columnSource, tableSource)

		// ディレクトリが指定されている場合のみ、カラムをソート
		if len(dirs) > 0 {
			tables, err = transformer.SortColumnByGormModelFile(tables, dirs)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err = consumer.ExportToMarkdown(os.Stdout, Title(), tables); err != nil {
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
	rootCmd.Flags().StringVarP(&title, "title", "t", "ROYLE", "document title")

	rootCmd.Flags().StringSliceVarP(&dirs, "dirs", "x", nil, "directories to search for GORM model files")

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

	rootCmd.Flags().StringVarP(&database, "database", "d", "", "mysql database name")
	if err := rootCmd.MarkFlagRequired("database"); err != nil {
		log.Fatal(err)
	}
}
