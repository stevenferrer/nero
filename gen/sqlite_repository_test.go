package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newSQLiteRepo(t *testing.T) {
	schema, err := buildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	sqliteRepo := newSQLiteRepo(schema)
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

func (s *SQLiteRepository) Create(c *Creator) (int64, error) {
	sb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.updatedAt)

	sql, args, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	stmnt, err := s.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmnt.Exec(args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SQLiteRepository) Query(q *Queryer) ([]*internal.Example, error) {
	pb := &predicate.Builder{}
	for _, pf := range q.pfs {
		pf(pb)
	}

	sb := squirrel.Select(q.columns...).From(q.collection)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			sb = sb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			sb = sb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			sb = sb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			sb = sb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			sb = sb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			sb = sb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
	}

	if q.limit > 0 {
		sb = sb.Limit(q.limit)
	}

	if q.offset > 0 {
		sb = sb.Offset(q.offset)
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	stmnt, err := s.db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmnt.Query(args...)
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

func (s *SQLiteRepository) Update(u *Updater) (int64, error) {
	pb := &predicate.Builder{}
	for _, pf := range u.pfs {
		pf(pb)
	}

	sb := squirrel.Update(u.collection).Set("name", u.name).Set("updated_at", u.updatedAt)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			sb = sb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			sb = sb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			sb = sb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			sb = sb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			sb = sb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			sb = sb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	stmnt, err := s.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmnt.Exec(args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *SQLiteRepository) Delete(d *Deleter) (int64, error) {
	pb := &predicate.Builder{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	sb := squirrel.Delete(d.collection)
	for _, p := range pb.Predicates() {
		switch p.Op {
		case predicate.Eq:
			sb = sb.Where(squirrel.Eq{
				p.Field: p.Val,
			})
		case predicate.NotEq:
			sb = sb.Where(squirrel.NotEq{
				p.Field: p.Val,
			})
		case predicate.Gt:
			sb = sb.Where(squirrel.Gt{
				p.Field: p.Val,
			})
		case predicate.GtOrEq:
			sb = sb.Where(squirrel.GtOrEq{
				p.Field: p.Val,
			})
		case predicate.Lt:
			sb = sb.Where(squirrel.Lt{
				p.Field: p.Val,
			})
		case predicate.LtOrEq:
			sb = sb.Where(squirrel.LtOrEq{
				p.Field: p.Val,
			})
		}
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	stmnt, err := s.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmnt.Exec(args...)
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
