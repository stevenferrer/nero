package nero

// Schemaer is an interface that wraps the Schema method
type Schemaer interface {
	Schema() *Schema
}

// Schema is a nero schema used for generating the repository
type Schema struct {
	// PkgName is the package name of the generated files
	PkgName string
	// Collection is the name of the collection/table
	Collection string
	// Columns is the list of columns
	Columns []*Column
	// Templates is the list of custom repository templates
	Templates []Templater
}

// SchemaBuilder is schema builder
type SchemaBuilder struct {
	schema *Schema
}

// NewSchemaBuilder returns a SchemaBuilder
func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{schema: &Schema{
		Columns:   []*Column{},
		Templates: []Templater{},
	}}
}

// Build builds the schema
func (s *SchemaBuilder) Build() *Schema {
	return s.schema
}

// Schema is a nero schema

// PkgName sets the pkg name
func (s *SchemaBuilder) PkgName(pkgName string) *SchemaBuilder {
	s.schema.PkgName = pkgName
	return s
}

// Collection sets the collection
func (s *SchemaBuilder) Collection(collection string) *SchemaBuilder {
	s.schema.Collection = collection
	return s
}

// Columns sets the columns
func (s *SchemaBuilder) Columns(columns ...*Column) *SchemaBuilder {
	s.schema.Columns = append(s.schema.Columns, columns...)
	return s
}

// Templates sets the templates
func (s *SchemaBuilder) Templates(templates ...Templater) *SchemaBuilder {
	s.schema.Templates = append(s.schema.Templates, templates...)
	return s
}
