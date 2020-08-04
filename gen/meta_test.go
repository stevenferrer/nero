package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newMeta(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	meta := newMeta(schema)
	expect := `
const (
	collection = "examples"
)

type Column int

func (c Column) String() string {
	switch c {
	case ColumnID:
		return "id"
	case ColumnName:
		return "name"
	case ColumnUpdatedAt:
		return "updated_at"
	case ColumnCreatedAt:
		return "created_at"
	}
	return "invalid"
}

const (
	ColumnID Column = iota
	ColumnName
	ColumnUpdatedAt
	ColumnCreatedAt
)
`
	expect = strings.TrimSpace(expect)
	got := strings.TrimSpace(fmt.Sprintf("%#v", meta))
	assert.Equal(t, expect, got)
}
