package internal

import (
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
	"github.com/sf9v/nero"
	stringsx "github.com/sf9v/nero/x/strings"
)

// Col is a column
type Col struct {
	// Name is the column name
	Name string
	// StructField overrides the struct field name
	StructField string
	// Type is the type info of column
	Type *mira.Type
	// Ident is an identity column .i.e. id
	Ident,
	// Auto is an auto-filled column
	Auto,
	// Nullable is a nullable column
	Nullable bool
}

// CamelName returns a camelized version of name
func (c *Col) CamelName() string {
	return stringsx.ToCamel(c.Name)
}

// LowerCamelName returns a lower-camelized version of name
func (c *Col) LowerCamelName() string {
	return stringsx.ToLowerCamel(c.Name)
}

// Field returns the field name
func (c *Col) Field() string {
	field := c.CamelName()
	if len(c.StructField) > 0 {
		field = c.StructField
	}
	return field
}

// Identifier returns the identifier
func (c *Col) Identifier() string {
	return stringsx.ToLowerCamel(c.Field())
}

// IdentifierPlural returns the plural form of identifier
func (c *Col) IdentifierPlural() string {
	return inflection.Plural(c.Identifier())
}

// HasPreds returns true if column can have predicate functions
func (c *Col) HasPreds() bool {
	kind := c.Type.T().Kind()
	return !(kind == reflect.Map ||
		kind == reflect.Slice)
}

// IsArray returns true if column is an array or a slice
func (c *Col) IsArray() bool {
	kind := c.Type.T().Kind()
	return kind == reflect.Array ||
		kind == reflect.Slice
}

var valueScannerType = reflect.TypeOf(new(nero.ValueScanner)).Elem()

// IsValueScanner returns true if column implements value scanner
func (c *Col) IsValueScanner() bool {
	t := reflect.TypeOf(c.Type.V())
	if t.Kind() != reflect.Ptr {
		t = reflect.New(t).Type()
	}

	return t.Implements(valueScannerType)
}
