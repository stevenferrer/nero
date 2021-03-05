package aggregate

// AggFunc is an aggregate list decorator
type AggFunc func([]*Aggregate) []*Aggregate
