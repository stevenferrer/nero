package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newAggregates(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	meta := newAggregates(schema)
	expect := `
type AggFunc func(*aggregate.Aggregates)

func Avg(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Avg,
		})
	}
}

func Count(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Count,
		})
	}
}

func Max(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Max,
		})
	}
}

func Min(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Min,
		})
	}
}

func Sum(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Sum,
		})
	}
}
`
	expect = strings.TrimSpace(expect)
	got := strings.TrimSpace(fmt.Sprintf("%#v", meta))
	assert.Equal(t, expect, got)
}
