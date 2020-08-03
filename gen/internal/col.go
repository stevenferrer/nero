package internal

import "github.com/iancoleman/strcase"

// Col is a column
type Col struct {
	// Name is the column name
	Name string
	// StructField overrides the struct field name
	StructField string
	// Typ is the type info of column
	Typ *Typ
	// Auto is an auto-filled column
	Auto bool
	// Ident is an identity column .i.e. id
	Ident bool
}

// CamelName returns a camelized version of name
func (c *Col) CamelName() string {
	return strcase.ToCamel(c.Name)
}

// LowerCamelName returns a lower-camelized version of name
func (c *Col) LowerCamelName() string {
	return strcase.ToLowerCamel(c.Name)
}
