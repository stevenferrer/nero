package nero

import (
	"github.com/stevenferrer/mira"
)

// SchemaBuilder is used for building a schema
type SchemaBuilder struct {
	s Schema
}

// NewSchemaBuilder takes a struct value and  returns a SchemaBuilder
func NewSchemaBuilder(v interface{}) *SchemaBuilder {
	return &SchemaBuilder{s: Schema{
		typeInfo:  mira.NewTypeInfo(v),
		fields:    make([]Field, 0, 2),
		templates: make([]Template, 0, 2),
	}}
}

// PkgName sets the package name
func (sb *SchemaBuilder) PkgName(pkgName string) *SchemaBuilder {
	sb.s.pkgName = pkgName
	return sb
}

// Table sets the database table/collection name
func (sb *SchemaBuilder) Table(table string) *SchemaBuilder {
	sb.s.table = table
	return sb
}

// Identity sets the identity field
func (sb *SchemaBuilder) Identity(field Field) *SchemaBuilder {
	sb.s.identity = field
	return sb
}

// Fields sets the fields
func (sb *SchemaBuilder) Fields(fields ...Field) *SchemaBuilder {
	sb.s.fields = append(sb.s.fields, fields...)
	return sb
}

// Templates sets the templates
func (sb *SchemaBuilder) Templates(templates ...Template) *SchemaBuilder {
	sb.s.templates = append(sb.s.templates, templates...)
	return sb
}

// Build builds the schema
func (sb *SchemaBuilder) Build() Schema {
	templates := sb.s.templates

	// use default template set
	if len(templates) == 0 {
		templates = []Template{
			NewPostgresTemplate(),
			NewSQLiteTemplate(),
		}
	}

	// get pkg imports
	importMap := map[string]int{}
	for _, fld := range append(sb.s.fields, sb.s.identity) {
		if fld.typeInfo.PkgPath() != "" {
			importMap[fld.typeInfo.PkgPath()] = 1
		}
	}

	imports := []string{sb.s.typeInfo.PkgPath()}
	for imp := range importMap {
		imports = append(imports, imp)
	}

	return Schema{
		typeInfo:  sb.s.typeInfo,
		pkgName:   sb.s.pkgName,
		table:     sb.s.table,
		identity:  sb.s.identity,
		fields:    sb.s.fields,
		imports:   imports,
		templates: templates,
	}
}
