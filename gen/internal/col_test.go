package internal

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sf9v/mira"
	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero/example"
)

func TestColMethods(t *testing.T) {
	col := Col{
		Name:        "id",
		StructField: "ID",
		Type:        mira.NewType(""),
		Ident:       true,
		Auto:        true,
	}
	assert.Equal(t, "Id", col.CamelName())
	assert.Equal(t, "id", col.LowerCamelName())
	assert.Equal(t, "ID", col.Field())
	assert.Equal(t, "id", col.Identifier())
	assert.Equal(t, "ids", col.IdentifierPlural())
	assert.True(t, col.HasPreds())

	col = Col{Type: mira.NewType(example.Map{})}
	assert.True(t, col.IsValueScanner())

	col = Col{Name: "tags", Type: mira.NewType([]string{})}
	assert.True(t, col.IsArray())
	assert.False(t, col.HasPreds())
	assert.False(t, col.IsValueScanner())

	col = Col{Name: "uuid", Type: mira.NewType(uuid.New())}
	assert.True(t, col.IsValueScanner())
}
