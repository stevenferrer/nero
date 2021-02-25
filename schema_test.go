package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero"
)

func TestSchemaBuilder(t *testing.T) {
	pkg := "mypkg"
	collection := "mycollection"
	schema := nero.NewSchemaBuilder().
		PkgName(pkg).
		Collection(collection).
		Columns(
			nero.NewColumnBuilder("id", int64(0)).
				Auto().Identity().StructField("ID").Build(),
			nero.NewColumnBuilder("name", "").Build(),
		).
		Build()

	assert.Equal(t, pkg, schema.PkgName)
	assert.Equal(t, collection, schema.Collection)
	assert.Len(t, schema.Columns, 2)
}
