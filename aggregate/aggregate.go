package aggregate

// Aggregate is an aggregate
type Aggregate struct {
	Field    string
	Operator Operator
}

type Aggregates []Aggregate

// Func is an aggregate func
type Func func(Aggregates) Aggregates

func Build(aggFuncs ...Func) Aggregates {
	aggregates := make(Aggregates, 0, len(aggFuncs))
	for _, aggFunc := range aggFuncs {
		aggregates = aggFunc(aggregates)
	}

	return aggregates
}

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
	// None is used to include a field in the result
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
