package nero

import (
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/stevenferrer/mira"
	stringsx "github.com/stevenferrer/nero/x/strings"
)

// Field is a field
type Field struct {
	// name is the field name
	name string
	// typeInfo is the field type info
	typeInfo *mira.TypeInfo
	// StructField overrides the struct field
	structField string
	// Auto is the auto-filled flag
	auto,
	// Optional is the optional flag
	optional bool
}

// TypeInfo returns the type info
func (f Field) TypeInfo() *mira.TypeInfo {
	return f.typeInfo
}

// Name returns the field name
func (f Field) Name() string {
	return f.name
}

// StructField returns the struct field
func (f Field) StructField() string {
	structField := stringsx.ToCamel(f.name)
	if len(f.structField) > 0 {
		structField = f.structField
	}

	return structField
}

// Identifier returns the lower-camelized struct field
func (f Field) Identifier() string {
	return stringsx.ToLowerCamel(f.StructField())
}

// IdentifierPlural returns the plural form of identifier
func (f Field) IdentifierPlural() string {
	return inflection.Plural(f.Identifier())
}

// IsArray returns true if field is an array or a slice
func (f Field) IsArray() bool {
	kind := f.typeInfo.T().Kind()
	return kind == reflect.Array ||
		kind == reflect.Slice
}

// IsNillable returns true if the field is nillable
func (f Field) IsNillable() bool {
	return f.typeInfo.IsNillable()
}

// IsValueScanner returns true if field implements value scanner
func (f Field) IsValueScanner() bool {
	t := reflect.TypeOf(f.typeInfo.V())
	if t.Kind() != reflect.Ptr {
		t = reflect.New(t).Type()
	}

	return t.Implements(reflect.TypeOf(new(ValueScanner)).Elem())
}

// IsAuto returns the auto flag
func (f Field) IsAuto() bool {
	return f.auto
}

// IsOptional returns the optional flag
func (f Field) IsOptional() bool {
	return f.optional
}

// IsComparable returns true if field is comparable i.e. with comparisong operators
func (f Field) IsComparable() bool {
	kind := f.typeInfo.T().Kind()
	return !(kind == reflect.Map ||
		kind == reflect.Slice)
}
