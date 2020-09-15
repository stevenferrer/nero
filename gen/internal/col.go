package internal

import (
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
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

func (c *Col) Field() string {
	field := c.CamelName()
	if len(c.StructField) > 0 {
		field = c.StructField
	}
	return field
}

func (c *Col) Identifier() string {
	return stringsx.ToLowerCamel(c.Field())
}

func (c *Col) IdentifierPlural() string {
	return inflection.Plural(c.Identifier())
}

func (c *Col) HasPreds() bool {
	kind := c.Type.T().Kind()
	return !(kind == reflect.Map ||
		kind == reflect.Slice)
}
