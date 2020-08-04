package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newSorts(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	predicates := newSorts(schema)
	expect := strings.TrimSpace(`
type SortFunc func(*sort.Sorts)

func Asc(col Column) SortFunc {
	return func(srts *sort.Sorts) {
		srts.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Asc,
		})
	}
}

func Desc(col Column) SortFunc {
	return func(srts *sort.Sorts) {
		srts.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Desc,
		})
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
	assert.Equal(t, expect, got)
}
