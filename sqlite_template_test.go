package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero"
)

func TestSQLiteTemplate(t *testing.T) {
	tmpl := nero.NewSQLiteTemplate().WithFilename("sqlite.go")
	assert.Equal(t, "sqlite.go", tmpl.Filename())

	_, err := nero.ParseTemplate(tmpl)
	require.NoError(t, err)
}
