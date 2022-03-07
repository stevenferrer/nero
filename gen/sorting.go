package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/sorting"
)

func newSortFile(schema nero.Schema) (*File, error) {
	tmpl, err := template.New("sort.tmpl").
		Funcs(nero.NewFuncMap()).Parse(sortTmpl)
	if err != nil {
		return nil, err
	}

	data := struct {
		Directions []sorting.Direction
		Schema     nero.Schema
	}{
		Directions: []sorting.Direction{
			sorting.Asc, sorting.Desc,
		},
		Schema: schema,
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return &File{name: "sorting.go", buf: buf.Bytes()}, nil
}

const sortTmpl = `
{{- fileHeaders -}}

package {{.Schema.PkgName}}

import (
	"github.com/stevenferrer/nero/sorting"
)

{{range $direction := .Directions}}
// {{$direction.String}} {{$direction.Desc}} sort direction
func {{$direction.String}}(fields ...Field) sorting.Func {
	return func(sortings sorting.Sortings) sorting.Sortings {
		for _, field := range fields {
			sortings = append(sortings, sorting.Sorting{
				Field: field.String(),
				Direction: sorting.{{$direction.String}},
			})
		}

		return sortings
	}
}
{{end}}
`
