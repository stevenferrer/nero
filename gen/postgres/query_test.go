package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newQueryBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newQueryBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) Query(ctx context.Context, q *Queryer) ([]*example.User, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return nil, err
	}

	list, err := pg.QueryTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return list, tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newQueryOneBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newQueryOneBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) QueryOne(ctx context.Context, q *Queryer) (*example.User, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return nil, err
	}

	item, err := pg.QueryOneTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return item, tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newQueryTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newQueryTxBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*example.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := pg.buildSelect(q)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Query").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(txx).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*example.User{}
	for rows.Next() {
		var item example.User
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Group,
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newQueryOneTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newQueryOneTxBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) QueryOneTx(ctx context.Context, tx nero.Tx, q *Queryer) (*example.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := pg.buildSelect(q)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "QueryOne").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	var item example.User
	err := qb.RunWith(txx).
		QueryRowContext(ctx).
		Scan(
			&item.ID,
			&item.Name,
			&item.Group,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newBuildSelectBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newBuildSelectBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	columns := []string{"\"id\"", "\"name\"", "\"group_res\"", "\"updated_at\"", "\"created_at\""}
	qb := squirrel.Select(columns...).
		From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := q.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case comparison.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case comparison.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case comparison.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case comparison.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case comparison.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		}
	}

	sfs := q.sfs
	sorts := &sort.Sorts{}
	for _, sf := range sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := fmt.Sprintf("%q", s.Col)
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
