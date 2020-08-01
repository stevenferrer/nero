// Code generated by nero, DO NOT EDIT.
package user

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	errors "github.com/pkg/errors"
	nero "github.com/sf9v/nero"
	predicate "github.com/sf9v/nero/predicate"
	user "github.com/sf9v/nero/test/integration/user"
	"strconv"
)

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

func (s *SQLiteRepository) Create(ctx context.Context, c *Creator) (string, error) {
	tx, err := s.Tx(ctx)
	if err != nil {
		return "", err
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
		return "", err
	}

	return id, nil
}

func (s *SQLiteRepository) Query(ctx context.Context, q *Queryer) ([]*user.User, error) {
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

func (s *SQLiteRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (string, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return "", errors.New("expecting tx to be *sql.Tx")
	}

	qb := sq.Insert(c.collection).
		Columns(c.columns...).
		Values(c.email, c.name, c.updatedAt).
		RunWith(txx)
	res, err := qb.ExecContext(ctx)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func (s *SQLiteRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*user.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range q.pfs {
		pf(pb)
	}

	qb := sq.Select(q.columns...).
		From(q.collection).
		RunWith(txx)
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(sq.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(sq.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(sq.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(sq.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(sq.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(sq.LtOrEq{
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

	list := []*user.User{}
	for rows.Next() {
		var item user.User
		err = rows.Scan(
			&item.ID,
			&item.Email,
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

	pb := &predicate.Predicates{}
	for _, pf := range u.pfs {
		pf(pb)
	}

	qb := sq.Update(u.collection).
		Set("email", u.email).
		Set("name", u.name).
		Set("updated_at", u.updatedAt).
		RunWith(txx)
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(sq.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(sq.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(sq.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(sq.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(sq.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(sq.LtOrEq{
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

	pb := &predicate.Predicates{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	qb := sq.Delete(d.collection).
		RunWith(txx)
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(sq.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			qb = qb.Where(sq.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			qb = qb.Where(sq.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			qb = qb.Where(sq.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			qb = qb.Where(sq.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			qb = qb.Where(sq.LtOrEq{
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
