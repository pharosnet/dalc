package generates

import (
	"bytes"
	"fmt"
	"github.com/pharosnet/dalc/cmd/dalc/internal/entry"
	"github.com/pharosnet/dalc/cmd/dalc/internal/parser/commons"
	"strings"
	"text/template"
)

func generateQuery(packageName string, jsonTag bool, queries []*entry.Query) (fs []*GenerateFile, err error) {

	dataList := makeGenerateQueryData(packageName, jsonTag, queries)
	tmplSelect, templateSelectErr := template.New("QUERY_SELECT_TEMPLATE").Parse(templateSelect)
	if templateSelectErr != nil {
		err = templateSelectErr
		return
	}
	tmplExec, templateExecErr := template.New("QUERY_EXEC_TEMPLATE").Parse(templateExec)
	if templateExecErr != nil {
		err = templateExecErr
		return
	}

	fs = make([]*GenerateFile, 0, 1)

	for _, data := range dataList {
		buf := bytes.NewBufferString("")
		if data.Exec {
			execErr := tmplExec.Execute(buf, data)
			if execErr != nil {
				err = execErr
				return
			}
		} else {
			execErr := tmplSelect.Execute(buf, data)
			if execErr != nil {
				err = execErr
				return
			}
		}
		file := &GenerateFile{
			Name:    fmt.Sprintf("query.%s.go", data.RawName),
			Content: buf.Bytes(),
		}
		fs = append(fs, file)
	}

	return
}

func makeGenerateQueryData(packageName string, jsonTag bool, queries []*entry.Query) (dataList []*GenerateQueryData) {
	dataList = make([]*GenerateQueryData, 0, 1)
	for _, query := range queries {
		data := &GenerateQueryData{}
		data.Package = packageName
		if query.Kind != entry.SelectQueryKind {
			data.Exec = true
		} else {
			data.Exec = false
		}
		data.RawName = strings.ToLower(query.RawName)
		data.Imports = make(map[string]string)
		data.LowName = strings.ToLower(query.Name[0:1]) + query.Name[1:]
		data.Name = query.Name

		data.QuerySQL = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(query.Sql, "\r", " "), "\n", " "), ";", "")

		data.RequestFields = make([]*QueryField, 0, 1)
		for _, expr := range query.CondExprList.ExprList {
			data.RequestFields = append(data.RequestFields, &QueryField{
				Name: expr.Name,
				Type: expr.GoType.Name,
				Tags: "",
			})
			if data.Exec {
				if expr.GoType.Package != "" && expr.GoType.Package != "github.com/pharosnet/dalc" {
					data.Imports[expr.GoType.Package] = expr.GoType.Package
				}
			} else {
				if expr.GoType.Package != "" && expr.GoType.Package != "sql" && expr.GoType.Package != "database/sql" && expr.GoType.Package != "github.com/pharosnet/dalc" && expr.GoType.Package != "context" {
					data.Imports[expr.GoType.Package] = expr.GoType.Package
				}
			}

		}

		if len(query.TableList) == 1 {
			table := query.TableList[0]
			mapped := 0
			for _, expr := range query.SelectExprList.ExprList {
				for _, column := range table.Ref.Columns {
					if strings.ToUpper(expr.ColumnName) == strings.ToUpper(column.Name) {
						mapped++
					}
				}
			}
			if mapped == len(table.Ref.Columns) {
				data.IsTable = true
				data.Table = table.Ref
			}
		}

		if !data.IsTable {
			data.ResultFields = make([]*QueryField, 0, 1)
			for _, expr := range query.SelectExprList.ExprList {
				data.ResultFields = append(data.ResultFields, &QueryField{
					Name: expr.Name,
					Type: expr.GoType.Name,
					Tags: func(jsonTag bool) string {
						if !jsonTag {
							return ""
						}
						tag := commons.CamelToSnakeLow(expr.Name)
						return fmt.Sprintf("`json:\"%s\"`", tag)
					}(jsonTag),
				})

				if data.Exec {
					if expr.GoType.Package != "" && expr.GoType.Package != "github.com/pharosnet/dalc" {
						data.Imports[expr.GoType.Package] = expr.GoType.Package
					}
				} else {
					if expr.GoType.Package != "" && expr.GoType.Package != "sql" && expr.GoType.Package != "database/sql" && expr.GoType.Package != "github.com/pharosnet/dalc" && expr.GoType.Package != "context" {
						data.Imports[expr.GoType.Package] = expr.GoType.Package
					}
				}
			}
		}

		dataList = append(dataList, data)
	}
	return
}

var templateSelect = `
package  {{ .Package }}

import (
    "context"
    "database/sql"
    "github.com/pharosnet/dalc"
    {{ range $key, $value := .Imports }}
        "{{ $key }}"
    {{ end }}
)

// ************* {{ .RawName }} *************
const {{ .LowName }}SQL = "{{ .QuerySQL }}"

type {{ .Name }}Request struct { {{ range $key, $field := .RequestFields}}
    {{ $field.Name }} {{ $field.Type }}{{ end }}
}

{{ if eq .IsTable true }}
    type {{ .Name }}ResultIterator func(ctx context.Context, result *{{ .Table.GoName }}) (err error)
{{ else }}
    type {{ .Name }}Result struct { {{ range $key, $field := .ResultFields}}
        {{ $field.Name }} {{ $field.Type }} {{ $field.Tags }}{{ end }}
    }

    type {{ .Name }}ResultIterator func(ctx context.Context, result *{{ .Name }}Result) (err error)
{{ end }}


func {{ .Name }}(ctx dalc.PreparedContext, request *{{ .Name }}Request, iterator {{ .Name }}ResultIterator) (err error) {

    args := dalc.NewArgs() {{ range $key, $field := .RequestFields}}
    args.Arg(request.{{ $field.Name }}){{ end }}


    err = dalc.Query(ctx, {{ .LowName }}SQL, args, func(ctx context.Context, rows *sql.Rows, rowErr error) (err error) {

        if rowErr != nil {
            err = rowErr
            return
        }

        {{ if eq .IsTable true }}
            result := &{{ .Table.GoName }}{}
            scanErr := result.scanSQLRow(rows)
        {{ else }}
            result := &{{ .Name }}Result{}
            scanErr := rows.Scan( {{ range $key, $field := .ResultFields}}
                &result.{{ $field.Name }},{{ end }}
            )
        {{ end }}

        if scanErr != nil {
            err = scanErr
            return
        }

        err = iterator(ctx, result)

        return
    })

    return
}
`

var templateExec = `
package  {{ .Package }}

import (
    "github.com/pharosnet/dalc"
    {{ range $key, $value := .Imports }}
        "{{ $key }}"
    {{ end }}
)

// ************* {{ .RawName }} *************
const {{ .LowName }}SQL = "{{ .QuerySQL }}"

type {{ .Name }}Request struct { {{ range $key, $field := .RequestFields}}
    {{ $field.Name }} {{ $field.Type }}{{ end }}
}

func {{ .Name }}(ctx dalc.PreparedContext, request *{{ .Name }}Request) (affected int64, err error) {

    args := dalc.NewArgs() {{ range $key, $field := .RequestFields}}
    args.Arg(request.{{ $field.Name }}){{ end }}

    affected, err = dalc.Execute(ctx, {{ .LowName }}SQL, args)

    return
}
`