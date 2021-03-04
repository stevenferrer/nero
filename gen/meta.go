package gen

import (
	"bytes"
	"text/template"

	"github.com/sf9v/nero"
)

func newMetaFile(schema *nero.Schema) (*bytes.Buffer, error) {
	tmpl, err := template.New("meta.tmpl").Parse(metaTmpl)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, schema)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// TODO: wrap all template data into a struct

const metaTmpl = `
// Code generated by nero, DO NOT EDIT.
package {{.PkgName}}

// Collection is the name of the collection
const Collection = "{{ .Collection }}"

// Column is a {{.TypeInfo.Name}} column
type Column int

// String returns the string representation of the Column
func (c Column) String() string {
	return [...]string{
	"{{.Identity.Name}}",
    {{range .Columns -}}
		"{{.Name}}",
    {{end -}}
	}[c]
}

const (
	Column{{.Identity.FieldName}} Column = iota
	{{range $e := .Columns -}}
		Column{{$e.FieldName}}
    {{end -}}
)`
