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
	case In:
		return "In"
	case NotIn:
		return "NotIn"
	}

	return "Invalid"
}

// Desc is a predicate operator description
func (op Operator) Desc() string {
	switch op {
	case Eq:
		return "equal"
	case NotEq:
		return "not equal"
	case Gt:
		return "greater than"
	case GtOrEq:
		return "greater than or equal"
	case Lt:
		return "less than"
	case LtOrEq:
		return "less than or equal"
	case IsNull:
		return "is null"
	case IsNotNull:
		return "is not null"
	case In:
		return "in"
	case NotIn:
		return "not in"
	}

	return ""
}

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
