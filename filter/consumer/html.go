package consumer

import (
	"html/template"
	"io"

	"github.com/digeon-inc/royle/pipe"
)

func ExportToHTML(output io.Writer, tables []pipe.TableMetaData) error {

	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
	<title>Table Specification</title>
</head>
<body>
<h1>Table Specification</h1>
	{{range .}}
		<h2 id="{{.TableName}}">{{.TableName}}</h2>
		<table border="1">
			<tr>
				<th>Name</th>
				<th>Type</th>
				<th>Nullable</th>
				<th>Constraint</th>
				<th>Referenced</th>
				<th>Default</th>
				<th>Extra</th>
			</tr>
			{{range .Columns}}
				<tr>
					<td>{{.ColumnName}}</td>
					<td>{{.ColumnType}}</td>
					<td>{{.IsNullable}}</td>
					<td>{{.ConstraintTypes}}</td>
					<td><a href="#{{.ReferencedTableName}}">{{.ReferencedTableName}}</a></td>
					<td>{{.ColumnDefault}}</td>
					<td>{{.Extra}}</td>
				</tr>
			{{end}}
		</table>
	{{end}}
</body>
</html>
`

	tmpl, err := template.New("tableTemplate").Parse(htmlTemplate)

	if err != nil {
		return err
	}

	return tmpl.Execute(output, tables)
}
