package internal

import (
	"testing"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/example"
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
	schema, err := BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	assert.Equal(t, "User", schema.Type.Name())
	assert.Equal(t, "users", schema.Collection)
	assert.Equal(t, "user", schema.Pkg)
	assert.Len(t, schema.Cols, 10)

	ident := schema.Ident
	assert.Equal(t, "id", ident.Name)
	assert.Equal(t, "ID", ident.StructField)

	// no ident defined
	_, err = BuildSchema(new(example1))
	assert.Error(t, err)

	// multiple idents defined
	_, err = BuildSchema(new(example2))
	assert.Error(t, err)
}
