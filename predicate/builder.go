package predicate

// TODO:
// - Database should also be the one to dictate the supported operators
// - The operations listed here are all the possible operators

// Builder is a predicat builder
type Builder struct {
	predicates []*Predicate
}

// Append adds predicate to the builder
func (b *Builder) Append(ps ...*Predicate) {
	b.predicates = append(b.predicates, ps...)
}

// Predicates returns the list of predicates
func (b *Builder) Predicates() []*Predicate {
	return b.predicates
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
	// Eq is equal
	Eq Op = "Eq"
	// NotEq is not equal
	NotEq Op = "NotEq"
	// Gt is greater than
	Gt Op = "Gt"
	// GtOrEq is greater than or equal
	GtOrEq Op = "GtOrEq"
	// Lt is less than
	Lt Op = "Lt"
	// LtOrEq is less than or equal
	LtOrEq Op = "LtOrEq"
)

// Ops is the list of predicate operators
var Ops = []Op{Eq, NotEq, Gt, GtOrEq, Lt, LtOrEq}
