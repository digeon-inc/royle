# MySQL documentation

## Table1

TableComment1

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| ColumnName1 | ColumnType1 | IsNullable1 |  |  | ColumnDefault1 | Extra1 | ColumnComment1 |
| ColumnName2 | ColumnType2 | IsNullable2 |  |  | ColumnDefault2 | Extra2 |  |

## Table2

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| ColumnName1 | ColumnType1 | IsNullable1 | FOREIGN KEY | [Table1](#Table1) | ColumnDefault1 | Extra1 |  |
