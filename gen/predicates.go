package gen

import (
	"bytes"
	"html/template"

	"github.com/sf9v/nero/comparison"
	gen "github.com/sf9v/nero/gen/internal"
)

func newPredicatesFile(schema *gen.Schema) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	v := struct {
		Ops    []comparison.Operator
		Schema *gen.Schema
	}{
		Ops: []comparison.Operator{
			comparison.Eq,
			comparison.NotEq,
			comparison.Gt,
			comparison.GtOrEq,
			comparison.Lt,
			comparison.LtOrEq,
			comparison.IsNull,
			comparison.IsNotNull,
			comparison.In,
			comparison.NotIn,
		},
		Schema: schema,
	}

	tmpl, err := template.New("predicates.tmpl").Funcs(template.FuncMap{
		"isNullOrNotOp": func(op comparison.Operator) bool {
			return op == comparison.IsNull ||
				op == comparison.IsNotNull
		},
		"isInOrNotOp": func(op comparison.Operator) bool {
			return op == comparison.In ||
				op == comparison.NotIn
		},
	}).Parse(predicatesTmpl)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(buf, v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

const predicatesTmpl = `
package {{.Schema.Pkg}}

import (
	"github.com/sf9v/nero/comparison"

	{{range $import := .Schema.SchemaImports -}}
		"{{$import}}"
	{{end -}}
	{{range $import := .Schema.ColumnImports -}}
		"{{$import}}"
	{{end -}}
)

type PredFunc func(*comparison.Predicates) 

{{range $col := .Schema.Cols -}}
	{{if $col.HasPreds -}}
		{{range $op := $.Ops -}}
			{{if isNullOrNotOp $op }}
				func {{$col.Field}}{{$op.String}} () PredFunc {
					return func(pb *comparison.Predicates) {
						pb.Add(&comparison.Predicate{
							Col: "{{$col.Name}}",
							Op: comparison.{{$op.String}},
						})
					}
				}
			{{else if isInOrNotOp $op }}
				func {{$col.Field}}{{$op.String}} ({{$col.IdentifierPlural}} {{printf "...%T" $col.Type.V}}) PredFunc {
					vals := []interface{}{}
					for _, v := range {{$col.IdentifierPlural}} {
						vals = append(vals, v)
					}

					return func(pb *comparison.Predicates) {
						pb.Add(&comparison.Predicate{
							Col: "{{$col.Identifier}}",
							Op: comparison.{{$op.String}},
							Val: vals,
						})
					}
				}
			{{else}}
				func {{$col.Field}}{{$op.String}} ({{$col.Identifier}} {{printf "%T" $col.Type.V}}) PredFunc {
					return func(pb *comparison.Predicates) {
						pb.Add(&comparison.Predicate{
							Col: "{{$col.Name}}",
							Op: comparison.{{$op.String}},
							Val: {{$col.Identifier}},
						})
					}
				}
			{{end}}
		{{end -}}
	{{end}}
{{end -}}
`
