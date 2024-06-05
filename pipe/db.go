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
	Comment             string
}

type TableMetadata struct {
	TableName    string
	TableComment string
}
