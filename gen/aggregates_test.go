package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newAggregates(t *testing.T) {
	meta := newAggregates()
	expect := `
// AggFunc is the aggregate function type
type AggFunc func(*aggregate.Aggregates)

// Avg is the average aggregate function
func Avg(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Avg,
		})
	}
}

// Count is the count aggregate function
func Count(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Count,
		})
	}
}

// Max is the max aggregate function
func Max(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Max,
		})
	}
}

// Min is the min aggregate function
func Min(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Min,
		})
	}
}

// Sum is the sum aggregate function
func Sum(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Sum,
		})
	}
}

/*
None is not an aggregate function and is only
used when you want to include a column in the result
*/
func None(col Column) AggFunc {
	return func(aggs *aggregate.Aggregates) {
		aggs.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.None,
		})
	}
}
`
	expect = strings.TrimSpace(expect)
	got := strings.TrimSpace(fmt.Sprintf("%#v", meta))
	assert.Equal(t, expect, got)
}
