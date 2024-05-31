package consumer_test

import (
	"io"
	"os"
	"testing"

	"github.com/digeon-inc/royle/filter/consumer"
	"github.com/digeon-inc/royle/pipe"
	"github.com/google/go-cmp/cmp"
)

func TestExportToMarkdown(t *testing.T) {
	type args struct {
		output io.Writer
		title  string
		tables []pipe.Table
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				output: os.Stdout,
				title:  "MySQL documentation",
				tables: []pipe.Table{
					{
						TableName: "Table1",
						Comment:   "TableComment1",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName1",
								ColumnDefault:       "ColumnDefault1",
								IsNullable:          "IsNullable1",
								ColumnType:          "ColumnType1",
								Extra:               "Extra1",
								Comment:             "ColumnComment1",
								ReferencedTableName: "",
								ConstraintTypes:     "",
							},
						},
					},
					{
						TableName: "Table2",
						Comment:   "",
						Columns: []pipe.Column{
							{
								ColumnName:          "ColumnName2",
								ColumnDefault:       "ColumnDefault2",
								IsNullable:          "IsNullable2",
								ColumnType:          "ColumnType2",
								Extra:               "Extra2",
								Comment:             "",
								ReferencedTableName: "Table1",
								ConstraintTypes:     "FOREIGN KEY",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 期待される出力をファイルから読み取る
			expectedContent, err := os.ReadFile("expected_output.md")
			if err != nil {
				t.Fatalf("Failed to read expected output file: %v", err)
			}

			// テスト対象の関数を実行し、ファイルに書き込む
			actualFile, err := os.Create("actual_output.md")
			if err != nil {
				t.Fatalf("Failed to create actual output file: %v", err)
			}
			defer os.Remove("actual_output.md")
			defer actualFile.Close()

			err = consumer.ExportToMarkdown(actualFile, tt.args.title, tt.args.tables)
			if err != nil {
				t.Fatalf("ExportToMarkdown() failed: %v", err)
			}

			// 期待される出力とファイルの内容を比較する
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
