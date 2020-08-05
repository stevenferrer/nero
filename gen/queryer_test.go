package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newQueryer(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	queryer := newQueryer(schema)
	expect := strings.TrimSpace(`
type Queryer struct {
	collection string
	columns    []string
	limit      uint64
	offset     uint64
	pfs        []PredFunc
	sfs        []SortFunc
}

func NewQueryer() *Queryer {
	return &Queryer{
		collection: collection,
		columns:    []string{"id", "name", "group_res", "updated_at", "created_at"},
	}
}

func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

func (q *Queryer) Limit(limit uint64) *Queryer {
	q.limit = limit
	return q
}

func (q *Queryer) Offset(offset uint64) *Queryer {
	q.offset = offset
	return q
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", queryer))
	assert.Equal(t, expect, got)
}
