package comparison

// Operator is comparison operator type
type Operator int

// List of comparison operators
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
