package gen

import (
	"bytes"
	"text/template"

	"github.com/sf9v/nero"
)

func newMetaFile(schema *nero.Schema) (*File, error) {
	tmpl, err := template.New("meta.tmpl").
		Funcs(nero.NewFuncMap()).Parse(metaTmpl)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, schema)
	if err != nil {
		return nil, err
	}

	return &File{name: "meta.go", buf: buf.Bytes()}, nil
}

// TODO: wrap all template data into a struct

const metaTmpl = `
{{- fileHeaders -}}

package {{.PkgName}}

// Collection is the name of the database collection
const Collection = "{{ .Collection }}"

// Field is a {{.TypeInfo.Name}} field
type Field int

// String returns the string representation of the field
func (f Field) String() string {
	return [...]string{
	"{{.Identity.Name}}",
    {{range .Fields -}}
		"{{.Name}}",
    {{end -}}
	}[f]
}

const (
	Field{{.Identity.StructField}} Field = iota + 1
	{{range $e := .Fields -}}
		Field{{$e.StructField}}
    {{end -}}
)`
