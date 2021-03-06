// Code generated by nero, DO NOT EDIT.
package playerrepo

import (
	"github.com/sf9v/nero/aggregate"
)

// Avg is the average aggregate operator
func Avg(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.Avg,
		})
	}
}

// Count is the count aggregate operator
func Count(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.Count,
		})
	}
}

// Max is the max aggregate operator
func Max(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.Max,
		})
	}
}

// Min is the min aggregate operator
func Min(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.Min,
		})
	}
}

// Sum is the sum aggregate operator
func Sum(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.Sum,
		})
	}
}

// None is the none aggregate operator
func None(field Field) aggregate.AggFunc {
	return func(aggs []*aggregate.Aggregate) []*aggregate.Aggregate {
		return append(aggs, &aggregate.Aggregate{
			Field: field.String(),
			Op:    aggregate.None,
		})
	}
}