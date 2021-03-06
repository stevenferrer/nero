package nero

import (
	"github.com/sf9v/mira"
)

// SchemaBuilder is schema builder
type SchemaBuilder struct {
	sc *Schema
}

// NewSchemaBuilder takes a model struct value and  returns a SchemaBuilder
func NewSchemaBuilder(v interface{}) *SchemaBuilder {
	return &SchemaBuilder{sc: &Schema{
		typeInfo:   mira.NewTypeInfo(v),
		fields:     []*Field{},
		templaters: []Templater{},
	}}
}

// Build builds the schema
func (sb *SchemaBuilder) Build() *Schema {
	templates := sb.sc.templaters

	// use default template set
	if len(templates) == 0 {
		templates = []Templater{NewPostgresTemplate()}
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
		templaters: templates,
	}
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
func (sb *SchemaBuilder) Templates(templaters ...Templater) *SchemaBuilder {
	sb.sc.templaters = append(sb.sc.templaters, templaters...)
	return sb
}
