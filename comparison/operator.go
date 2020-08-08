package comparison

// Operator is comparison operator
type Operator int

func (op Operator) String() string {
	switch op {
	case Eq:
		return "Eq"
	case NotEq:
		return "NotEq"
	case Gt:
		return "Gt"
	case GtOrEq:
		return "GtOrEq"
	case Lt:
		return "Lt"
	case LtOrEq:
		return "LtOrEq"
	case IsNull:
		return "IsNull"
	case IsNotNull:
		return "IsNotNull"
	}

	return "Invalid"
}

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
)
