package customtypes_test

import (
	"os"
	"path"
	"testing"

	"github.com/stevenferrer/nero/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero/test/customtypes"
)

func TestCustomTypes(t *testing.T) {
	c := customtypes.Custom{}
	files, err := gen.Generate(c.Schema())
	require.NoError(t, err)
	assert.Len(t, files, 7, "should have 7 generated files")

	// create base directory
	basePath := path.Join("customrepo")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	for _, f := range files {
		err = f.Render(basePath)
		require.NoError(t, err)
	}
}
