package aggregate

// Function is an aggregate function
type Function int

func (f Function) String() string {
	switch f {
	case Avg:
		return "Avg"
	case Count:
		return "Count"
	case Max:
		return "Max"
	case Min:
		return "Min"
	case Sum:
		return "Sum"
	case None:
		return "None"
	}

	return ""
}

// Desc is a aggregate function description
func (f Function) Desc() string {
	switch f {
	case Avg:
		return "average"
	case Count:
		return "count"
	case Max:
		return "max"
	case Min:
		return "min"
	case Sum:
		return "sum"
	case None:
		return "none"
	}

	return ""
}

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
