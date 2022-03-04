package aggregate

// Aggregate is an aggregate
type Aggregate struct {
	Field    string
	Operator Operator
}

// Func is an aggregate func
type Func func([]Aggregate) []Aggregate
