package pipe

type Table struct {
	TableName string
	Comment   string
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
	Comment             string
}
