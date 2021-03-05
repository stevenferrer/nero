package comparison

// Predicate is a predicate parameter
type Predicate struct {
	Col string
	Op  Operator
	Arg interface{}
}
