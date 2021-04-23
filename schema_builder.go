package nero

import (
	"github.com/sf9v/mira"
)

// SchemaBuilder is used for building a schema
type SchemaBuilder struct {
	sc *Schema
}

// NewSchemaBuilder takes a struct value and  returns a SchemaBuilder
func NewSchemaBuilder(v interface{}) *SchemaBuilder {
	return &SchemaBuilder{sc: &Schema{
		typeInfo:  mira.NewTypeInfo(v),
		fields:    []*Field{},
		templates: []Template{},
	}}
}

// PkgName sets the package name
func (sb *SchemaBuilder) PkgName(pkgName string) *SchemaBuilder {
	sb.sc.pkgName = pkgName
	return sb
}

// Collection sets teh collection
func (sb *SchemaBuilder) Collection(collection string) *SchemaBuilder {
	sb.sc.collection = collection
	return sb
}

// Identity sets the identity field
func (sb *SchemaBuilder) Identity(field *Field) *SchemaBuilder {
	sb.sc.identity = field
	return sb
}

// Fields sets the fields
func (sb *SchemaBuilder) Fields(fields ...*Field) *SchemaBuilder {
	sb.sc.fields = append(sb.sc.fields, fields...)
	return sb
}

// Templates sets the templates
func (sb *SchemaBuilder) Templates(templates ...Template) *SchemaBuilder {
	sb.sc.templates = append(sb.sc.templates, templates...)
	return sb
}

// Build builds the schema
func (sb *SchemaBuilder) Build() *Schema {
	templates := sb.sc.templates

	// use default template set
	if len(templates) == 0 {
		templates = []Template{
			NewPostgresTemplate(),
			NewSQLiteTemplate(),
		}
	}

	// get pkg imports
	importMap := map[string]int{}
	for _, fld := range append(sb.sc.fields, sb.sc.identity) {
		if fld.typeInfo.PkgPath() != "" {
			importMap[fld.typeInfo.PkgPath()] = 1
		}
	}

	imports := []string{sb.sc.typeInfo.PkgPath()}
	for imp := range importMap {
		imports = append(imports, imp)
	}

	return &Schema{
		typeInfo:   sb.sc.typeInfo,
		pkgName:    sb.sc.pkgName,
		collection: sb.sc.collection,
		identity:   sb.sc.identity,
		fields:     sb.sc.fields,
		imports:    imports,
		templates:  templates,
	}
}
