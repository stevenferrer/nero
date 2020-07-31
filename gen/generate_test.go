package gen

import (
	"testing"

	"github.com/sf9v/nero"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type example1 struct{}

func (*example1) Schema() *nero.Schema {
	return &nero.Schema{}
}

type example2 struct{}

func (*example2) Schema() *nero.Schema {
	return &nero.Schema{
		Columns: nero.Columns{
			nero.NewColumn("id1", int64(0)).Ident(),
			nero.NewColumn("id2", int64(0)).Ident(),
		},
	}
}

func Test_buildSchema(t *testing.T) {
	schema, err := buildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	assert.Equal(t, "Example", schema.Typ.Name)
	assert.Equal(t, "examples", schema.Coln)
	assert.Equal(t, "example", schema.Pkg)
	assert.Len(t, schema.Cols, 4)

	ident := schema.Ident
	assert.Equal(t, "id", ident.Name)
	assert.Equal(t, "ID", ident.FieldName)

	// no ident defined
	_, err = buildSchema(new(example1))
	assert.Error(t, err)

	// multiple idents defined
	_, err = buildSchema(new(example2))
	assert.Error(t, err)
}

func TestGenerate(t *testing.T) {
	outFiles, err := Generate(new(gen.Example))
	assert.NoError(t, err)
	assert.Len(t, outFiles, 8)

	for _, outFile := range outFiles {
		require.NotEmpty(t, outFile.Name)
		require.NotNil(t, outFile.Buffer)
	}
}
