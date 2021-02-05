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
func (a *Aggregates) Add(ags ...*Aggregate) {
	a.list = append(a.list, ags...)
}

// All returns all predicates
func (a *Aggregates) All() []*Aggregate {
	return a.list
}
