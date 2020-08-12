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
// Query is the query builder for User
type Queryer struct {
	limit  uint64
	offset uint64
	pfs    []PredFunc
	sfs    []SortFunc
}

// NewQueryer returns a query builder
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where adds predicates to the query
func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

// Sort adds sort/order to the query
func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

// Limit adds limit to the query
func (q *Queryer) Limit(limit uint64) *Queryer {
	q.limit = limit
	return q
}

// Offset adds offset to the query
func (q *Queryer) Offset(offset uint64) *Queryer {
	q.offset = offset
	return q
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", queryer))
	assert.Equal(t, expect, got)
}
