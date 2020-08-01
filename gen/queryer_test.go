package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newQueryer(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
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
}

func NewQueryer() *Queryer {
	return &Queryer{
		collection: collection,
		columns:    []string{"id", "name", "updated_at", "created_at"},
	}
}

func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
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
