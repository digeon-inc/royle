package integration_test

import (
	"os"
	"testing"

	"github.com/digeon-inc/royle/filter/consumer"
	"github.com/digeon-inc/royle/filter/producer"
	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm/schema"
)

func TestIntegration(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name             string
		args             args
		wantTables       []pipe.Table
		wantSortedTables []pipe.Table
	}{
		{
			name: "success",
			args: args{
				title: "MySQL documentation",
			},
			wantTables: []pipe.Table{
				{
					TableName: "orders",
					Comment:   "Stores basic information about orders",
					Columns: []pipe.Column{
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "product_name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "quantity",
							ColumnDefault:       "1",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							Comment:             "Quantity of the product being ordered, defaults to 1",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "user_id",
							ColumnDefault:       "",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "users",
							ConstraintTypes:     "FOREIGN KEY",
						},
					},
				},
				{
					TableName: "users",
					Comment:   "Stores basic information about users",
					Columns: []pipe.Column{
						{
							ColumnName:          "email",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "UNIQUE",
						},
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
					},
				},
			},
			wantSortedTables: []pipe.Table{
				{
					TableName: "orders",
					Comment:   "Stores basic information about orders",
					Columns: []pipe.Column{
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "product_name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "user_id",
							ColumnDefault:       "",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "users",
							ConstraintTypes:     "FOREIGN KEY",
						},
						{
							ColumnName:          "quantity",
							ColumnDefault:       "1",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							Comment:             "Quantity of the product being ordered, defaults to 1",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
					},
				},
				{
					TableName: "users",
					Comment:   "Stores basic information about users",
					Columns: []pipe.Column{
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "email",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							Comment:             "",
							ReferencedTableName: "",
							ConstraintTypes:     "UNIQUE",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			columnSource, err := producer.FetchColumnMetadata(informationSchemaDB, DBName())
			if err != nil {
				t.Errorf("FetchColumnMetadata error = %v", err)
			}

			tableSource, err := producer.FetchTableMetadata(informationSchemaDB, DBName())
			if err != nil {
				t.Errorf("FetchTableMetadata error = %v", err)
			}
			t.Log(tableSource)

			gotTables := transformer.MergeMetadataIntoTables(columnSource, tableSource)
			if !cmp.Equal(gotTables, tt.wantTables) {
				t.Errorf("diff =%v", cmp.Diff(gotTables, tt.wantTables))
			}

			gotSortedTables, err := transformer.SortColumnByGormModelFile(schema.NamingStrategy{}, gotTables, []string{"test_data", "another_test_data"})
			if err != nil {
				t.Errorf("SortColumnByGormModelFile error = %v", err)
			}
			if !cmp.Equal(gotSortedTables, tt.wantSortedTables) {
				t.Errorf("diff =%v", cmp.Diff(gotSortedTables, tt.wantSortedTables))
			}

			// 生成したmdを書き込むファイルを作成する。
			actualFile, err := os.Create("actual_output.md")
			if err != nil {
				t.Fatalf("Failed to create actual output file: %v", err)
			}
			defer os.Remove("actual_output.md")
			defer actualFile.Close()

			err = consumer.ExportToMarkdown(actualFile, tt.args.title, gotSortedTables)
			if err != nil {
				t.Fatalf("ExportToMarkdown error: %v", err)
			}

			// 期待される出力と実際に書き込まれたファイルの内容を比較する。
			expectedContent, err := os.ReadFile("expected_output.md")
			if err != nil {
				t.Fatalf("Failed to read expected output file: %v", err)
			}
			actualContent, err := os.ReadFile("actual_output.md")
			if err != nil {
				t.Fatalf("Failed to read actual output file: %v", err)
			}
			if diff := cmp.Diff(string(expectedContent), string(actualContent)); diff != "" {
				t.Errorf("Mismatch between expected and actual output (-want +got):\n%s", diff)
			}

		})
	}
}
