package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresTemplate(t *testing.T) {
	tmpl := NewPostgresTemplate().WithFilename("pg.go")

	assert.Equal(t, "pg.go", tmpl.Filename())

	_, err := ParseTemplate(tmpl.Content())
	require.NoError(t, err)
}
