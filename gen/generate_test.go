package gen

import (
	"os"
	"path"
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
		Columns: []*nero.Column{
			nero.NewColumn("id1", int64(0)).Ident(),
			nero.NewColumn("id2", int64(0)).Ident(),
		},
	}
}

func TestGenerateExample(t *testing.T) {
	files, err := Generate(new(gen.Example))
	assert.NoError(t, err)
	assert.Len(t, files, 11)

	for _, file := range files {
		require.NotEmpty(t, file.Name())
		require.NotEmpty(t, file.Bytes())
	}

	// create base directory
	basePath := path.Join("gen", "example")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	// render files
	err = files.Render(basePath)
	assert.NoError(t, err)

	_, err = Generate(new(example1))
	assert.Error(t, err)
}
