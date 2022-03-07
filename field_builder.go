package nero

import (
	"github.com/stevenferrer/mira"
)

// FieldBuilder is a field builder
type FieldBuilder struct {
	f Field
}

// NewFieldBuilder takes a field name and a value and returns a FieldBuilder
func NewFieldBuilder(name string, v interface{}) *FieldBuilder {
	return &FieldBuilder{Field{
		name:     name,
		typeInfo: mira.NewTypeInfo(v),
	}}
}

// Auto sets the auto-populated flag
func (fb *FieldBuilder) Auto() *FieldBuilder {
	fb.f.auto = true
	return fb
}

// Optional sets the optional flag
func (fb *FieldBuilder) Optional() *FieldBuilder {
	fb.f.optional = true
	return fb
}

// StructField sets the struct field
func (fb *FieldBuilder) StructField(structField string) *FieldBuilder {
	fb.f.structField = structField
	return fb
}

// Build builds the field
func (fb *FieldBuilder) Build() Field {
	return Field{
		name:        fb.f.name,
		typeInfo:    fb.f.typeInfo,
		auto:        fb.f.auto,
		optional:    fb.f.optional,
		structField: fb.f.structField,
	}
}
