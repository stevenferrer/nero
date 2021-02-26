package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero"
)

func TestColumnBuilder(t *testing.T) {
	column := nero.NewColumnBuilder("id", int64(0)).
		Auto().StructField("ID").
		Optional().Comparable().
		Build()

	assert.True(t, column.IsOptional())
	assert.True(t, column.IsAuto())
	assert.True(t, column.IsComparable())

	assert.NotNil(t, column.TypeInfo())
	assert.Equal(t, "id", column.Name())
	assert.Equal(t, "ID", column.FieldName())
	assert.Equal(t, "id", column.Identifier())
	assert.Equal(t, "ids", column.IdentifierPlural())
	assert.Equal(t, true, column.CanHavePreds())
	assert.Equal(t, false, column.IsArray())
	assert.Equal(t, false, column.IsNillable())
	assert.Equal(t, false, column.IsValueScanner())
}
