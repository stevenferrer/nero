package aggregate

// Aggregate is an aggregate parameter
type Aggregate struct {
	Field string
	Op    Operator
}

// AggFunc is an aggregate list decorator
type AggFunc func([]Aggregate) []Aggregate
