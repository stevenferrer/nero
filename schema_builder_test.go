package nero_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero"
)

type MyStruct struct {
	ID   *big.Int
	Name string
}

func TestSchemaBuilder(t *testing.T) {
	pkg := "mypkg"
	collection := "mycollection"
	ms := &MyStruct{}
	schemaBuilder := nero.NewSchemaBuilder(ms).
		PkgName(pkg).Collection(collection).
		Identity(
			nero.NewColumnBuilder("id", ms.ID).
				Auto().StructField("ID").Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", ms.Name).Build(),
		)

	schema := schemaBuilder.Build()

	assert.Equal(t, pkg, schema.PkgName())
	assert.Equal(t, collection, schema.Collection())
	assert.NotNil(t, schema.Identity())
	assert.Len(t, schema.Columns(), 1)
	assert.Len(t, schema.Imports(), 2)
	assert.Len(t, schema.Templaters(), 1)
	assert.NotNil(t, schema.TypeInfo())
	assert.Equal(t, "MyStruct", schema.TypeName())
	assert.Equal(t, "MyStructs", schema.TypeNamePlural())
	assert.Equal(t, "myStruct", schema.TypeIdentifier())
	assert.Equal(t, "myStructs", schema.TypeIdentifierPlural())

	tmpl := nero.NewPostgresTemplate()
	schema = schemaBuilder.Templates(tmpl).Build()
	assert.Len(t, schema.Templaters(), 1)
}
