package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newDeleteBlock(t *testing.T) {
	block := newDeleteBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := pg.DeleteTx(ctx, tx, d)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newDeleteTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newDeleteTxBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	pb := &predicate.Predicates{}
	for _, pf := range d.pfs {
		pf(pb)
	}

	qb := squirrel.Delete("\"users\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	for _, p := range pb.All() {
		switch p.Op {
		case predicate.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case predicate.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case predicate.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case predicate.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case predicate.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case predicate.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		}
	}
	if log := pg.log; log != nil {
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
