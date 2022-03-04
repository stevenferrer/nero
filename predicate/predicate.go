package predicate

// Predicate is a predicate
type Predicate struct {
	Field    string
	Operator Operator
	Argument interface{}
}

type Predicates []Predicate

// Func is a predicate func
type Func func(Predicates) Predicates

func Build(predFuncs ...Func) Predicates {
	preds := make(Predicates, 0, len(predFuncs))
	for _, predFunc := range predFuncs {
		preds = predFunc(preds)
	}

	return preds
}

// Operator is a predicate operator
type Operator int

// List of predicate operators
const (
	// Eq is an equal operator
	Eq Operator = iota
	// NotEq is a not equal operator
	NotEq
	// Gt is a greater than operator
	Gt
	// GtOrEq is a greater than or equal operator
	GtOrEq
	// Lt is a less than operator
	Lt
	// LtOrEq is a less than or equal operator
	LtOrEq
	// IsNull is an "is null" operator
	IsNull
	// IsNotNull is an "is not null" operator
	IsNotNull
	// In is used to check if a value is in the list
	In
	// In is used to check if a value is not in the list
	NotIn
)

func (o Operator) String() string {
	return [...]string{
		"Eq",
		"NotEq",
		"Gt",
		"GtOrEq",
		"Lt",
		"LtOrEq",
		"IsNull",
		"IsNotNull",
		"In",
		"NotIn",
	}[o]
}

// Desc is a predicate operator description
func (o Operator) Desc() string {
	return [...]string{
		"equal",
		"not equal",
		"greater than",
		"greater than or equal",
		"less than",
		"less than or equal",
		"is null",
		"is not null",
		"in",
		"not in",
	}[o]
}
