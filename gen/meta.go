package gen

import (
	"bytes"
	"text/template"

	gen "github.com/sf9v/nero/gen/internal"
)

func newMetaFile(schema *gen.Schema) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	tmpl, err := template.New("meta.tmpl").Parse(metaTmpl)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(buf, schema)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

const metaTmpl = `
package {{.Pkg}}

// Collection is the name of the collection
const Collection = "{{ .Collection }}"

// Column is a {{.Type.Name}} column
type Column int

// String implements Stringer
func (c Column) String() string {
	switch c {
    {{range .Cols -}}
        case Column{{.Field -}}:
            return "{{.Name -}}"
    {{end -}}
	}

	return "invalid"
}

const (
	{{- range $i, $e := .Cols -}}
        {{if (eq $i 0)}}
            Column{{$e.Field }} Column = iota
        {{else -}}
            Column{{$e.Field}}
        {{end -}}
    {{end -}}
)`
