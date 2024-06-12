# MySQL documentation

## orders

Stores basic information about orders

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| id | int | NO | PRIMARY KEY |  |  | auto_increment |  |
| product_name | varchar(255) | NO |  |  |  |  |  |
| user_id | int | YES | FOREIGN KEY | [users](#users) |  |  |  |
| quantity | int | YES |  |  | 1 |  | Quantity of the product being ordered, defaults to 1 |

## user_details

Stores basic information about users details

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| user_detail_id | varchar(30) | NO | PRIMARY KEY |  |  |  |  |
| user_id | int | YES | FOREIGN KEY | [users](#users) |  |  |  |
| name | varchar(255) | NO |  |  |  |  |  |
| created_at | datetime(3) | YES |  |  |  |  |  |
| updated_at | datetime(3) | YES |  |  |  |  |  |

## users

Stores basic information about users

| Name | Type | Nullable | Constraints | Referenced | Default | Extra | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| id | int | NO | PRIMARY KEY |  |  | auto_increment |  |
| name | varchar(255) | NO |  |  |  |  |  |
| email | varchar(255) | NO | UNIQUE |  |  |  |  |
