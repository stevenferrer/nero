package aggregate

// Aggregate is an aggregate
type Aggregate struct {
	Col string
	Fn  Function
}

// Aggregates is an aggregate builder
type Aggregates struct {
	list []*Aggregate
}

// Add adds aggregates to the list
func (aggs *Aggregates) Add(ags ...*Aggregate) {
	aggs.list = append(aggs.list, ags...)
}

// All returns all predicates
func (aggs *Aggregates) All() []*Aggregate {
	return aggs.list
}
