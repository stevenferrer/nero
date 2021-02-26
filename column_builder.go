package nero

import "github.com/sf9v/mira"

// ColumnBuilder is a column
type ColumnBuilder struct {
	*Column
}

// NewColumnBuilder returns a ColumnBuilder
func NewColumnBuilder(name string, v interface{}) *ColumnBuilder {
	return &ColumnBuilder{&Column{
		name:     name,
		typeInfo: mira.NewType(v),
	}}
}

// Build builds the column
func (cb *ColumnBuilder) Build() *Column {
	return &Column{
		name:        cb.name,
		typeInfo:    cb.typeInfo,
		auto:        cb.auto,
		comparable:  cb.comparable,
		optional:    cb.optional,
		structField: cb.structField,
	}
}

// Auto sets the column as auto-filled
func (cb *ColumnBuilder) Auto() *ColumnBuilder {
	cb.auto = true
	return cb
}

// Comparable sets the column as comparable
func (cb *ColumnBuilder) Comparable() *ColumnBuilder {
	cb.comparable = true
	return cb
}

// Optional sets the column as optional
func (cb *ColumnBuilder) Optional() *ColumnBuilder {
	cb.optional = true
	return cb
}

// StructField sets the struct field name override. This is useful when the
// inferred struct field is different from the actual field e.g. The struct
// field of the model is "ID" but being referred to as "Id" in the generated code
func (cb *ColumnBuilder) StructField(structField string) *ColumnBuilder {
	cb.structField = structField
	return cb
}
