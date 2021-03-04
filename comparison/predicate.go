package comparison

// Predicate is a predicate
type Predicate struct {
	Col string
	Op  Operator
	Arg interface{}
}
