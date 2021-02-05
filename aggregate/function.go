package aggregate

// Function is an aggregate function
type Function int

const (
	// Avg is average aggregate function
	Avg Function = iota
	// Count is the count aggregate function
	Count
	// Max is the max aggregate function
	Max
	// Min is the min aggregate function
	Min
	// Sum is the sum aggregate function
	Sum
	// None is not an aggregate function and is only used
	// when you want to include a column in the result
	None
)

func (f Function) String() string {
	return [...]string{
		"Avg",
		"Count",
		"Max",
		"Min",
		"Sum",
		"None",
	}[f]
}

// Desc is a aggregate function description
func (f Function) Desc() string {
	return [...]string{
		"average",
		"count",
		"max",
		"min",
		"sum",
		"none",
	}[f]
}
