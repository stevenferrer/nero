package predicate

// TODO:
// - Database should also be the one to dictate the supported operators
// - The operations listed here are all the possible operators

type Builder struct {
	predicates []*Predicate
}

func (b *Builder) Append(ps ...*Predicate) {
	b.predicates = append(b.predicates, ps...)
}

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
type Op int

const (
	// Eq is equal
	Eq Op = iota
	// NotEq is not equal
	NotEq
	// Gt is greater than
	Gt
	// GtOrEq is greater than or equal
	GtOrEq
	// Lt is less than
	Lt
	// LtOrEq is less than or equal
	LtOrEq
)
