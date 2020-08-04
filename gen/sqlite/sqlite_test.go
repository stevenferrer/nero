package sqlite

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSQLiteRepo(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	stmnt := NewSQLiteRepo(schema)
	expect := strings.TrimSpace(`
type SQLiteRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}

var _ = Repository(&SQLiteRepository{})

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (sqlr *SQLiteRepository) Debug(out io.Writer) *SQLiteRepository {
	lg := zerolog.New(out).With().Timestamp().Logger()
	return &SQLiteRepository{
		db:  sqlr.db,
		log: &lg,
	}
}

func (sqlr *SQLiteRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return sqlr.db.BeginTx(ctx, nil)
}

func (sqlr *SQLiteRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return 0, err
	}

	id, err := sqlr.CreateTx(ctx, tx, c)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return id, tx.Commit()
}

func (sqlr *SQLiteRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return err
	}

	err = sqlr.CreateManyTx(ctx, tx, cs...)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func (sqlr *SQLiteRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.updatedAt).
		RunWith(txx)
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Create").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (sqlr *SQLiteRepository) CreateManyTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(cs[0].collection).
		Columns(cs[0].columns...)
	for _, c := range cs {
		qb = qb.Values(c.name, c.updatedAt)
	}
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "CreateMany").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	_, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sqlr *SQLiteRepository) Query(ctx context.Context, q *Queryer) ([]*internal.Example, error) {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return nil, err
	}

	list, err := sqlr.QueryTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return list, tx.Commit()
}

func (sqlr *SQLiteRepository) QueryOne(ctx context.Context, q *Queryer) (*internal.Example, error) {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return nil, err
	}

	item, err := sqlr.QueryOneTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return item, tx.Commit()
}

func (sqlr *SQLiteRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*internal.Example, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := sqlr.buildSelect(q)
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Query").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(txx).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*internal.Example{}
	for rows.Next() {
		var item internal.Example
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &item)
	}

	return list, nil
}

func (sqlr *SQLiteRepository) QueryOneTx(ctx context.Context, tx nero.Tx, q *Queryer) (*internal.Example, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := sqlr.buildSelect(q)
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "QueryOne").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	var item internal.Example
	err := qb.RunWith(txx).
		QueryRowContext(ctx).
		Scan(
			&item.ID,
			&item.Name,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (sqlr *SQLiteRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	qb := squirrel.Select(q.columns...).
		From(q.collection)

	pb := &predicate.Predicates{}
	for _, pf := range q.pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Col: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Col: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Col: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Col: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Col: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Col: p.Val,
			})
		}
	}

	sorts := &sort.Sorts{}
	for _, sf := range q.sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := s.Col
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if q.limit > 0 {
		qb = qb.Limit(q.limit)
	}

	if q.offset > 0 {
		qb = qb.Offset(q.offset)
	}

	return qb
}

func (sqlr *SQLiteRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := sqlr.UpdateTx(ctx, tx, u)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (sqlr *SQLiteRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range u.pfs {
		pf(pb)
	}

	qb := squirrel.Update(u.collection)
	if u.name != "" {
		qb = qb.Set("name", u.name)
	}
	if u.updatedAt != nil {
		qb = qb.Set("updated_at", u.updatedAt)
	}

	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Col: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Col: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Col: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Col: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Col: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Col: p.Val,
			})
		}
	}
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Update").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (sqlr *SQLiteRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := sqlr.DeleteTx(ctx, tx, d)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (sqlr *SQLiteRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	qb := squirrel.Delete(d.collection).
		RunWith(txx)
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Col: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Col: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Col: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Col: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Col: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Col: p.Val,
			})
		}
	}
	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Delete").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (sqlr *SQLiteRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	tx, err := sqlr.Tx(ctx)
	if err != nil {
		return err
	}

	err = sqlr.AggregateTx(ctx, tx, a)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func (sqlr *SQLiteRepository) AggregateTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	aggs := &aggregate.Aggregates{}
	for _, aggf := range a.aggfs {
		aggf(aggs)
	}
	cols := []string{}
	for _, agg := range aggs.All() {
		col := agg.Col
		switch agg.Fn {
		case aggregate.Avg:
			cols = append(cols, "AVG("+col+") avg_"+col)
		case aggregate.Count:
			cols = append(cols, "COUNT("+col+") count_"+col)
		case aggregate.Max:
			cols = append(cols, "MAX("+col+") max_"+col)
		case aggregate.Min:
			cols = append(cols, "MIN("+col+") min_"+col)
		case aggregate.Sum:
			cols = append(cols, "SUM("+col+") sum_"+col)
		}
	}

	groups := []string{}
	for _, group := range a.groups {
		groups = append(groups, group.String())
	}

	qb := squirrel.Select(cols...).
		From(a.collection).GroupBy(groups...)

	preds := &predicate.Predicates{}
	for _, pf := range a.pfs {
		pf(preds)
	}
	for _, p := range preds.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Col: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Col: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Col: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Col: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Col: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Col: p.Val,
			})
		}
	}

	sorts := &sort.Sorts{}
	for _, sf := range a.sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := s.Col
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if log := sqlr.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Aggregate").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(txx).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	dv := reflect.ValueOf(a.dest).Elem()
	dt := reflect.TypeOf(dv.Interface()).Elem()
	if dt.NumField() != len(cols) {
		return errors.New("aggregate columns and destination struct field count should match")
	}

	for rows.Next() {
		de := reflect.New(dt).Elem()
		dest := make([]interface{}, de.NumField())
		for i := 0; i < de.NumField(); i++ {
			dest[i] = de.Field(i).Addr().Interface()
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		dv.Set(reflect.Append(dv, de))
	}

	return nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", stmnt))
	assert.Equal(t, expect, got)
}

func Test_newTypeDefBlock(t *testing.T) {
	block := newTypeDefBlock()
	expect := strings.TrimSpace(`
type SQLiteRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newTypeAssertBlock(t *testing.T) {
	block := newTypeAssertBlock()
	expect := `var _ = Repository(&SQLiteRepository{})`
	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newDebugLogBlock(t *testing.T) {
	block := newDebugLogBlock("Query")
	expect := strings.TrimSpace(`
if log := sqlr.log; log != nil {
	sql, args, err := qb.ToSql()
	log.Debug().Str("op", "Query").Str("stmnt", sql).
		Interface("args", args).Err(err).Msg("")
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
