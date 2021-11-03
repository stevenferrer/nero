package gen_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero/gen"
	"github.com/stevenferrer/nero/gen/internal"
)

func TestGenerate(t *testing.T) {
	u := internal.User{}
	files, err := gen.Generate(u.Schema())
	assert.NoError(t, err)
	assert.Len(t, files, 6)

	for _, file := range files {
		require.NotEmpty(t, file.Filename())
		require.NotEmpty(t, file.Bytes())
	}

	// create base directory
	basePath := "userrepo"
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	assert.NoError(t, err)
	for _, f := range files {
		err = f.Render(basePath)
		require.NoError(t, err)
	}
}
