// Code generated by nero, DO NOT EDIT.
package playerrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/aggregate"
	"github.com/stevenferrer/nero/predicate"
	"github.com/stevenferrer/nero/sorting"
	"github.com/stevenferrer/nero/test/demo-test/playerpkg"
)

// SQLiteRepository is a repository that uses SQLite3 as data store
type SQLiteRepository struct {
	db     *sql.DB
	logger nero.Logger
	debug  bool
}

var _ Repository = (*SQLiteRepository)(nil)

// NewSQLiteRepository returns a new SQLiteRepository
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// Debug enables debug mode
func (repo *SQLiteRepository) Debug() *SQLiteRepository {
	l := log.New(os.Stdout, "[nero] ", log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
	return &SQLiteRepository{
		db:     repo.db,
		debug:  true,
		logger: l,
	}
}

// WithLogger overrides the default logger
func (repo *SQLiteRepository) WithLogger(logger nero.Logger) *SQLiteRepository {
	repo.logger = logger
	return repo
}

// BeginTx starts a transaction
func (repo *SQLiteRepository) BeginTx(ctx context.Context) (nero.Tx, error) {
	return repo.db.BeginTx(ctx, nil)
}

// Create creates a Player
func (repo *SQLiteRepository) Create(ctx context.Context, c *Creator) (string, error) {
	return repo.create(ctx, repo.db, c)
}

// CreateInTx creates a Player in a transaction
func (repo *SQLiteRepository) CreateInTx(ctx context.Context, tx nero.Tx, c *Creator) (string, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return "", errors.New("expecting tx to be *sql.Tx")
	}

	return repo.create(ctx, txx, c)
}

func (repo *SQLiteRepository) create(ctx context.Context, runner nero.SQLRunner, c *Creator) (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}

	columns := []string{
		"\"email\"",
		"\"name\"",
		"\"age\"",
		"\"race\"",
	}

	values := []interface{}{
		c.email,
		c.name,
		c.age,
		c.race,
	}

	if !isZero(c.updatedAt) {
		columns = append(columns, "updated_at")
		values = append(values, c.updatedAt)
	}

	qb := squirrel.Insert("\"players\"").Columns(columns...).
		Values(values...).RunWith(runner)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Create, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	_, err := qb.ExecContext(ctx)
	if err != nil {
		return "", err
	}

	var id string
	err = repo.db.QueryRowContext(ctx, "select last_insert_rowid()").Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// CreateMany batch creates Players
func (repo *SQLiteRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	return repo.createMany(ctx, repo.db, cs...)
}

// CreateManyInTx batch creates Players in a transaction
func (repo *SQLiteRepository) CreateManyInTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return repo.createMany(ctx, txx, cs...)
}

func (repo *SQLiteRepository) createMany(ctx context.Context, runner nero.SQLRunner, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	columns := []string{
		"\"email\"",
		"\"name\"",
		"\"age\"",
		"\"race\"",
		"\"updated_at\"",
	}
	qb := squirrel.Insert("\"players\"").Columns(columns...)
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}

		qb = qb.Values(
			c.email,
			c.name,
			c.age,
			c.race,
			c.updatedAt,
		)
	}

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: CreateMany, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	_, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Query queries Players
func (repo *SQLiteRepository) Query(ctx context.Context, q *Queryer) ([]playerpkg.Player, error) {
	return repo.query(ctx, repo.db, q)
}

// QueryInTx queries Players in a transaction
func (repo *SQLiteRepository) QueryInTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]playerpkg.Player, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.query(ctx, txx, q)
}

func (repo *SQLiteRepository) query(ctx context.Context, runner nero.SQLRunner, q *Queryer) ([]playerpkg.Player, error) {
	qb := repo.buildSelect(q)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Query, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []playerpkg.Player{}
	for rows.Next() {
		var player playerpkg.Player
		err = rows.Scan(
			&player.ID,
			&player.Email,
			&player.Name,
			&player.Age,
			&player.Race,
			&player.UpdatedAt,
			&player.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

// QueryOne queries a Player
func (repo *SQLiteRepository) QueryOne(ctx context.Context, q *Queryer) (playerpkg.Player, error) {
	return repo.queryOne(ctx, repo.db, q)
}

// QueryOneInTx queries a Player in a transaction
func (repo *SQLiteRepository) QueryOneInTx(ctx context.Context, tx nero.Tx, q *Queryer) (playerpkg.Player, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return (playerpkg.Player{}), errors.New("expecting tx to be *sql.Tx")
	}

	return repo.queryOne(ctx, txx, q)
}

func (repo *SQLiteRepository) queryOne(ctx context.Context, runner nero.SQLRunner, q *Queryer) (playerpkg.Player, error) {
	qb := repo.buildSelect(q)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: QueryOne, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	var player playerpkg.Player
	err := qb.RunWith(runner).
		QueryRowContext(ctx).
		Scan(
			&player.ID,
			&player.Email,
			&player.Name,
			&player.Age,
			&player.Race,
			&player.UpdatedAt,
			&player.CreatedAt,
		)
	if err != nil {
		return (playerpkg.Player{}), err
	}

	return player, nil
}

func (repo *SQLiteRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	columns := []string{
		"\"id\"",
		"\"email\"",
		"\"name\"",
		"\"age\"",
		"\"race\"",
		"\"updated_at\"",
		"\"created_at\"",
	}
	qb := squirrel.Select(columns...).From("\"players\"")

	predicates := predicate.Build(q.predFuncs...)
	qb = squirrel.SelectBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), predicates))

	sortings := sorting.Build(q.sortFuncs...)
	qb = repo.buildSorting(qb, sortings)

	if q.limit > 0 {
		qb = qb.Limit(uint64(q.limit))
	}

	if q.offset > 0 {
		qb = qb.Offset(uint64(q.offset))
	}

	return qb
}

func (repo *SQLiteRepository) buildPreds(sb squirrel.StatementBuilderType, preds []predicate.Predicate) squirrel.StatementBuilderType {
	for _, pred := range preds {
		ph := "?"
		fieldX, arg := pred.Field, pred.Argument

		args := []interface{}{}
		if fieldY, ok := arg.(Field); ok { // a field
			ph = fmt.Sprintf("%q", fieldY)
		} else if vals, ok := arg.([]interface{}); ok { // array of values
			args = append(args, vals...)
		} else { // single value
			args = append(args, arg)
		}

		switch pred.Operator {
		case predicate.Eq:
			sb = sb.Where(fmt.Sprintf("%q = "+ph, fieldX), args...)
		case predicate.NotEq:
			sb = sb.Where(fmt.Sprintf("%q <> "+ph, fieldX), args...)
		case predicate.Gt:
			sb = sb.Where(fmt.Sprintf("%q > "+ph, fieldX), args...)
		case predicate.GtOrEq:
			sb = sb.Where(fmt.Sprintf("%q >= "+ph, fieldX), args...)
		case predicate.Lt:
			sb = sb.Where(fmt.Sprintf("%q < "+ph, fieldX), args...)
		case predicate.LtOrEq:
			sb = sb.Where(fmt.Sprintf("%q <= "+ph, fieldX), args...)
		case predicate.IsNull, predicate.IsNotNull:
			fmtStr := "%q IS NULL"
			if pred.Operator == predicate.IsNotNull {
				fmtStr = "%q IS NOT NULL"
			}
			sb = sb.Where(fmt.Sprintf(fmtStr, fieldX))
		case predicate.In, predicate.NotIn:
			fmtStr := "%q IN (%s)"
			if pred.Operator == predicate.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}

			phs := make([]string, 0, len(args))
			for range args {
				phs = append(phs, "?")
			}

			sb = sb.Where(fmt.Sprintf(fmtStr, fieldX, strings.Join(phs, ",")), args...)
		}
	}

	return sb
}

func (repo *SQLiteRepository) buildSorting(qb squirrel.SelectBuilder, sortings []sorting.Sorting) squirrel.SelectBuilder {
	for _, s := range sortings {
		field := fmt.Sprintf("%q", s.Field)
		switch s.Direction {
		case sorting.Asc:
			qb = qb.OrderBy(field + " ASC")
		case sorting.Desc:
			qb = qb.OrderBy(field + " DESC")
		}
	}

	return qb
}

// Update updates a Player or many Players
func (repo *SQLiteRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	return repo.update(ctx, repo.db, u)
}

// UpdateInTx updates a Player many Players in a transaction
func (repo *SQLiteRepository) UpdateInTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.update(ctx, txx, u)
}

func (repo *SQLiteRepository) update(ctx context.Context, runner nero.SQLRunner, u *Updater) (int64, error) {
	qb := squirrel.Update("\"players\"")

	cnt := 0

	if !isZero(u.email) {
		qb = qb.Set("\"email\"", u.email)
		cnt++
	}

	if !isZero(u.name) {
		qb = qb.Set("\"name\"", u.name)
		cnt++
	}

	if !isZero(u.age) {
		qb = qb.Set("\"age\"", u.age)
		cnt++
	}

	if !isZero(u.race) {
		qb = qb.Set("\"race\"", u.race)
		cnt++
	}

	if !isZero(u.updatedAt) {
		qb = qb.Set("\"updated_at\"", u.updatedAt)
		cnt++
	}

	if cnt == 0 {
		return 0, nil
	}

	predicates := predicate.Build(u.predFuncs...)
	qb = squirrel.UpdateBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), predicates))

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Update, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Delete deletes a Player or many Players
func (repo *SQLiteRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	return repo.delete(ctx, repo.db, d)
}

// DeleteInTx deletes a Player or many Players in a transaction
func (repo *SQLiteRepository) DeleteInTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.delete(ctx, txx, d)
}

func (repo *SQLiteRepository) delete(ctx context.Context, runner nero.SQLRunner, d *Deleter) (int64, error) {
	qb := squirrel.Delete("\"players\"")

	predicates := predicate.Build(d.predFuncs...)
	qb = squirrel.DeleteBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), predicates))

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Delete, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Aggregate runs an aggregate query
func (repo *SQLiteRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	return repo.aggregate(ctx, repo.db, a)
}

// AggregateInTx runs an aggregate query in a transaction
func (repo *SQLiteRepository) AggregateInTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return repo.aggregate(ctx, txx, a)
}

func (repo *SQLiteRepository) aggregate(ctx context.Context, runner nero.SQLRunner, a *Aggregator) error {
	aggregates := aggregate.Build(a.aggFuncs...)
	columns := make([]string, 0, len(aggregates))
	for _, agg := range aggregates {
		field := agg.Field
		qf := fmt.Sprintf("%q", field)
		switch agg.Operator {
		case aggregate.Avg:
			columns = append(columns, "AVG("+qf+") avg_"+field)
		case aggregate.Count:
			columns = append(columns, "COUNT("+qf+") count_"+field)
		case aggregate.Max:
			columns = append(columns, "MAX("+qf+") max_"+field)
		case aggregate.Min:
			columns = append(columns, "MIN("+qf+") min_"+field)
		case aggregate.Sum:
			columns = append(columns, "SUM("+qf+") sum_"+field)
		case aggregate.None:
			columns = append(columns, qf)
		}
	}

	qb := squirrel.Select(columns...).From("\"players\"")

	groupBys := make([]string, 0, len(a.groupBys))
	for _, groupBy := range a.groupBys {
		groupBys = append(groupBys, fmt.Sprintf("%q", groupBy.String()))
	}
	qb = qb.GroupBy(groupBys...)

	predicates := predicate.Build(a.predFuncs...)
	qb = squirrel.SelectBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), predicates))

	sortings := sorting.Build(a.sortFuncs...)
	qb = repo.buildSorting(qb, sortings)

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Aggregate, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	v := reflect.ValueOf(a.v).Elem()
	t := reflect.TypeOf(v.Interface()).Elem()
	if len(columns) != t.NumField() {
		return errors.Errorf("column count (%v) and destination struct field count (%v) doesn't match", len(columns), t.NumField())
	}

	for rows.Next() {
		ve := reflect.New(t).Elem()
		dest := make([]interface{}, ve.NumField())
		for i := 0; i < ve.NumField(); i++ {
			dest[i] = ve.Field(i).Addr().Interface()
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		v.Set(reflect.Append(v, ve))
	}

	return nil
}