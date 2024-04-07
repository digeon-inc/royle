package consumer_test

import (
	"io"
	"os"
	"testing"

	"github.com/digeon-inc/royle/filter/consumer"
	"github.com/digeon-inc/royle/pipe"
)

func TestExportToMarkdown(t *testing.T) {

	type args struct {
		output io.Writer
		title  string
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
				title:  "MySQL documentation",
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
								Comment:             "Comment1",
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
			// TODO: ファイルの差分をとって内容が変わってないかチェックする。
			if err := consumer.ExportToMarkdown(tt.args.output, tt.args.title, tt.args.tables); (err != nil) != tt.hasErr {
				t.Errorf("ExportToMarkdown() error = %v, hasErr %v", err, tt.hasErr)
			}
		})
	}
}
