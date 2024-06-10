package transformer_test

import (
	"testing"

	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
	"github.com/google/go-cmp/cmp"
)

func TestSortColumnByGorm(t *testing.T) {
	type args struct {
		got  []pipe.Table
		want []pipe.Table
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				got: []pipe.Table{
					{
						TableName: "user",
						Comment:   "user comment",
						Columns: []pipe.Column{
							{
								ColumnName: "user_type",
								ColumnType: "string",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "email_to_update",
								ColumnType: "string",
							},
							{
								ColumnName: "hashed_password",
								ColumnType: "string",
							},
							{
								ColumnName: "user_id",
								ColumnType: "string",
							},
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "deleted_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "email",
								ColumnType: "string",
							},
						},
					},
					{
						TableName: "user_detail",
						Comment:   "user detail comment",
						Columns: []pipe.Column{
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "name",
								ColumnType: "string",
							},
							{
								ColumnName:          "user_id",
								ColumnType:          "string",
								ReferencedTableName: "user",
								ConstraintTypes:     "FOREIGN KEY",
							},
							{
								ColumnName: "user_detail_id",
								ColumnType: "string",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
						},
					},
				},
				want: []pipe.Table{
					{
						TableName: "user",
						Comment:   "user comment",
						Columns: []pipe.Column{
							{
								ColumnName: "user_id",
								ColumnType: "string",
							},
							{
								ColumnName: "email",
								ColumnType: "string",
							},
							{
								ColumnName: "email_to_update",
								ColumnType: "string",
							},
							{
								ColumnName: "hashed_password",
								ColumnType: "string",
							},
							{
								ColumnName: "user_type",
								ColumnType: "string",
							},
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "deleted_at",
								ColumnType: "time.Time",
							},
						},
					},
					{
						TableName: "user_detail",
						Comment:   "user detail comment",
						Columns: []pipe.Column{
							{
								ColumnName: "user_detail_id",
								ColumnType: "string",
							},
							{
								ColumnName:          "user_id",
								ColumnType:          "string",
								ReferencedTableName: "user",
								ConstraintTypes:     "FOREIGN KEY",
							},
							{
								ColumnName: "name",
								ColumnType: "string",
							},
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := transformer.SortColumnByGorm(tt.args.got, "test_data")
			if err != nil {
				t.Errorf("SortColumnByGorm error = %v", err)
			}
			if !cmp.Equal(actual, tt.args.want) {
				t.Errorf("diff =%v", cmp.Diff(actual, tt.args.want))
			}
		})
	}
}
