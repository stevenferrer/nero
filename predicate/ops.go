package predicate

// Op is predicate operator
type Op string

const (
	// Eq is an equal operator
	Eq Op = "Eq"
	// NotEq is a not equal operator
	NotEq Op = "NotEq"
	// Gt is a greater than operator
	Gt Op = "Gt"
	// GtOrEq is a greater than or equal operator
	GtOrEq Op = "GtOrEq"
	// Lt is a less than operator
	Lt Op = "Lt"
	// LtOrEq is a less than or equal operator
	LtOrEq Op = "LtOrEq"
)
