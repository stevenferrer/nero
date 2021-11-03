package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevenferrer/nero"
)

func TestFieldBuilder(t *testing.T) {
	field := nero.NewFieldBuilder("id", int64(0)).Auto().
		StructField("ID").Optional().Build()

	assert.True(t, field.IsOptional())
	assert.True(t, field.IsAuto())

	assert.NotNil(t, field.TypeInfo())
	assert.Equal(t, "id", field.Name())
	assert.Equal(t, "ID", field.StructField())
	assert.Equal(t, "id", field.Identifier())
	assert.Equal(t, "ids", field.IdentifierPlural())
	assert.Equal(t, true, field.IsComparable())
	assert.Equal(t, false, field.IsArray())
	assert.Equal(t, false, field.IsNillable())
	assert.Equal(t, false, field.IsValueScanner())
}
