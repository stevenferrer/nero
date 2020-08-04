package aggregate

type Aggregate struct {
	Col string
	Fn  Function
}

type Aggregates struct {
	list []*Aggregate
}

func (aggs *Aggregates) Add(ags ...*Aggregate) {
	aggs.list = append(aggs.list, ags...)
}

func (aggs *Aggregates) All() []*Aggregate {
	return aggs.list
}
