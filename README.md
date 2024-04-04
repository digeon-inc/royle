# royle

MYSQLドキュメントを生成するコマンドラインアプリケーション

## Install

### golang
  
```zsh
go install github.com/digeon-inc/royle@latest
```

### Mac, Windows or Linux

[リリースノート](https://github.com/digeon-inc/royle/releases)

## Usage

```zsh
royle --host mysql --password password --port 3306 --user docker --dbname template 
```

## Flags

### -h, --help
コマンドについての説明

### -o, --filename
出力するファイルの名前 (デフォルト "output")

### -f, --format
出力するファイルのフォーマット (デフォルト "md")

### -s, --host
MYSQLのホスト (必須)

### -p, --password
MYSQLのパスワード (必須)

### -r, --port
MYSQLのポート (必須)

### -u, --user
mysqlのユーザー (必須)

### -n, --dbname
MYSQLのデータベース名 (必須)

## Output example

### Table Specification

#### orders

| Name | Type | Nullable | Constraints | Referenced | Default | Extra |
|-------------|----------------|-------------|-------------|-------|------------------------|-------------------|
| id | int | NO | PRIMARY KEY |  |  | auto_increment |
| product_name | varchar(255) | NO |  |  |  |  |
| quantity | int | YES |  |  | 1 |  |
| user_id | int | YES | FOREIGN KEY | [users](#users) |  |  |

#### users

| Name | Type | Nullable | Constraints | Referenced | Default | Extra |
|-------------|----------------|-------------|-------------|-------|------------------------|-------------------|
| email | varchar(255) | NO | UNIQUE |  |  |  |
| id | int | NO | PRIMARY KEY |  |  | auto_increment |
| name | varchar(255) | NO |  |  |  |  |

## Output columns

### Name
カラムの名前。

### TYPE
カラムのデータ型。

### Nullable
カラムの NULL 値可能性。 この値は、NULL 値をカラムに格納できる場合は YES で、格納できない場合は NO

### Constraints
制約のタイプ。 値は、UNIQUE, PRIMARY KEY, FOREIGN KEY または (MySQL 8.0.16) CHECK

### Referenced
FOREIGN KEY が参照しているテーブルの名前。

### Default
カラムのデフォルト値。

### Extra
カラムについての追加情報。
- AUTO_INCREMENT 属性
- ON UPDATE CURRENT_TIMESTAMP 属性
- 生成されたカラムの STORED GENERATED または VIRTUAL GENERATED
- 式のデフォルト値を持つカラムの DEFAULT_GENERATED
