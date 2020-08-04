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
	}

	return "Invalid"
}

const (
	// Avg is average aggregate function
	Avg Function = iota
	// Count is count aggregate function
	Count
	// Max is max aggregate function
	Max
	// Min is min aggregate function
	Min
	// Sum is sum aggregate function
	Sum
)
