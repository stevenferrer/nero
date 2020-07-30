package internal

import "github.com/iancoleman/strcase"

type Col struct {
	Name string
	// FieldName overrides the struct field name
	FieldName string
	Typ       *Typ
	Auto      bool
	Ident     bool
}

func (c *Col) CamelName() string {
	return strcase.ToCamel(c.Name)
}

func (c *Col) LowerCamelName() string {
	return strcase.ToLowerCamel(c.Name)
}
