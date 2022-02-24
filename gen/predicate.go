package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/comparison"
)

func newPredicateFile(schema nero.Schema) (*File, error) {
	tmpl, err := template.New("predicates.tmpl").
		Funcs(nero.NewFuncMap()).Parse(predicatesTmpl)
	if err != nil {
		return nil, err
	}

	data := struct {
		EqOps   []comparison.Operator
		LtGtOps []comparison.Operator
		NullOps []comparison.Operator
		InOps   []comparison.Operator
		Schema  nero.Schema
	}{
		EqOps: []comparison.Operator{
			comparison.Eq,
			comparison.NotEq,
		},
		LtGtOps: []comparison.Operator{
			comparison.Gt,
			comparison.GtOrEq,
			comparison.Lt,
			comparison.LtOrEq,
		},
		NullOps: []comparison.Operator{
			comparison.IsNull,
			comparison.IsNotNull,
		},
		InOps: []comparison.Operator{
			comparison.In,
			comparison.NotIn,
		},
		Schema: schema,
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return &File{name: "predicate.go", buf: buf.Bytes()}, nil
}

const predicatesTmpl = `
{{- fileHeaders -}}

package {{.Schema.PkgName}}

import (
	"github.com/lib/pq"
	"github.com/stevenferrer/nero/comparison"
	{{range $import := .Schema.Imports -}}
		"{{$import}}"
	{{end -}}
)

{{ $fields := prependToFields .Schema.Identity .Schema.Fields }}

{{range $field := $fields -}}
	{{if $field.IsComparable  -}}
        {{ range $op := $.EqOps }} 
            // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
            func {{$field.StructField}}{{$op.String}} ({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) comparison.PredFunc {
                return func(preds []comparison.Predicate) []comparison.Predicate {
                    return append(preds, comparison.Predicate{
                        Field: "{{$field.Name}}",
                        Op: comparison.{{$op.String}},
                        {{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
                            Arg: pq.Array({{$field.Identifier}}),
                        {{else -}}
                            Arg: {{$field.Identifier}},
                        {{end -}}
                    })
                }
            }
        {{end}}

        {{ range $op := $.LtGtOps }}
            {{if or $field.TypeInfo.IsNumeric (isType $field.TypeInfo.V "time.Time")}}
                // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
                func {{$field.StructField}}{{$op.String}} ({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) comparison.PredFunc {
                    return func(preds []comparison.Predicate) []comparison.Predicate {
                        return append(preds, comparison.Predicate{
                            Field: "{{$field.Name}}",
                            Op: comparison.{{$op.String}},
                            {{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
                                Arg: pq.Array({{$field.Identifier}}),
                            {{else -}}
                                Arg: {{$field.Identifier}},
                            {{end -}}
                        })
                    }
                }
            {{end}}
        {{end }}

        {{ range $op := $.NullOps }}
            {{if $field.IsNillable}}
                // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
                func {{$field.StructField}}{{$op.String}} () comparison.PredFunc {
                    return func(preds []comparison.Predicate) []comparison.Predicate {
                        return append(preds, comparison.Predicate{
                            Field: "{{$field.Name}}",
                            Op: comparison.{{$op.String}},
                        })
                    }
                }
            {{end}}
        {{end}}

        {{ range $op := $.InOps }}
            // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
            func {{$field.StructField}}{{$op.String}} ({{$field.IdentifierPlural}} ...{{rawType $field.TypeInfo.V}}) comparison.PredFunc {
                args := []interface{}{}
                for _, v := range {{$field.IdentifierPlural}} {
                    args = append(args, v)
                }

                return func(preds []comparison.Predicate) []comparison.Predicate {
                    return append(preds, comparison.Predicate{
                        Field: "{{$field.Name}}",
                        Op: comparison.{{$op.String}},
                        Arg: args,
                    })
                }
            }
        {{end}}
	{{end}}
{{end -}}

{{ range $op := $.EqOps }} 
    // FieldX{{$op.String}}FieldY fieldX {{$op.Desc}} fieldY
    //
    // fieldX and fieldY must be of the same type
    func FieldX{{$op.String}}FieldY (fieldX, fieldY Field) comparison.PredFunc {
        return func(preds []comparison.Predicate) []comparison.Predicate {
            return append(preds, comparison.Predicate{
                Field: fieldX.String(),
                Op: comparison.{{$op.String}},
                Arg: fieldY,
            })
        }
    }
{{end}}

{{ range $op := $.LtGtOps }} 
    // FieldX{{$op.String}}FieldY fieldX {{$op.Desc}} fieldY
    // 
    // fieldX and fieldY must be of the same type
    func FieldX{{$op.String}}FieldY (fieldX, fieldY Field) comparison.PredFunc {
        return func(preds []comparison.Predicate) []comparison.Predicate {
            return append(preds, comparison.Predicate{
                Field: fieldX.String(),
                Op: comparison.{{$op.String}},
                Arg: fieldY,
            })
        }
    }
{{end}}
`
