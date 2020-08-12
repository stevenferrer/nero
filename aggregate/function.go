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

	return "Invalid"
}

func (f Function) Description() string {
	switch f {
	case Avg:
		return "Avg is the average aggregate function"
	case Count:
		return "Count is the count aggregate function"
	case Max:
		return "Max is the max aggregate function"
	case Min:
		return "Min is the min aggregate function"
	case Sum:
		return "Sum is the sum aggregate function"
	case None:
		return `None is not an aggregate function and is only 
used when you want to include a column in the result`
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
