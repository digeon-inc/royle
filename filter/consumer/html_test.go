package consumer_test

import (
	"io"
	"os"
	"testing"

	"gitlab.com/digeon-inc/templates/open-mysql/filter/consumer"
	"gitlab.com/digeon-inc/templates/open-mysql/pipe"
)

func TestExportToHTML(t *testing.T) {

	type args struct {
		output io.Writer
		tables []pipe.TableMetaData
	}
	tests := []struct {
		name   string
		args   args
		hasErr bool
	}{
		{
			name: "success",
			args: args{
				output: os.Stdout,
				tables: []pipe.TableMetaData{
					{
						TableName: "table1",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName1",
								ColumnDefault:       "ColumnDefault1",
								IsNullable:          "IsNullable1",
								ColumnType:          "ColumnType1",
								Extra:               "Extra1",
								ReferencedTableName: "ReferencedTableName1",
								ConstraintTypes:     "ConstraintTypes1",
							},
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

			if err := consumer.ExportToHTML(tt.args.output, tt.args.tables); (err != nil) != tt.hasErr {
				t.Errorf("ExportToHTML() error = %v, hasErr %v", err, tt.hasErr)
			}
		})
	}
}
