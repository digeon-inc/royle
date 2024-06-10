# royle

MYSQLドキュメントを生成するコマンドラインアプリケーション

## インストール

### golang

```bash
go install github.com/digeon-inc/royle@latest
```

### Mac, Windows or Linux

[バイナリのリリースノート](https://github.com/digeon-inc/royle/releases)

## 使い方

### Linux or Mac

```bash
royle --host mysql --password password --port 3306 --user docker --database template > doc.md
```

#### pandocを使う

[pandoc](https://pandoc.org/)と合わせて使えば様々なフォーマットに変換することができます。

```bash
 royle --host mysql --password password --port 3306 --user docker --database template | pandoc -o doc.html
```


### Windows

```powershell
royle --host mysql --password password --port 3306 --user docker --database template | Out-File -FilePath doc.md -Encoding utf8
```

## コマンドのフラグ

### -h, --help
コマンドについての説明

### -t, --title
ドキュメントのタイトル (default "ROYLE")

### -s, --host
MYSQLのホスト (必須)

### -p, --password
MYSQLのパスワード (必須)

### -r, --port
MYSQLのポート (必須)

### -u, --user
mysqlのユーザー (必須)

### -d, --database
MYSQLのデータベース名 (必須)

## ドキュメント生成例

[example.sql](https://github.com/digeon-inc/royle/blob/main/example.sql)で作ったテーブルから生成したマークダウンをpandocでhtmlにした[ドキュメント](https://digeon-inc.github.io/royle/)

## ドキュメントの各項目について

`##`に続く見出しはテーブル名である。また、その見出しのすぐ下にある説明はテーブルのコメントである。

### 生成したドキュメントの表について

#### Name
カラムの名前。

#### TYPE
カラムのデータ型。

#### Nullable
カラムの NULL 値可能性。 この値は、NULL 値をカラムに格納できる場合は YES で、格納できない場合は NO

#### Constraints
制約のタイプ。 値は、UNIQUE, PRIMARY KEY, FOREIGN KEY または (MySQL 8.0.16) CHECK

#### Referenced
FOREIGN KEY が参照しているテーブルの名前。

#### Default
カラムのデフォルト値。

#### Extra
カラムについての追加情報。
- AUTO_INCREMENT 属性
- ON UPDATE CURRENT_TIMESTAMP 属性
- 生成されたカラムの STORED GENERATED または VIRTUAL GENERATED
- 式のデフォルト値を持つカラムの DEFAULT_GENERATED

#### Comment
カラムのコメント。

## モチベーション

### 自動化による効率化

royleを使うことでドキュメント生成を自動化でき、手作業での作業時間を節約できます。これにより、開発者やチームはより多くの時間を実際のコードの開発や修正に費やすことができます。
Github Actionの例は[こちら](https://github.com/digeon-inc/royle/blob/main/.github/workflows/page.yaml)。

### 正確性の確保

royleはドキュメントを生成するときにmysqlのテーブル情報を毎回取得し、正確な情報を提供します。特に開発者がormを使っている場合、データベース上でテーブルがどのように表現されているかをormのドキュメントなしで正確に知ることができます。

## document

[コードについて](https://github.com/digeon-inc/royle/blob/main/doc/DOC.md)