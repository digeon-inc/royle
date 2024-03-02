package transformer

import "gitlab.com/digeon-inc/templates/open-mysql/pipe"

// columnsがテーブル順であることが前提
func ConvertColumnMetadataToTableMetaData(cols []pipe.ColumnMetadata) []pipe.TableMetaData {
	result := make([]pipe.TableMetaData, 0, 100)
	currentTableName := ""
	currentColumns := make([]pipe.Column, 0, 20)
	for i, col := range cols {
		if currentTableName != col.TableName {
			if currentTableName != "" {
				result = append(result, pipe.TableMetaData{TableName: currentTableName, Columns: currentColumns})
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
			ReferencedTableName: col.ReferencedTableName,
			ConstraintTypes:     col.ConstraintTypes,
		})

		if i == len(cols)-1 {
			result = append(result, pipe.TableMetaData{TableName: currentTableName, Columns: currentColumns})
		}

	}
	return result
}
