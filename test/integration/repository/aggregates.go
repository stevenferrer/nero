package repository

import (
	"github.com/sf9v/nero/aggregate"
)

type AggFunc func(*aggregate.Aggregates)

func Avg(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Avg,
		})
	}
}

func Count(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Count,
		})
	}
}

func Max(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Max,
		})
	}
}

func Min(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Min,
		})
	}
}

func Sum(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Sum,
		})
	}
}

func None(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.None,
		})
	}
}
