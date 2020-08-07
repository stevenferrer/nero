package gen

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/example"
)

type example1 struct{}

func (*example1) Schema() *nero.Schema {
	return &nero.Schema{}
}

func TestGenerate(t *testing.T) {
	files, err := Generate(new(example.User))
	assert.NoError(t, err)
	assert.Len(t, files, 7)

	for _, file := range files {
		require.NotEmpty(t, file.Name())
		require.NotEmpty(t, file.Bytes())
	}

	// create base directory
	basePath := path.Join("gen", "user")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	// render files
	err = files.Render(basePath)
	assert.NoError(t, err)

	_, err = Generate(new(example1))
	assert.Error(t, err)
}
