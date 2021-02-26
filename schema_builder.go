package nero

import (
	"github.com/sf9v/mira"
)

// SchemaBuilder is schema builder
type SchemaBuilder struct {
	*Schema
}

// NewSchemaBuilder takes a model struct value and  returns a SchemaBuilder
func NewSchemaBuilder(v interface{}) *SchemaBuilder {
	return &SchemaBuilder{&Schema{
		typeInfo:   mira.NewType(v),
		columns:    []*Column{},
		templaters: []Templater{},
	}}
}

// Build builds the schema
func (sb *SchemaBuilder) Build() *Schema {
	templates := sb.templaters

	// use default template set
	if len(templates) == 0 {
		templates = []Templater{NewPostgresTemplate()}
	}

	// get pkg imports
	importMap := map[string]int{}
	for _, c := range append(sb.columns, sb.identity) {
		if c.typeInfo.PkgPath() != "" {
			importMap[c.typeInfo.PkgPath()] = 1
		}
	}

	imports := []string{sb.typeInfo.PkgPath()}
	for imp := range importMap {
		imports = append(imports, imp)
	}

	return &Schema{
		typeInfo:   sb.typeInfo,
		pkgName:    sb.pkgName,
		collection: sb.collection,
		identity:   sb.identity,
		columns:    sb.columns,
		imports:    imports,
		templaters: templates,
	}
}

// Schema is a nero schema

// PkgName sets the pkg name
func (sb *SchemaBuilder) PkgName(pkgName string) *SchemaBuilder {
	sb.pkgName = pkgName
	return sb
}

// Collection sets the collection
func (sb *SchemaBuilder) Collection(collection string) *SchemaBuilder {
	sb.collection = collection
	return sb
}

// Identity sets the identity column
func (sb *SchemaBuilder) Identity(column *Column) *SchemaBuilder {
	sb.identity = column
	return sb
}

// Columns sets the columns
func (sb *SchemaBuilder) Columns(columns ...*Column) *SchemaBuilder {
	sb.columns = append(sb.columns, columns...)
	return sb
}

// Templates sets the templates
func (sb *SchemaBuilder) Templates(templaters ...Templater) *SchemaBuilder {
	sb.templaters = append(sb.templaters, templaters...)
	return sb
}
