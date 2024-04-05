package consumer

import (
	"io"
	"text/template"

	"github.com/digeon-inc/royle/pipe"
)

func ExportToMarkdown(output io.Writer, title string, tables []pipe.TableMetaData) error {
	markdownTemplate := `
# {{.Title}}
{{range .Tables}}
## {{.TableName}}
| Name | Type | Nullable | Constraints | Referenced | Default | Extra |
|-------------|----------------|-------------|-------------|-------|------------------------|-------------------|
{{range .Columns}}| {{.ColumnName}} | {{.ColumnType}} | {{.IsNullable}} | {{.ConstraintTypes}} | {{if ne .ReferencedTableName ""}}[{{.ReferencedTableName}}](#{{.ReferencedTableName}}){{end}} | {{.ColumnDefault}} | {{.Extra}} |
{{end}}
{{end}}
`

	tmpl, err := template.New("tableTemplate").Parse(markdownTemplate)
	if err != nil {
		return err
	}

	data := struct {
		Title  string
		Tables []pipe.TableMetaData
	}{
		Title:  title,
		Tables: tables,
	}

	return tmpl.Execute(output, data)

}
