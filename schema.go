package nero

// Schemaer is a contract for a nero schema
type Schemaer interface {
	Schema() *Schema
}

// Schema is a nero schema
type Schema struct {
	// Pkg is the package name of the generated files
	Pkg string
	// Collection is the name of the collection/table
	Collection string
	// Columns is the list of columns
	Columns Columns
}

// Columns is an array of columns
type Columns []*column

type column struct {
	Name      string
	T         interface{}
	FieldName string
	IsAuto    bool
	IsIdent   bool
}

// NewColumn creates a new column
func NewColumn(name string, t interface{}) *column {
	return &column{Name: name, T: t}
}

// Auto is an auto-filled column e.g.
// auto-increment id, auto-filled date etc.
func (c *column) Auto() *column {
	c.IsAuto = true
	return c
}

func (c *column) Ident() *column {
	c.IsIdent = true
	return c
}

// Field is the struct field name. Use when nero generated the
// wrong field for your struct e.g. your struct field is "ID"
// but nero generated "Id" instead.
func (c *column) Field(field string) *column {
	c.FieldName = field
	return c
}
