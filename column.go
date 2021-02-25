package nero

// Column is a column
type Column struct {
	// Name is the column name
	Name string
	// T is the column type
	T interface{}
	// StructField overrides the struct field
	StructField string
	// Identity is an identity column
	Identity,
	// Auto is an auto-filled column and thus will not have any setter
	Auto,
	// Optional is an optional column i.e. not required
	Optional,
	// ColumnComparable is a column that can be compared
	// with other columns in the same collection/table
	ColumnComparable bool
}

// ColumnBuilder is a column
type ColumnBuilder struct {
	column *Column
}

// NewColumnBuilder returns a ColumnBuilder
func NewColumnBuilder(name string, t interface{}) *ColumnBuilder {
	return &ColumnBuilder{
		column: &Column{
			Name: name,
			T:    t,
		},
	}
}

// Build builds the column
func (c *ColumnBuilder) Build() *Column {
	return c.column
}

// Auto is an auto-filled column i.e. auto-increment
// primary key id, auto-filled date etc.
func (c *ColumnBuilder) Auto() *ColumnBuilder {
	c.column.Auto = true
	return c
}

// Identity is an identity column
func (c *ColumnBuilder) Identity() *ColumnBuilder {
	c.column.Identity = true
	return c
}

// ColumnComparable is a column that can be compared
// with other columns in the same collection/table
func (c *ColumnBuilder) ColumnComparable() *ColumnBuilder {
	c.column.ColumnComparable = true
	return c
}

// StructField overrides the struct field name. Use this when the
// inferred struct field is wrong. e.g. The struct field of the model
// is "ID" but being referred to as "Id" in the generated code
func (c *ColumnBuilder) StructField(structField string) *ColumnBuilder {
	c.column.StructField = structField
	return c
}

// Optional is an optional column i.e. not required
func (c *ColumnBuilder) Optional() *ColumnBuilder {
	c.column.Optional = true
	return c
}
