package internal

import (
	"github.com/sf9v/mira"
	"github.com/sf9v/nero/x/strings"
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
	return strings.ToCamel(c.Name)
}

// LowerCamelName returns a lower-camelized version of name
func (c *Col) LowerCamelName() string {
	return strings.ToLowerCamel(c.Name)
}
