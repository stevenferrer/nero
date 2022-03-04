package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/predicate"
)

func newPredicateFile(schema nero.Schema) (*File, error) {
	tmpl, err := template.New("predicates.tmpl").
		Funcs(nero.NewFuncMap()).Parse(predicatesTmpl)
	if err != nil {
		return nil, err
	}

	data := struct {
		EqOps   []predicate.Operator
		LtGtOps []predicate.Operator
		NullOps []predicate.Operator
		InOps   []predicate.Operator
		Schema  nero.Schema
	}{
		EqOps: []predicate.Operator{
			predicate.Eq,
			predicate.NotEq,
		},
		LtGtOps: []predicate.Operator{
			predicate.Gt,
			predicate.GtOrEq,
			predicate.Lt,
			predicate.LtOrEq,
		},
		NullOps: []predicate.Operator{
			predicate.IsNull,
			predicate.IsNotNull,
		},
		InOps: []predicate.Operator{
			predicate.In,
			predicate.NotIn,
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
	"github.com/stevenferrer/nero/predicate"
	{{range $import := .Schema.Imports -}}
		"{{$import}}"
	{{end -}}
)

{{ $fields := prependToFields .Schema.Identity .Schema.Fields }}

{{range $field := $fields -}}
	{{if $field.IsComparable  -}}
        {{ range $op := $.EqOps }} 
            // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
            func {{$field.StructField}}{{$op.String}} ({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) predicate.Func {
                return func(predicates []predicate.Predicate) []predicate.Predicate {
                    return append(predicates, predicate.Predicate{
                        Field: "{{$field.Name}}",
                        Operator: predicate.{{$op.String}},
                        {{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
                            Argument: pq.Array({{$field.Identifier}}),
                        {{else -}}
                            Argument: {{$field.Identifier}},
                        {{end -}}
                    })
                }
            }
        {{end}}

        {{ range $op := $.LtGtOps }}
            {{if or $field.TypeInfo.IsNumeric (isType $field.TypeInfo.V "time.Time")}}
                // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
                func {{$field.StructField}}{{$op.String}} ({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) predicate.Func {
                    return func(predicates []predicate.Predicate) []predicate.Predicate {
                        return append(predicates, predicate.Predicate{
                            Field: "{{$field.Name}}",
                            Operator: predicate.{{$op.String}},
                            {{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
                                Argument: pq.Array({{$field.Identifier}}),
                            {{else -}}
                                Argument: {{$field.Identifier}},
                            {{end -}}
                        })
                    }
                }
            {{end}}
        {{end }}

        {{ range $op := $.NullOps }}
            {{if $field.IsNillable}}
                // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
                func {{$field.StructField}}{{$op.String}} () predicate.Func {
                    return func(predicates []predicate.Predicate) []predicate.Predicate {
                        return append(predicates, predicate.Predicate{
                            Field: "{{$field.Name}}",
                            Operator: predicate.{{$op.String}},
                        })
                    }
                }
            {{end}}
        {{end}}

        {{ range $op := $.InOps }}
            // {{$field.StructField}}{{$op.String}} {{$op.Desc}} operator on {{$field.StructField}} field
            func {{$field.StructField}}{{$op.String}} ({{$field.IdentifierPlural}} ...{{rawType $field.TypeInfo.V}}) predicate.Func {
                args := []interface{}{}
                for _, v := range {{$field.IdentifierPlural}} {
                    args = append(args, v)
                }

                return func(predicates []predicate.Predicate) []predicate.Predicate {
                    return append(predicates, predicate.Predicate{
                        Field: "{{$field.Name}}",
                        Operator: predicate.{{$op.String}},
                        Argument: args,
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
    func FieldX{{$op.String}}FieldY (fieldX, fieldY Field) predicate.Func {
        return func(predicates []predicate.Predicate) []predicate.Predicate {
            return append(predicates, predicate.Predicate{
                Field: fieldX.String(),
                Operator: predicate.{{$op.String}},
                Argument: fieldY,
            })
        }
    }
{{end}}

{{ range $op := $.LtGtOps }} 
    // FieldX{{$op.String}}FieldY fieldX {{$op.Desc}} fieldY
    // 
    // fieldX and fieldY must be of the same type
    func FieldX{{$op.String}}FieldY (fieldX, fieldY Field) predicate.Func {
        return func(predicates []predicate.Predicate) []predicate.Predicate {
            return append(predicates, predicate.Predicate{
                Field: fieldX.String(),
                Operator: predicate.{{$op.String}},
                Argument: fieldY,
            })
        }
    }
{{end}}
`
