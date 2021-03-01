package nero

import (
	"github.com/sf9v/mira"
)

// SchemaBuilder is schema builder
type SchemaBuilder struct {
	schema *Schema
}

// NewSchemaBuilder takes a model struct value and  returns a SchemaBuilder
func NewSchemaBuilder(v interface{}) *SchemaBuilder {
	return &SchemaBuilder{schema: &Schema{
		typeInfo:   mira.NewTypeInfo(v),
		columns:    []*Column{},
		templaters: []Templater{},
	}}
}

// Build builds the schema
func (sb *SchemaBuilder) Build() *Schema {
	templates := sb.schema.templaters

	// use default template set
	if len(templates) == 0 {
		templates = []Templater{NewPostgresTemplate()}
	}

	// get pkg imports
	importMap := map[string]int{}
	for _, c := range append(sb.schema.columns, sb.schema.identity) {
		if c.typeInfo.PkgPath() != "" {
			importMap[c.typeInfo.PkgPath()] = 1
		}
	}

	imports := []string{sb.schema.typeInfo.PkgPath()}
	for imp := range importMap {
		imports = append(imports, imp)
	}

	return &Schema{
		typeInfo:   sb.schema.typeInfo,
		pkgName:    sb.schema.pkgName,
		collection: sb.schema.collection,
		identity:   sb.schema.identity,
		columns:    sb.schema.columns,
		imports:    imports,
		templaters: templates,
	}
}

// Schema is a nero schema

// PkgName sets the pkg name
func (sb *SchemaBuilder) PkgName(pkgName string) *SchemaBuilder {
	sb.schema.pkgName = pkgName
	return sb
}

// Collection sets the collection
func (sb *SchemaBuilder) Collection(collection string) *SchemaBuilder {
	sb.schema.collection = collection
	return sb
}

// Identity sets the identity column
func (sb *SchemaBuilder) Identity(column *Column) *SchemaBuilder {
	sb.schema.identity = column
	return sb
}

// Columns sets the columns
func (sb *SchemaBuilder) Columns(columns ...*Column) *SchemaBuilder {
	sb.schema.columns = append(sb.schema.columns, columns...)
	return sb
}

// Templates sets the templates
func (sb *SchemaBuilder) Templates(templaters ...Templater) *SchemaBuilder {
	sb.schema.templaters = append(sb.schema.templaters, templaters...)
	return sb
}
