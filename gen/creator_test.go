package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func Test_newCreator(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	creator := newCreator(schema)
	expect := strings.TrimSpace(`
type Creator struct {
	collection string
	columns    []string
	name       string
	group      string
	updatedAt  *time.Time
}

func NewCreator() *Creator {
	return &Creator{
		collection: collection,
		columns:    []string{"name", "group_res", "updated_at"},
	}
}

func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

func (c *Creator) Group(group string) *Creator {
	c.group = group
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
