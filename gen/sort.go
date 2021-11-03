package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/sort"
)

func newSortFile(schema *nero.Schema) (*File, error) {
	tmpl, err := template.New("sort.tmpl").
		Funcs(nero.NewFuncMap()).Parse(sortTmpl)
	if err != nil {
		return nil, err
	}

	data := struct {
		Directions []sort.Direction
		Schema     *nero.Schema
	}{
		Directions: []sort.Direction{
			sort.Asc, sort.Desc,
		},
		Schema: schema,
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return &File{name: "sort.go", buf: buf.Bytes()}, nil
}

const sortTmpl = `
{{- fileHeaders -}}

package {{.Schema.PkgName}}

import (
	"github.com/stevenferrer/nero/sort"
)

{{range $direction := .Directions}}
// {{$direction.String}} {{$direction.Desc}} sort direction
func {{$direction.String}}(field Field) sort.SortFunc {
	return func(sorts []*sort.Sort) []*sort.Sort {
		return append(sorts, &sort.Sort{
			Field: field.String(),
			Direction: sort.{{$direction.String}},
		})
	}
}
{{end}}
`
