package example

import (
	"os"
	"path"
	"testing"

	"github.com/sf9v/nero/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	outFiles, err := gen.Generate(new(User))
	require.NoError(t, err)

	// write files
	basePath := path.Join("gen/user")
	err = os.MkdirAll(basePath, os.ModePerm)
	require.NoError(t, err)

	for _, outFile := range outFiles {
		filePath := path.Join(basePath, outFile.Name)
		f, err := os.Create(filePath)
		assert.NoError(t, err)

		_, err = f.Write(outFile.Bytes())
		assert.NoError(t, err)
	}
}
