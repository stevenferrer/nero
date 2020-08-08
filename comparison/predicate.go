package comparison

// Predicate is a predicate
type Predicate struct {
	Col string
	Val interface{}
	Op  Operator
}

// Predicates is a predicate builder
type Predicates struct {
	list []*Predicate
}

// Add adds predicates to the list
func (p *Predicates) Add(ps ...*Predicate) {
	p.list = append(p.list, ps...)
}

// All returns all predicates
func (p *Predicates) All() []*Predicate {
	return p.list
}
