package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/aggregate"
)

func newAggregateFile(schema nero.Schema) (*File, error) {
	tmpl, err := template.New("aggregates.tmpl").
		Funcs(nero.NewFuncMap()).Parse(aggregatesTmpl)
	if err != nil {
		return nil, err
	}

	data := struct {
		Operators []aggregate.Operator
		Schema    nero.Schema
	}{
		Operators: []aggregate.Operator{
			aggregate.Avg, aggregate.Count,
			aggregate.Max, aggregate.Min,
			aggregate.Sum, aggregate.None,
		},
		Schema: schema,
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return &File{name: "aggregate.go", buf: buf.Bytes()}, nil
}

const aggregatesTmpl = `
{{- fileHeaders -}}

package {{.Schema.PkgName}}

import (
	"github.com/stevenferrer/nero/aggregate"
)

{{range $op := .Operators}}
// {{$op.String}} is the {{$op.Desc}} aggregate operator
func {{$op.String}}(fields ...Field) aggregate.Func {
	return func(aggregates []aggregate.Aggregate) []aggregate.Aggregate {
		for _, field := range fields {
			aggregates = append(aggregates, aggregate.Aggregate{
				Field: field.String(),
				Operator: aggregate.{{$op.String}},
			})
		}
		return aggregates
	}
}
{{end}}
`
