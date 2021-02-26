package nero

import (
	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"
)

// Schemaer is an interface that wraps the Schema method
type Schemaer interface {
	Schema() *Schema
}

// Schema is a schema used for generating the repository
type Schema struct {
	// pkgName is the package name of the generated files
	pkgName string
	// Collection is the name of the collection/table
	collection string
	// typeInfo is the type info the schema model
	typeInfo *mira.Type
	// Identity is the identity column
	identity *Column
	// Columns is the list of columns
	columns []*Column
	// Imports are list of package imports
	imports []string
	// Templates is the list of custom repository templaters
	templaters []Templater
}

// PkgName returns the pkg name
func (s *Schema) PkgName() string {
	return s.pkgName
}

// Collection returns the collection
func (s *Schema) Collection() string {
	return s.collection
}

// Identity returns the identity column
func (s *Schema) Identity() *Column {
	return s.identity
}

// Columns returns the columns
func (s *Schema) Columns() []*Column {
	return s.columns[:]
}

// Imports returns the pkg imports
func (s *Schema) Imports() []string {
	return s.imports[:]
}

// Templaters returns the templaters
func (s *Schema) Templaters() []Templater {
	return s.templaters[:]
}

// TypeInfo returns the type info
func (s *Schema) TypeInfo() *mira.Type {
	return s.typeInfo
}

// TypeName returns the type name
func (s *Schema) TypeName() string {
	return s.typeInfo.Name()
}

// TypeIdentifier returns the type identifier
func (s *Schema) TypeIdentifier() string {
	return stringsx.ToLowerCamel(s.TypeName())
}

// TypeIdentifierPlural returns the plural form of type identifier
func (s *Schema) TypeIdentifierPlural() string {
	return inflection.Plural(s.TypeIdentifier())
}
