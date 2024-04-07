package producer

import (
	"database/sql"
	_ "embed"

	"github.com/digeon-inc/royle/pipe"
)

//go:embed query.sql
var querySQL string // query.sql ファイルをバイナリに埋め込む

func FetchColumnMetadata(db *sql.DB, schemaName string) ([]pipe.ColumnMetadata, error) {
	rows, err := db.Query(querySQL, schemaName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]pipe.ColumnMetadata, 0, 1000)
	var (
		tableName           sql.NullString
		columnName          sql.NullString
		columnDefault       sql.NullString
		isNullable          sql.NullString
		columnType          sql.NullString
		extra               sql.NullString
		referencedTableName sql.NullString
		constraintTypes     sql.NullString
		comment             sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(&tableName, &columnName, &columnDefault, &isNullable, &columnType, &extra, &comment, &referencedTableName, &constraintTypes); err != nil {
			return nil, err
		}
		result = append(result, pipe.ColumnMetadata{
			TableName:           tableName.String,
			ColumnName:          columnName.String,
			ColumnDefault:       columnDefault.String,
			IsNullable:          isNullable.String,
			ColumnType:          columnType.String,
			Extra:               extra.String,
			Comment:             comment.String,
			ReferencedTableName: referencedTableName.String,
			ConstraintTypes:     constraintTypes.String,
		})
	}

	return result, nil
}
