// Code generated by nero, DO NOT EDIT.
package playerrepo

import (
	"context"
	"reflect"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/aggregate"
	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/sort"
	"github.com/stevenferrer/nero/test/integration/player"
)

// Repository is an interface that provides the methods
// for interacting with a Player repository
type Repository interface {
	// BeginTx starts a transaction
	BeginTx(context.Context) (nero.Tx, error)
	// Create creates a Player
	Create(context.Context, *Creator) (id string, err error)
	// CreateInTx creates a Player in a transaction
	CreateInTx(context.Context, nero.Tx, *Creator) (id string, err error)
	// CreateMany batch creates Players
	CreateMany(context.Context, ...*Creator) error
	// CreateManyInTx batch creates Players in a transaction
	CreateManyInTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries Players
	Query(context.Context, *Queryer) ([]*player.Player, error)
	// QueryTx queries Players in a transaction
	QueryInTx(context.Context, nero.Tx, *Queryer) ([]*player.Player, error)
	// QueryOne queries a Player
	QueryOne(context.Context, *Queryer) (*player.Player, error)
	// QueryOneTx queries a Player in a transaction
	QueryOneInTx(context.Context, nero.Tx, *Queryer) (*player.Player, error)
	// Update updates a Player or many Players
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates a Player many Players in a transaction
	UpdateInTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes a Player or many Players
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes a Player or many Players in a transaction
	DeleteInTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate performs an aggregate query
	Aggregate(context.Context, *Aggregator) error
	// Aggregate performs an aggregate query in a transaction
	AggregateInTx(context.Context, nero.Tx, *Aggregator) error
}

// Creator is a create builder
type Creator struct {
	email     string
	name      string
	age       int
	race      player.Race
	updatedAt *time.Time
}

// NewCreator returns a Creator
func NewCreator() *Creator {
	return &Creator{}
}

// Email sets the Email field
func (c *Creator) Email(email string) *Creator {
	c.email = email
	return c
}

// Name sets the Name field
func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

// Age sets the Age field
func (c *Creator) Age(age int) *Creator {
	c.age = age
	return c
}

// Race sets the Race field
func (c *Creator) Race(race player.Race) *Creator {
	c.race = race
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}

// Validate validates the fields
func (c *Creator) Validate() error {
	var err error
	if isZero(c.email) {
		err = multierror.Append(err, nero.NewErrRequiredField("email"))
	}

	if isZero(c.name) {
		err = multierror.Append(err, nero.NewErrRequiredField("name"))
	}

	if isZero(c.age) {
		err = multierror.Append(err, nero.NewErrRequiredField("age"))
	}

	if isZero(c.race) {
		err = multierror.Append(err, nero.NewErrRequiredField("race"))
	}

	return err
}

// Queryer is a query builder
type Queryer struct {
	limit     uint
	offset    uint
	predFuncs []comparison.PredFunc
	sortFuncs []sort.SortFunc
}

// NewQueryer returns a Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where applies predicates
func (q *Queryer) Where(predFuncs ...comparison.PredFunc) *Queryer {
	q.predFuncs = append(q.predFuncs, predFuncs...)
	return q
}

// Sort applies sorting expressions
func (q *Queryer) Sort(sortFuncs ...sort.SortFunc) *Queryer {
	q.sortFuncs = append(q.sortFuncs, sortFuncs...)
	return q
}

// Limit applies limit
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset applies offset
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder
type Updater struct {
	email     string
	name      string
	age       int
	race      player.Race
	updatedAt *time.Time
	predFuncs []comparison.PredFunc
}

// NewUpdater returns an Updater
func NewUpdater() *Updater {
	return &Updater{}
}

// Email sets the Email field
func (c *Updater) Email(email string) *Updater {
	c.email = email
	return c
}

// Name sets the Name field
func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

// Age sets the Age field
func (c *Updater) Age(age int) *Updater {
	c.age = age
	return c
}

// Race sets the Race field
func (c *Updater) Race(race player.Race) *Updater {
	c.race = race
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	c.updatedAt = updatedAt
	return c
}

// Where applies predicates
func (u *Updater) Where(predFuncs ...comparison.PredFunc) *Updater {
	u.predFuncs = append(u.predFuncs, predFuncs...)
	return u
}

// Deleter is a delete builder
type Deleter struct {
	predFuncs []comparison.PredFunc
}

// NewDeleter returns a Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where applies predicates
func (d *Deleter) Where(predFuncs ...comparison.PredFunc) *Deleter {
	d.predFuncs = append(d.predFuncs, predFuncs...)
	return d
}

// Aggregator is an aggregate query builder
type Aggregator struct {
	v         interface{}
	aggFuncs  []aggregate.AggFunc
	predFuncs []comparison.PredFunc
	sortFuncs []sort.SortFunc
	groupBys  []Field
}

// NewAggregator expects a v and returns an Aggregator
// where 'v' argument must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{v: v}
}

// Aggregate applies aggregate functions
func (a *Aggregator) Aggregate(aggFuncs ...aggregate.AggFunc) *Aggregator {
	a.aggFuncs = append(a.aggFuncs, aggFuncs...)
	return a
}

// Where applies predicates
func (a *Aggregator) Where(predFuncs ...comparison.PredFunc) *Aggregator {
	a.predFuncs = append(a.predFuncs, predFuncs...)
	return a
}

// Sort applies sorting expressions
func (a *Aggregator) Sort(sortFuncs ...sort.SortFunc) *Aggregator {
	a.sortFuncs = append(a.sortFuncs, sortFuncs...)
	return a
}

// Group applies group clauses
func (a *Aggregator) GroupBy(fields ...Field) *Aggregator {
	a.groupBys = append(a.groupBys, fields...)
	return a
}

// rollback performs a rollback
func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}

// isZero checks if v is a zero-value
func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
