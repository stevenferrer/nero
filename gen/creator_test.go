package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newCreator(t *testing.T) {
	schema, err := buildSchema(new(internal.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	creator := newCreator(schema)
	expect := strings.TrimSpace(`
type Creator struct {
	collection string
	columns    []string
	name       string
	updatedAt  *time.Time
}

func NewCreator() *Creator {
	return &Creator{
		collection: collection,
		columns:    []string{"name", "updated_at"},
	}
}

func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", creator))
	assert.Equal(t, expect, got)
}
