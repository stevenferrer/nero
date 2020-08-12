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

func (op Operator) Description() string {
	switch op {
	case Eq:
		return "Equal"
	case NotEq:
		return "Not equal"
	case Gt:
		return "Greater than"
	case GtOrEq:
		return "Greater than or equal"
	case Lt:
		return "Less than"
	case LtOrEq:
		return "Less than or equal"
	case IsNull:
		return "Is null"
	case IsNotNull:
		return "Is not null"
	}

	return "Invalid"
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
)
