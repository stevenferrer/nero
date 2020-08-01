package sqlite

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSQLiteRepoC(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	sqliteRepo := NewSQLiteRepoC(schema)
	expect := strings.TrimSpace(`
type SQLiteRepository struct {
	db *sql.DB
}

var _ = Repository(&SQLiteRepository{})

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (s *SQLiteRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return s.db.BeginTx(ctx, &sql.TxOptions{})
}

func (s *SQLiteRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil && tx != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "rollback error: %v", rollbackErr)
			}
			return
		}

		if err = tx.Commit(); err != nil {
			err = errors.Wrap(err, "commit error")
		}
	}()

	id, err := s.CreateTx(ctx, tx, c)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SQLiteRepository) Query(ctx context.Context, q *Queryer) ([]*internal.Example, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil && tx != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "rollback error: %v", rollbackErr)
			}
			return
		}

		if err = tx.Commit(); err != nil {
			err = errors.Wrap(err, "commit error")
		}
	}()

	list, err := s.QueryTx(ctx, tx, q)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *SQLiteRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil && tx != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "rollback error: %v", rollbackErr)
			}
			return
		}

		if err = tx.Commit(); err != nil {
			err = errors.Wrap(err, "commit error")
		}
	}()

	rowsAffected, err := s.UpdateTx(ctx, tx, u)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *SQLiteRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil && tx != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Wrapf(err, "rollback error: %v", rollbackErr)
			}
			return
		}

		if err = tx.Commit(); err != nil {
			err = errors.Wrap(err, "commit error")
		}
	}()

	rowsAffected, err := s.DeleteTx(ctx, tx, d)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *SQLiteRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.updatedAt).
		RunWith(txx)
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

func (s *SQLiteRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*internal.Example, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Builder{}
	for _, pf := range q.pfs {
		pf(pb)
	}

	qb := squirrel.Select(q.columns...).
		From(q.collection).
		RunWith(txx)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
	}

	if q.limit > 0 {
		qb = qb.Limit(q.limit)
	}

	if q.offset > 0 {
		qb = qb.Offset(q.offset)
	}

	rows, err := qb.QueryContext(ctx)
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

func (s *SQLiteRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Builder{}
	for _, pf := range u.pfs {
		pf(pb)
	}

	qb := squirrel.Update(u.collection).
		Set("name", u.name).
		Set("updated_at", u.updatedAt).
		RunWith(txx)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
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

func (s *SQLiteRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Builder{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	qb := squirrel.Delete(d.collection).
		RunWith(txx)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", sqliteRepo))
	assert.Equal(t, expect, got)
}
