package internal

import (
	"testing"

	"github.com/sf9v/nero"
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
		Columns: []*nero.Column{
			nero.NewColumn("id1", int64(0)).Ident(),
			nero.NewColumn("id2", int64(0)).Ident(),
		},
	}
}

func TestBuildSchema(t *testing.T) {
	schema, err := BuildSchema(new(Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	assert.Equal(t, "Example", schema.Typ.Name)
	assert.Equal(t, "examples", schema.Coln)
	assert.Equal(t, "example", schema.Pkg)
	assert.Len(t, schema.Cols, 4)

	ident := schema.Ident
	assert.Equal(t, "id", ident.Name)
	assert.Equal(t, "ID", ident.Field)

	// no ident defined
	_, err = BuildSchema(new(example1))
	assert.Error(t, err)

	// multiple idents defined
	_, err = BuildSchema(new(example2))
	assert.Error(t, err)
}
