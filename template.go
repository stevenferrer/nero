package nero

import (
	"fmt"
	"reflect"
	"text/template"

	"github.com/stevenferrer/mira"
)

// Template is an interface that wraps the Filename and Content method
type Template interface {
	// Filename is the filename of the generated file
	Filename() string
	// Content is returns the template content
	Content() string
}

// ParseTemplate parses the repository template
func ParseTemplate(t Template) (*template.Template, error) {
	return template.New(t.Filename() + ".tmpl").
		Funcs(NewFuncMap()).Parse(t.Content())
}

// NewFuncMap returns a template func map
func NewFuncMap() template.FuncMap {
	return template.FuncMap{
		"type":            typeFunc,
		"rawType":         rawTypeFunc,
		"zeroValue":       zeroValueFunc,
		"prependToFields": prependToFields,
		"fileHeaders":     fileHeadersFunc,
		"isType":          isTypeFunc,
	}
}

// typeFunc returns the type of the value
func typeFunc(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		return fmt.Sprintf("%T", v)
	}

	ev := reflect.New(resolveType(t)).Elem().Interface()
	return fmt.Sprintf("%T", ev)
}

// rawTypeFunc returns the raw type of the value
func rawTypeFunc(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// resolveType resolves the type of the value
func resolveType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr:
		return resolveType(t.Elem())
	}
	return t
}

// zeroValueFunc returns zero value as a string
func zeroValueFunc(v interface{}) string {
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

// prependToFields prepends a field to the list of fields
func prependToFields(field Field, fields []Field) []Field {
	return append([]Field{field}, fields...)
}

const fileHeaders = `
// Code generated by nero, DO NOT EDIT.
`

// fileHeadersFunc returns the standard file headers
func fileHeadersFunc() string {
	return fileHeaders
}

func isTypeFunc(v interface{}, typeStr string) bool {
	return typeFunc(v) == typeStr
}
