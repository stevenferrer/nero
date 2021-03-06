package nero_test

import (
	"testing"

	"github.com/sf9v/nero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresTemplate(t *testing.T) {
	tmpl := nero.NewPostgresTemplate().WithFilename("pg.go")
	assert.Equal(t, "pg.go", tmpl.Filename())

	_, err := nero.ParseTemplater(tmpl)
	require.NoError(t, err)
}
