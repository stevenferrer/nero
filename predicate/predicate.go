package predicate

// Predicate is a predicate
type Predicate struct {
	Field    string
	Operator Operator
	Argument interface{}
}

// Func is a predicate func
type Func func([]Predicate) []Predicate
