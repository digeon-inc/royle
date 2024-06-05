package transformer_test

import (
	"testing"

	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
	"github.com/google/go-cmp/cmp"
)

func TestMergeMetadataIntoTables(t *testing.T) {
	type args struct {
		columns []pipe.ColumnMetadata
		tables  []pipe.TableMetadata
		want    []pipe.Table
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				// columnはテーブル順が必須である。
				columns: []pipe.ColumnMetadata{
					{
						TableName:           "TableName1",
						ColumnName:          "ColumnName1-1",
						ColumnDefault:       "ColumnDefault1-1",
						IsNullable:          "IsNullable1-1",
						ColumnType:          "ColumnType1-1",
						Extra:               "Extra1-1",
						Comment:             "ColumnComment1-1",
						ReferencedTableName: "",
						ConstraintTypes:     "",
					},
					{
						TableName:           "TableName1",
						ColumnName:          "ColumnName1-2",
						ColumnDefault:       "ColumnDefault1-2",
						IsNullable:          "IsNullable1-2",
						ColumnType:          "ColumnType1-2",
						Extra:               "Extra1-2",
						Comment:             "",
						ReferencedTableName: "",
						ConstraintTypes:     "",
					},
					{
						TableName:           "TableName2",
						ColumnName:          "ColumnName2",
						ColumnDefault:       "ColumnDefault2",
						IsNullable:          "IsNullable2",
						ColumnType:          "ColumnType2",
						Extra:               "Extra2",
						Comment:             "Comment2",
						ReferencedTableName: "Table1",
						ConstraintTypes:     "ConstraintTypes2",
					},
					{
						TableName:           "TableName3",
						ColumnName:          "ColumnName3",
						ColumnDefault:       "ColumnDefault3",
						IsNullable:          "IsNullable3",
						ColumnType:          "ColumnType3",
						Extra:               "Extra3",
						Comment:             "Comment3",
						ReferencedTableName: "Table1",
						ConstraintTypes:     "",
					},
				},
				// tableは順不同。
				tables: []pipe.TableMetadata{
					{
						TableName:    "TableName3",
						TableComment: "TableComment3",
					},
					{
						TableName:    "TableName1",
						TableComment: "TableComment1",
					},
					{
						TableName:    "TableName2",
						TableComment: "",
					},
				},
				want: []pipe.Table{
					{
						TableName: "TableName1",
						Comment:   "TableComment1",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName1-1",
								ColumnDefault:       "ColumnDefault1-1",
								IsNullable:          "IsNullable1-1",
								ColumnType:          "ColumnType1-1",
								Extra:               "Extra1-1",
								Comment:             "ColumnComment1-1",
								ReferencedTableName: "",
								ConstraintTypes:     "",
							},
							{
								ColumnName:          "ColumnName1-2",
								ColumnDefault:       "ColumnDefault1-2",
								IsNullable:          "IsNullable1-2",
								ColumnType:          "ColumnType1-2",
								Extra:               "Extra1-2",
								Comment:             "",
								ReferencedTableName: "",
								ConstraintTypes:     "",
							},
						},
					},
					{
						TableName: "TableName2",
						Comment:   "",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName2",
								ColumnDefault:       "ColumnDefault2",
								IsNullable:          "IsNullable2",
								ColumnType:          "ColumnType2",
								Extra:               "Extra2",
								Comment:             "Comment2",
								ReferencedTableName: "Table1",
								ConstraintTypes:     "ConstraintTypes2",
							},
						},
					},
					{
						TableName: "TableName3",
						Comment:   "TableComment3",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName3",
								ColumnDefault:       "ColumnDefault3",
								IsNullable:          "IsNullable3",
								ColumnType:          "ColumnType3",
								Extra:               "Extra3",
								Comment:             "Comment3",
								ReferencedTableName: "Table1",
								ConstraintTypes:     "",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformer.MergeMetadataIntoTables(tt.args.columns, tt.args.tables)
			if !cmp.Equal(got, tt.args.want) {
				t.Errorf("diff =%v", cmp.Diff(got, tt.args.want))
			}
		})
	}
}
