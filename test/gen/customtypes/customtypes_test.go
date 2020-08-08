package customtypes_test

import (
	"os"
	"path"
	"testing"

	"github.com/sf9v/nero/gen"
	"github.com/sf9v/nero/test/gen/customtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomTypes(t *testing.T) {
	files, err := gen.Generate(new(customtypes.Custom))
	require.NoError(t, err)
	assert.Len(t, files, 6)

	// create base directory
	basePath := path.Join("gen", "user")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	err = files.Render(basePath)
	assert.NoError(t, err)
}
