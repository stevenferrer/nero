package gen

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func TestGen(t *testing.T) {
	files, err := Generate(new(example.User))
	assert.NoError(t, err)
	assert.Len(t, files, 6)

	for _, file := range files {
		require.NotEmpty(t, file.FileName())
		require.NotEmpty(t, file.Bytes())
	}

	// create base directory
	basePath := path.Join("gen", "user")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	assert.NoError(t, err)
	for _, file := range files {
		err = file.Render(basePath)
		require.NoError(t, err)
	}
}
