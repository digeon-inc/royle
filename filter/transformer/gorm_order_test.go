package transformer_test

import (
	"testing"

	"github.com/digeon-inc/royle/filter/transformer"
	"github.com/digeon-inc/royle/pipe"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm/schema"
)

func TestSortColumnByGorm(t *testing.T) {
	type args struct {
		tables []pipe.Table
		dirs   []string
	}
	tests := []struct {
		name string
		args args
		want []pipe.Table
	}{
		{
			name: "success",
			args: args{
				tables: []pipe.Table{
					{
						TableName: "users",
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
						TableName: "user_details",
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
								ReferencedTableName: "users",
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
				dirs: []string{"test_data"},
			},
			want: []pipe.Table{
				{
					TableName: "users",
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
					TableName: "user_details",
					Comment:   "user detail comment",
					Columns: []pipe.Column{
						{
							ColumnName: "user_detail_id",
							ColumnType: "string",
						},
						{
							ColumnName:          "user_id",
							ColumnType:          "string",
							ReferencedTableName: "users",
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
		{
			name: "success （カラムがMysqlのデータベース内には存在するが、ファイルには記述されてない場合）",
			args: args{
				tables: []pipe.Table{
					{
						TableName: "users",
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
								ColumnName: "mysql_only",
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
								ColumnName: "mysql_only2",
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
						TableName: "user_details",
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
								ReferencedTableName: "users",
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
				dirs: []string{"test_data"},
			},
			want: []pipe.Table{
				{
					TableName: "users",
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
						{
							ColumnName: "mysql_only",
							ColumnType: "string",
						},
						{
							ColumnName: "mysql_only2",
							ColumnType: "string",
						},
					},
				},
				{
					TableName: "user_details",
					Comment:   "user detail comment",
					Columns: []pipe.Column{
						{
							ColumnName: "user_detail_id",
							ColumnType: "string",
						},
						{
							ColumnName:          "user_id",
							ColumnType:          "string",
							ReferencedTableName: "users",
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
		{
			name: "success （テーブルがMysqlのデータベース内には存在するが、ファイルには記述されてない場合）",
			args: args{
				tables: []pipe.Table{
					{
						TableName: "users",
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
						TableName: "mysql_only",
						Comment:   "This table has not been written to a file",
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
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
						},
					},
					{
						TableName: "user_details",
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
								ReferencedTableName: "users",
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
				dirs: []string{"test_data"},
			},
			want: []pipe.Table{
				{
					TableName: "users",
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
					TableName: "mysql_only",
					Comment:   "This table has not been written to a file",
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
							ColumnName: "updated_at",
							ColumnType: "time.Time",
						},
					},
				},
				{
					TableName: "user_details",
					Comment:   "user detail comment",
					Columns: []pipe.Column{
						{
							ColumnName: "user_detail_id",
							ColumnType: "string",
						},
						{
							ColumnName:          "user_id",
							ColumnType:          "string",
							ReferencedTableName: "users",
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
		{
			name: "success （ディレクトリが複数の場合）",
			args: args{
				tables: []pipe.Table{
					{
						TableName: "users",
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
						TableName: "user_details",
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
								ReferencedTableName: "users",
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
					{
						TableName: "orders",
						Comment:   "orders comment",
						Columns: []pipe.Column{
							{
								ColumnName: "total_amount",
								ColumnType: "string",
							},
							{
								ColumnName: "order_id",
								ColumnType: "string",
							},
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "deleted_at",
								ColumnType: "gorm.DeletedAt",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
						},
					},
				},
				dirs: []string{"test_data", "another_test_data"},
			},
			want: []pipe.Table{
				{
					TableName: "users",
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
					TableName: "user_details",
					Comment:   "user detail comment",
					Columns: []pipe.Column{
						{
							ColumnName: "user_detail_id",
							ColumnType: "string",
						},
						{
							ColumnName:          "user_id",
							ColumnType:          "string",
							ReferencedTableName: "users",
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
				{
					TableName: "orders",
					Comment:   "orders comment",
					Columns: []pipe.Column{
						{
							ColumnName: "order_id",
							ColumnType: "string",
						},
						{
							ColumnName: "total_amount",
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
							ColumnType: "gorm.DeletedAt",
						},
					},
				},
			},
		},
		{
			name: "success （ファイル内に複数のモデルがある場合）",
			args: args{
				tables: []pipe.Table{
					{
						TableName: "regions",
						Comment:   "regions comment",
						Columns: []pipe.Column{
							{
								ColumnName: "country",
								ColumnType: "string",
							},
							{
								ColumnName: "region_id",
								ColumnType: "int",
							},
							{
								ColumnName: "region_name",
								ColumnType: "string",
							},
						},
					},
					{
						TableName: "shippers",
						Comment:   "shippers comment",
						Columns: []pipe.Column{
							{
								ColumnName: "phone",
								ColumnType: "string",
							},
							{
								ColumnName: "shipper_id",
								ColumnType: "int",
							},
							{
								ColumnName: "shipper_name",
								ColumnType: "string",
							},
						},
					},
					{
						TableName: "orders",
						Comment:   "orders comment",
						Columns: []pipe.Column{
							{
								ColumnName: "total_amount",
								ColumnType: "string",
							},
							{
								ColumnName: "order_id",
								ColumnType: "string",
							},
							{
								ColumnName: "created_at",
								ColumnType: "time.Time",
							},
							{
								ColumnName: "deleted_at",
								ColumnType: "gorm.DeletedAt",
							},
							{
								ColumnName: "updated_at",
								ColumnType: "time.Time",
							},
						},
					},
				},
				dirs: []string{"another_test_data"},
			},
			want: []pipe.Table{
				{
					TableName: "regions",
					Comment:   "regions comment",
					Columns: []pipe.Column{
						{
							ColumnName: "region_id",
							ColumnType: "int",
						},
						{
							ColumnName: "region_name",
							ColumnType: "string",
						},
						{
							ColumnName: "country",
							ColumnType: "string",
						},
					},
				},
				{
					TableName: "shippers",
					Comment:   "shippers comment",
					Columns: []pipe.Column{
						{
							ColumnName: "shipper_id",
							ColumnType: "int",
						},
						{
							ColumnName: "shipper_name",
							ColumnType: "string",
						},
						{
							ColumnName: "phone",
							ColumnType: "string",
						},
					},
				},
				{
					TableName: "orders",
					Comment:   "orders comment",
					Columns: []pipe.Column{
						{
							ColumnName: "order_id",
							ColumnType: "string",
						},
						{
							ColumnName: "total_amount",
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
							ColumnType: "gorm.DeletedAt",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := transformer.SortColumnByGormModelFile(schema.NamingStrategy{}, tt.args.tables, tt.args.dirs)
			if err != nil {
				t.Errorf("SortColumnByGorm error = %v", err)
			}
			if !cmp.Equal(actual, tt.want) {
				t.Errorf("diff =%v", cmp.Diff(actual, tt.want))
			}
		})
	}
}
