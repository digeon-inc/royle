# royleのプログラムについて

## 概要 

本プログラムはパイプラインアーキテクチャを採用している。主な流れは下のようになる。[統合テスト](https://github.com/digeon-inc/royle/blob/main/integration_test/integration_test.go)で実際の流れを見ればより具体的に理解できる。

```mermaid
graph LR;
    A[producer.FetchColumnMetadata] -->C[transformer.MergeMetadataIntoTables];
    B[producer.FetchTableMetadata] -->C;
    C -->D[consumer.ExportToMarkdown];
```

各フィルターの役割は以下のようになる。

### producer（データの生成器）

#### FetchColumnMetadata

FetchColumnMetadataでは、**カラム**についての情報をmysqlデータベースのスキーマであるinformation_schemaからsqlを使って取得している。

### FetchTableMetadata

FetchTableMetadataでは、**テーブル**についての情報をmysqlデータベースのinformation_schemaからsqlを使って取得している。

### transformer（データの変換器）

#### MergeMetadataIntoTables

MergeMetadataIntoTablesでは、FetchColumnMetadataとFetchTableMetadataから取得したテーブルとカラムについての情報をカラム名をkeyとして結合している。単一のsqlでテーブル情報とカラム情報を結合しなかった理由は、メンテナンス性が下がるからだ。単一のsqlでやってしまうとどうしても文が長くなり、単一責任の原則にも反してまい可読性が悪化する。

### consumer（データの表示器）

#### ExportToMarkdown

MergeMetadataIntoTablesで変換したデータをmarkdown形式で出力する。出力先は標準出力である。
