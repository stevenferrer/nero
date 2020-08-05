package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func Test_newAggregator(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	meta := newAggregator(schema)
	expect := `
type Aggregator struct {
	collection string
	dest       interface{}
	aggfs      []AggFunc
	pfs        []PredFunc
	sfs        []SortFunc
	groups     []Column
}

func NewAggregator(dest interface{}) *Aggregator {
	return &Aggregator{
		collection: collection,
		dest:       dest,
	}
}

func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
	return a
}
`
	expect = strings.TrimSpace(expect)
	got := strings.TrimSpace(fmt.Sprintf("%#v", meta))
	assert.Equal(t, expect, got)
}
