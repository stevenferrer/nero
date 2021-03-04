package aggregate

// Operator is an aggregate operator
type Operator int

const (
	// Avg is average operator
	Avg Operator = iota
	// Count is the count operator
	Count
	// Max is the max operator
	Max
	// Min is the min operator
	Min
	// Sum is the sum operator
	Sum
	// None is not an aggregate function and is only used
	// when you want to include a column in the result
	None
)

func (o Operator) String() string {
	return [...]string{
		"Avg",
		"Count",
		"Max",
		"Min",
		"Sum",
		"None",
	}[o]
}

// Desc is a aggregate function description
func (o Operator) Desc() string {
	return [...]string{
		"average",
		"count",
		"max",
		"min",
		"sum",
		"none",
	}[o]
}
