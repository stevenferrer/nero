package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero"
)

func TestPostgresTemplate(t *testing.T) {
	tmpl := nero.NewPostgresTemplate().WithFilename("pg.go")
	assert.Equal(t, "pg.go", tmpl.Filename())

	_, err := nero.ParseTemplate(tmpl)
	require.NoError(t, err)
}
