package transformer

import (
	"github.com/digeon-inc/royle/pipe"
)

// columnsがテーブル順であることが前提
func MergeMetadataIntoTables(cols []pipe.ColumnMetadata, tables []pipe.TableMetadata) []pipe.Table {
	result := make([]pipe.Table, 0, 100)
	currentTableName := ""
	currentColumns := make([]pipe.Column, 0, 20)

	for i, col := range cols {
		// カラムはテーブル順なので、テーブル名が変わったら、変わる前のテーブルの全カラムが取り出されたことを意味する。
		if currentTableName != col.TableName {
			if currentTableName != "" {
				addTableToResult(&result, currentTableName, currentColumns, tables)
			}
			currentTableName = col.TableName
			currentColumns = make([]pipe.Column, 0, 50)
		}
		currentColumns = append(currentColumns, pipe.Column{
			ColumnName:          col.ColumnName,
			ColumnDefault:       col.ColumnDefault,
			IsNullable:          col.IsNullable,
			ColumnType:          col.ColumnType,
			Extra:               col.Extra,
			Comment:             col.Comment,
			ReferencedTableName: col.ReferencedTableName,
			ConstraintTypes:     col.ConstraintTypes,
		})

		// 最後のテーブルは次のテーブルと比較できず追加されないから、明示的に追加する。
		if i == len(cols)-1 {
			addTableToResult(&result, currentTableName, currentColumns, tables)
		}

	}
	return result
}

func findMatchingTable(tables []pipe.TableMetadata, targetTableName string) *pipe.TableMetadata {
	for _, table := range tables {
		if table.TableName == targetTableName {
			return &table
		}
	}
	return nil
}

func addTableToResult(result *[]pipe.Table, tableName string, columns []pipe.Column, tables []pipe.TableMetadata) {
	if currentTableMetadata := findMatchingTable(tables, tableName); currentTableMetadata != nil {
		*result = append(*result, pipe.Table{
			TableName: currentTableMetadata.TableName,
			Comment:   currentTableMetadata.TableComment,
			Columns:   columns,
		})
	} else {
		*result = append(*result, pipe.Table{
			TableName: tableName,
			Columns:   columns,
		})
	}
}
