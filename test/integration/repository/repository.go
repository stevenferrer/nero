// Code generated by nero, DO NOT EDIT.
package repository

import (
	"context"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero"
	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/test/integration/user"
)

// Repository is a repository for User
type Repository interface {
	// Tx begins a new transaction
	Tx(context.Context) (nero.Tx, error)
	// Create creates a new User
	Create(context.Context, *Creator) (id string, err error)
	// CreateTx creates a new type User inside a transaction
	CreateTx(context.Context, nero.Tx, *Creator) (id string, err error)
	// CreateMany creates many User
	CreateMany(context.Context, ...*Creator) error
	// CreateManyTx creates many User inside a transaction
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries many User
	Query(context.Context, *Queryer) ([]*user.User, error)
	// QueryTx queries many User inside a transaction
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*user.User, error)
	// QueryOne queries one User
	QueryOne(context.Context, *Queryer) (*user.User, error)
	// QueryOneTx queries one User inside a transaction
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*user.User, error)
	// Update updates User
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates User inside a transaction
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes User
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes User inside a transaction
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate performs aggregate query
	Aggregate(context.Context, *Aggregator) error
	// Aggregate performs aggregate query inside a transaction
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}

// Creator is a create builder for User
type Creator struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	tags      []string
	updatedAt *time.Time
}

// NewCreator is a factory for Creator
func NewCreator() *Creator {
	return &Creator{}
}

// UID is a setter for the UID field
func (c *Creator) UID(uid ksuid.KSUID) *Creator {
	c.uid = uid
	return c
}

// Email is a setter for the Email field
func (c *Creator) Email(email string) *Creator {
	c.email = email
	return c
}

// Name is a setter for the Name field
func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

// Age is a setter for the Age field
func (c *Creator) Age(age int) *Creator {
	c.age = age
	return c
}

// Group is a setter for the Group field
func (c *Creator) Group(group user.Group) *Creator {
	c.group = group
	return c
}

// Kv is a setter for the Kv field
func (c *Creator) Kv(kv example.Map) *Creator {
	c.kv = kv
	return c
}

// Tags is a setter for the Tags field
func (c *Creator) Tags(tags []string) *Creator {
	c.tags = tags
	return c
}

// UpdatedAt is a setter for the UpdatedAt field
func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}

// Validate validates the creator fields
func (c *Creator) Validate() error {
	if isZero(c.uid) {
		return nero.NewErrRequiredField("uid")
	}

	if isZero(c.email) {
		return nero.NewErrRequiredField("email")
	}

	if isZero(c.name) {
		return nero.NewErrRequiredField("name")
	}

	if isZero(c.age) {
		return nero.NewErrRequiredField("age")
	}

	if isZero(c.group) {
		return nero.NewErrRequiredField("group")
	}

	if isZero(c.kv) {
		return nero.NewErrRequiredField("kv")
	}

	if isZero(c.tags) {
		return nero.NewErrRequiredField("tags")
	}

	return nil
}

// Queryer is a query builder for User
type Queryer struct {
	limit  uint
	offset uint
	pfs    []PredFunc
	sfs    []SortFunc
}

// NewQueryer is a factory for Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where adds predicates to the query
func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

// Sort adds sorting expressions to the query
func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

// Limit adds limit clause to the query
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset adds offset clause to the query
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder for User
type Updater struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	tags      []string
	updatedAt *time.Time
	pfs       []PredFunc
}

// NewUpdater is a factory for Updater
func NewUpdater() *Updater {
	return &Updater{}
}

// UID is a setter for the UID field
func (c *Updater) UID(uid ksuid.KSUID) *Updater {
	c.uid = uid
	return c
}

// Email is a setter for the Email field
func (c *Updater) Email(email string) *Updater {
	c.email = email
	return c
}

// Name is a setter for the Name field
func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

// Age is a setter for the Age field
func (c *Updater) Age(age int) *Updater {
	c.age = age
	return c
}

// Group is a setter for the Group field
func (c *Updater) Group(group user.Group) *Updater {
	c.group = group
	return c
}

// Kv is a setter for the Kv field
func (c *Updater) Kv(kv example.Map) *Updater {
	c.kv = kv
	return c
}

// Tags is a setter for the Tags field
func (c *Updater) Tags(tags []string) *Updater {
	c.tags = tags
	return c
}

// UpdatedAt is a setter for the UpdatedAt field
func (c *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	c.updatedAt = updatedAt
	return c
}

// Where adds predicates to the update builder
func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}

// Deleter is a delete builder for User
type Deleter struct {
	pfs []PredFunc
}

// NewDeleter is a factory for Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where adds predicates to the delete builder
func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}

// Aggregator is an aggregate builder for User
type Aggregator struct {
	v      interface{}
	aggfs  []AggFunc
	pfs    []PredFunc
	sfs    []SortFunc
	groups []Column
}

// NewAggregator is a factory for Aggregator
// 'v' argument must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{
		v: v,
	}
}

// Aggregate adds aggregate functions to the aggregate builder
func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

// Where adds predicates to the aggregate builder
func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

// Sort adds sorting expressions to the aggregate builder
func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

// Group adds grouping clause to the aggregate builder
func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
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

// isZero checks of value is zero
func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
