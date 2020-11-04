package nero

// Schemaer is a contract for generating a repository
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
	Columns []*Column
}

// Column is a column
type Column struct {
	cfg *ColumnConfig
}

// ColumnConfig is a column config
type ColumnConfig struct {
	Name        string
	T           interface{}
	StructField string
	Auto,
	Ident,
	Nullable,
	ColumnComparable bool
}

// NewColumn creates a new column
func NewColumn(name string, t interface{}) *Column {
	return &Column{
		cfg: &ColumnConfig{
			Name: name,
			T:    t,
		},
	}
}

// Cfg returns the column config
func (c *Column) Cfg() *ColumnConfig {
	return c.cfg
}

// Auto is an auto-filled column i.e. auto-increment
// primary key id, auto-filled date etc.
func (c *Column) Auto() *Column {
	c.cfg.Auto = true
	return c
}

// Ident is an identity column
func (c *Column) Ident() *Column {
	c.cfg.Ident = true
	return c
}

// Nullable is a nullable column
func (c *Column) Nullable() *Column {
	c.cfg.Nullable = true
	return c
}

// ColumnComparable enables comparison with other column
func (c *Column) ColumnComparable() *Column {
	c.cfg.ColumnComparable = true
	return c
}

// StructField is the struct field name. Use this when the
// inferred struct field is wrong. e.g. The struct field
// is "ID" but being referred to as "Id" in the generated code.
func (c *Column) StructField(structField string) *Column {
	c.cfg.StructField = structField
	return c
}
