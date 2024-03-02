package producer

import (
	"database/sql"

	"github.com/digeon-inc/royle/pipe"
)

func FetchColumnMetadata(db *sql.DB, schemaName string) ([]pipe.ColumnMetadata, error) {
	query := `
SELECT
	c.TABLE_NAME,
	c.COLUMN_NAME,
	c.COLUMN_DEFAULT,
	c.IS_NULLABLE,
	c.COLUMN_TYPE,
	c.EXTRA,
	MAX(k.REFERENCED_TABLE_NAME) AS REFERENCED_TABLE_NAME,
	GROUP_CONCAT(DISTINCT t.CONSTRAINT_TYPE) AS CONSTRAINT_TYPES
FROM
	COLUMNS c
LEFT OUTER JOIN
	KEY_COLUMN_USAGE k ON c.TABLE_NAME = k.TABLE_NAME AND c.COLUMN_NAME = k.COLUMN_NAME
LEFT OUTER JOIN
	TABLE_CONSTRAINTS t ON k.CONSTRAINT_NAME = t.CONSTRAINT_NAME AND k.TABLE_NAME = t.TABLE_NAME
WHERE
	c.TABLE_SCHEMA = ?
GROUP BY
	c.TABLE_NAME,
	c.COLUMN_NAME,
	c.COLUMN_DEFAULT,
	c.IS_NULLABLE,
	c.COLUMN_TYPE,
	c.EXTRA;
`
	rows, err := db.Query(query, schemaName)
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
	)

	for rows.Next() {
		if err := rows.Scan(&tableName, &columnName, &columnDefault, &isNullable, &columnType, &extra, &referencedTableName, &constraintTypes); err != nil {
			return nil, err
		}
		result = append(result, pipe.ColumnMetadata{
			TableName:           tableName.String,
			ColumnName:          columnName.String,
			ColumnDefault:       columnDefault.String,
			IsNullable:          isNullable.String,
			ColumnType:          columnType.String,
			Extra:               extra.String,
			ReferencedTableName: referencedTableName.String,
			ConstraintTypes:     constraintTypes.String,
		})
	}

	return result, nil
}
