package nero

import (
	"fmt"
	"reflect"
	"text/template"

	"github.com/sf9v/mira"
	"github.com/sf9v/nero/comparison"
)

// Templater is an interface that wraps the Filename and Template method
type Templater interface {
	// Filename is the filename of the generated file
	Filename() string
	// Template is template for generating the repository implementation
	Template() string
}

// ParseTemplater parses the repository templater
func ParseTemplater(tmpl Templater) (*template.Template, error) {
	tmplt, err := template.New(tmpl.Filename() + ".tmpl").
		Funcs(NewFuncMap()).
		Parse(tmpl.Template())

	return tmplt, err
}

// NewFuncMap returns a template func map
func NewFuncMap() template.FuncMap {
	return template.FuncMap{
		"realType":         realTypeFunc,
		"rawType":          rawTypeFunc,
		"zero":             zeroFunc,
		"isNullOp":         isNullOp,
		"isInOp":           isInOp,
		"prependToColumns": prependToColumns,
	}
}

func isNullOp(op comparison.Operator) bool {
	return op == comparison.IsNull ||
		op == comparison.IsNotNull
}

func isInOp(op comparison.Operator) bool {
	return op == comparison.In ||
		op == comparison.NotIn
}

func realTypeFunc(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		return fmt.Sprintf("%T", v)
	}

	ev := reflect.New(resolveType(t)).Elem().Interface()
	return fmt.Sprintf("%T", ev)
}

func rawTypeFunc(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func resolveType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr:
		return resolveType(t.Elem())
	}
	return t
}

func zeroFunc(v interface{}) string {
	ti := mira.NewTypeInfo(v)

	if ti.IsNillable() {
		return "nil"
	}

	if ti.IsNumeric() {
		return "0"
	}

	switch ti.T().Kind() {
	case reflect.Bool:
		return "false"
	case reflect.Struct,
		reflect.Array:
		return fmt.Sprintf("(%T{})", v)
	}

	return "\"\""

}

// appendToColumns appends a column to a slice of columns
func prependToColumns(column *Column, columns []*Column) []*Column {
	return append([]*Column{column}, columns...)
}
