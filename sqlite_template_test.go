package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero"
)

func TestSQLiteTemplate(t *testing.T) {
	tmpl := nero.NewSQLiteTemplate().WithFilename("sqlite.go")
	assert.Equal(t, "sqlite.go", tmpl.Filename())

	_, err := nero.ParseTemplater(tmpl)
	require.NoError(t, err)
}
