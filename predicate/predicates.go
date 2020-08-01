package predicate

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

// Predicate is a predicate
type Predicate struct {
	Op    Op
	Field string
	Val   interface{}
}

// Op is predicate operator
type Op string

const (
	// Eq is an equal operator
	Eq Op = "Eq"
	// NotEq is a not equal operator
	NotEq Op = "NotEq"
	// Gt is a greater than operator
	Gt Op = "Gt"
	// GtOrEq is a greater than or equal operator
	GtOrEq Op = "GtOrEq"
	// Lt is a less than operator
	Lt Op = "Lt"
	// LtOrEq is a less than or equal operator
	LtOrEq Op = "LtOrEq"
)

// Ops is the list of predicate operators
var Ops = []Op{Eq, NotEq, Gt, GtOrEq, Lt, LtOrEq}
