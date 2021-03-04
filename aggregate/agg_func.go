package aggregate

// AggFunc is an aggregate decorator
type AggFunc func([]*Aggregate) []*Aggregate
