package pipe

type ColumnMetadata struct {
	TableName           string
	ColumnName          string
	ColumnDefault       string
	IsNullable          string
	ColumnType          string
	Extra               string
	ReferencedTableName string
	ConstraintTypes     string
}

type TableMetaData struct {
	TableName string
	Columns   []Column
}

type Column struct {
	ColumnName          string
	ColumnDefault       string
	IsNullable          string
	ColumnType          string
	Extra               string
	ReferencedTableName string
	ConstraintTypes     string
}
