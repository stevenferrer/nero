package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newSortsBlock(t *testing.T) {
	block := newSortsBlock()
	expect := strings.TrimSpace(`
sorts := &sort.Sorts{}
for _, sf := range sfs {
	sf(sorts)
}
for _, s := range sorts.All() {
	col := fmt.Sprintf("%q", s.Col)
	switch s.Direction {
	case sort.Asc:
		qb = qb.OrderBy(col + " ASC")
	case sort.Desc:
		qb = qb.OrderBy(col + " DESC")
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
