package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero"
)

func TestColumnBuilder(t *testing.T) {
	column := nero.NewColumnBuilder("id", int64(0)).Identity().
		Auto().StructField("ID").Optional().Build()
	assert.Equal(t, "id", column.Name)
	_, ok := column.T.(int64)
	assert.True(t, ok)
	assert.True(t, column.Identity)
	assert.True(t, column.Auto)
	assert.Equal(t, "ID", column.StructField)
	assert.True(t, column.Optional)

	column = nero.NewColumnBuilder("comparable", "").
		ColumnComparable().Build()
	assert.True(t, column.ColumnComparable)
}
