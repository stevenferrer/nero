package comparison

// Predicate is a predicate parameter
type Predicate struct {
	Field string
	Op    Operator
	Arg   interface{}
}

// PredFunc is a predicate list decorator
type PredFunc func([]Predicate) []Predicate
