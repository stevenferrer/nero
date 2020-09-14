package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sf9v/nero"

	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/test/integration/user"
	"time"
)

type Repository interface {
	Tx(context.Context) (nero.Tx, error)
	Create(context.Context, *Creator) (id string, err error)
	CreateTx(context.Context, nero.Tx, *Creator) (id string, err error)
	CreateMany(context.Context, ...*Creator) error
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	Query(context.Context, *Queryer) ([]*user.User, error)
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*user.User, error)
	QueryOne(context.Context, *Queryer) (*user.User, error)
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*user.User, error)
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	Aggregate(context.Context, *Aggregator) error
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}

type Creator struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	updatedAt *time.Time
}

func NewCreator() *Creator {
	return &Creator{}
}

func (c *Creator) UID(uid ksuid.KSUID) *Creator {
	c.uid = uid
	return c
}

func (c *Creator) Email(email string) *Creator {
	c.email = email
	return c
}

func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

func (c *Creator) Age(age int) *Creator {
	c.age = age
	return c
}

func (c *Creator) Group(group user.Group) *Creator {
	c.group = group
	return c
}

func (c *Creator) Kv(kv example.Map) *Creator {
	c.kv = kv
	return c
}

func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}

type Queryer struct {
	limit  uint
	offset uint
	pfs    []PredFunc
	sfs    []SortFunc
}

func NewQueryer() *Queryer {
	return &Queryer{}
}

func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

type Updater struct {
	uid       ksuid.KSUID
	email     string
	name      string
	age       int
	group     user.Group
	kv        example.Map
	updatedAt *time.Time
	pfs       []PredFunc
}

func NewUpdater() *Updater {
	return &Updater{}
}

func (c *Updater) UID(uid ksuid.KSUID) *Updater {
	c.uid = uid
	return c
}

func (c *Updater) Email(email string) *Updater {
	c.email = email
	return c
}

func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

func (c *Updater) Age(age int) *Updater {
	c.age = age
	return c
}

func (c *Updater) Group(group user.Group) *Updater {
	c.group = group
	return c
}

func (c *Updater) Kv(kv example.Map) *Updater {
	c.kv = kv
	return c
}

func (c *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	c.updatedAt = updatedAt
	return c
}

func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}

type Deleter struct {
	pfs []PredFunc
}

func NewDeleter() *Deleter {
	return &Deleter{}
}

func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}

type Aggregator struct {
	v      interface{}
	aggfs  []AggFunc
	pfs    []PredFunc
	sfs    []SortFunc
	groups []Column
}

func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{
		v: v,
	}
}

func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
	return a
}

func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}
