package nero

import (
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"
)

// Column is a column
type Column struct {
	// name is the column name
	name string
	// typeInfo is inferred column type info
	typeInfo *mira.Type
	// StructField overrides the struct field
	structField string
	// Auto is an auto-filled column flag
	auto,
	// Optional is an optional column flag
	optional,
	// ColumnComparable is a flag that enables
	// comparison against other columns
	comparable bool
}

// TypeInfo returns the type info
func (c *Column) TypeInfo() *mira.Type {
	return c.typeInfo
}

// Name returns the column name
func (c *Column) Name() string {
	return c.name
}

// FieldName returns the field name
func (c *Column) FieldName() string {
	field := stringsx.ToCamel(c.name)
	if len(c.structField) > 0 {
		field = c.structField
	}

	return field
}

// Identifier returns the lower-camelized field name
func (c *Column) Identifier() string {
	return stringsx.ToLowerCamel(c.FieldName())
}

// IdentifierPlural returns the plural form of lower-camelized field name
func (c *Column) IdentifierPlural() string {
	return inflection.Plural(c.Identifier())
}

// CanHavePreds returns true if column can have predicate functions
func (c *Column) CanHavePreds() bool {
	kind := c.typeInfo.T().Kind()
	return !(kind == reflect.Map ||
		kind == reflect.Slice)
}

// IsArray returns true if column is an array or a slice
func (c *Column) IsArray() bool {
	kind := c.typeInfo.T().Kind()
	return kind == reflect.Array ||
		kind == reflect.Slice
}

// IsNillable returns true if the column is nillable
func (c *Column) IsNillable() bool {
	return c.typeInfo.IsNillable()
}

// IsValueScanner returns true if column implements value scanner
func (c *Column) IsValueScanner() bool {
	t := reflect.TypeOf(c.typeInfo.V())
	if t.Kind() != reflect.Ptr {
		t = reflect.New(t).Type()
	}

	return t.Implements(reflect.TypeOf(new(ValueScanner)).Elem())
}

// IsAuto returns the auto flag
func (c *Column) IsAuto() bool {
	return c.auto
}

// IsOptional returns the optional flag
func (c *Column) IsOptional() bool {
	return c.optional
}

// IsComparable returns the column comparable flag
func (c *Column) IsComparable() bool {
	return c.comparable
}
