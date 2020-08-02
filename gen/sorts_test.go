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

func IDAsc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "id",
			Direction: sort.Asc,
		})
	}
}

func IDDesc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "id",
			Direction: sort.Desc,
		})
	}
}

func NameAsc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "name",
			Direction: sort.Asc,
		})
	}
}

func NameDesc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "name",
			Direction: sort.Desc,
		})
	}
}

func UpdatedAtAsc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "updated_at",
			Direction: sort.Asc,
		})
	}
}

func UpdatedAtDesc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "updated_at",
			Direction: sort.Desc,
		})
	}
}

func CreatedAtAsc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "created_at",
			Direction: sort.Asc,
		})
	}
}

func CreatedAtDesc() SortFunc {
	return func(srt *sort.Sorts) {
		srt.Add(&sort.Sort{
			Field:     "created_at",
			Direction: sort.Desc,
		})
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
	assert.Equal(t, expect, got)
}
