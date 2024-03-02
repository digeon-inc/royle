package integration_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/digeon-inc/royle/filter/producer"
	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
)

func TestFetchColumnMetadata(t *testing.T) {

	tests := []struct {
		name   string
		want   []pipe.TableMetaData
		hasErr bool
	}{
		{
			name: "success",
			want: []pipe.TableMetaData{
				{
					TableName: "orders",
					Columns: []pipe.Column{
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "product_name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "quantity",
							ColumnDefault:       "1",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
						{
							ColumnName:          "user_id",
							ColumnDefault:       "",
							IsNullable:          "YES",
							ColumnType:          "int",
							Extra:               "",
							ReferencedTableName: "users",
							ConstraintTypes:     "FOREIGN KEY",
						},
					},
				},
				{
					TableName: "users",
					Columns: []pipe.Column{
						{
							ColumnName:          "email",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							ReferencedTableName: "",
							ConstraintTypes:     "UNIQUE",
						},
						{
							ColumnName:          "id",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "int",
							Extra:               "auto_increment",
							ReferencedTableName: "",
							ConstraintTypes:     "PRIMARY KEY",
						},
						{
							ColumnName:          "name",
							ColumnDefault:       "",
							IsNullable:          "NO",
							ColumnType:          "varchar(255)",
							Extra:               "",
							ReferencedTableName: "",
							ConstraintTypes:     "",
						},
					},
				},
			},
			hasErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			source, err := producer.FetchColumnMetadata(informationSchemaDB, DBName())
			if (err != nil) != tt.hasErr {
				t.Errorf("FetchColumnMetadata error = %v, hasErr %v", err, tt.hasErr)
			}
			got := transformer.ConvertColumnMetadataToTableMetaData(source)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff =%v", cmp.Diff(got, tt.want))
			}
		})
	}
}
