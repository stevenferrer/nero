package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newSorts(t *testing.T) {
	predicates := newSorts()
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
