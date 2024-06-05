package producer

import (
	"database/sql"
	_ "embed"

	"github.com/digeon-inc/royle/pipe"
)

//go:embed table.sql
var tableSQL string // table.sql ファイルをバイナリに埋め込む

func FetchTableMetadata(db *sql.DB, schemaName string) ([]pipe.TableMetadata, error) {
	rows, err := db.Query(tableSQL, schemaName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]pipe.TableMetadata, 0, 100)
	var (
		tableName    sql.NullString
		tableComment sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(&tableName, &tableComment); err != nil {
			return nil, err
		}
		result = append(result, pipe.TableMetadata{
			TableName:    tableName.String,
			TableComment: tableComment.String,
		})
	}

	return result, nil
}
