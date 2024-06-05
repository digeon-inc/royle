# MySQL documentation

## orders

Stores basic information about orders

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| id | int | NO | PRIMARY KEY |  |  | auto_increment |  |
| product_name | varchar(255) | NO |  |  |  |  |  |
| quantity | int | YES |  |  | 1 |  | Quantity of the product being ordered, defaults to 1 |
| user_id | int | YES | FOREIGN KEY | [users](#users) |  |  |  |

## users

Stores basic information about users

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| email | varchar(255) | NO | UNIQUE |  |  |  |  |
| id | int | NO | PRIMARY KEY |  |  | auto_increment |  |
| name | varchar(255) | NO |  |  |  |  |  |
