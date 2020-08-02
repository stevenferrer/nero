package pg

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPGRepoC(t *testing.T) {
	schema, err := gen.BuildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	pgRepo := NewPGRepoC(schema)
	expect := strings.TrimSpace(`
type PGRepository struct {
	db *sql.DB
}

var _ = Repository(&PGRepository{})

func NewPGRepository(db *sql.DB) *PGRepository {
	return &PGRepository{
		db: db,
	}
}

func (s *PGRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}

func (s *PGRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}
	id, err := s.CreateTx(ctx, tx, c)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return id, tx.Commit()
}

func (s *PGRepository) Query(ctx context.Context, q *Queryer) ([]*internal.Example, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return nil, err
	}

	list, err := s.QueryTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return list, tx.Commit()
}

func (s *PGRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := s.UpdateTx(ctx, tx, u)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (s *PGRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := s.DeleteTx(ctx, tx, d)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (s *PGRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.updatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	var id int64
	err := qb.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PGRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*internal.Example, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range q.pfs {
		pf(pb)
	}

	qb := squirrel.Select(q.columns...).
		From(q.collection).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	for _, p := range pb.All() {
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

func (s *PGRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range u.pfs {
		pf(pb)
	}

	qb := squirrel.Update(u.collection).
		Set("name", u.name).
		Set("updated_at", u.updatedAt).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	for _, p := range pb.All() {
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

func (s *PGRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	qb := squirrel.Delete(d.collection).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	for _, p := range pb.All() {
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

	got := strings.TrimSpace(fmt.Sprintf("%#v", pgRepo))
	assert.Equal(t, expect, got)
}
