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
