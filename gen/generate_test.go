package gen_test

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/gen"
)

func TestGenerate(t *testing.T) {
	files, err := gen.Generate(&example.User{})
	assert.NoError(t, err)
	assert.Len(t, files, 6)

	for _, file := range files {
		require.NotEmpty(t, file.Filename())
		require.NotEmpty(t, file.Buf())
	}

	// create base directory
	basePath := path.Join("gen", "user")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	assert.NoError(t, err)
	for _, f := range files {
		err = f.Render(basePath)
		require.NoError(t, err)
	}
}
