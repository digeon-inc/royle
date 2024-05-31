package transformer

import "github.com/digeon-inc/royle/pipe"


// columnsがテーブル順であることが前提
func ConvertColumnMetadataToTableMetaData(cols []pipe.ColumnMetadata, tables []pipe.TableMetadata) []pipe.Table {
	result := make([]pipe.Table, 0, 100)
	currentTableName := ""
	currentColumns := make([]pipe.Column, 0, 20)
	for i, col := range cols {
		if currentTableName != col.TableName {
			if currentTableName != "" {
				result = append(result, pipe.Table{TableName: currentTableName, Columns: currentColumns})
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

		if i == len(cols)-1 {
			currentTableMetadata := findMatchingTable(tables,currentTableName) 
			result = append(result, pipe.Table{TableName: currentTableMetadata.TableName,Comment: currentTableMetadata.TableComment, Columns: currentColumns})
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
