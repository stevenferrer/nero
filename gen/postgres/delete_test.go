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
	return pg.delete(ctx, pg.db, d)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newDeleteTxBlock(t *testing.T) {
	block := newDeleteTxBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.delete(ctx, txx, d)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newDeleteRunnerBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newDeleteRunnerBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) delete(ctx context.Context, runner nero.SqlRunner, d *Deleter) (int64, error) {
	qb := squirrel.Delete("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := d.pfs
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
		case comparison.In, comparison.NotIn:
			args := p.Val.([]interface{})
			if len(args) == 0 {
				continue
			}
			qms := []string{}
			for range args {
				qms = append(qms, "?")
			}
			fmtStr := "%q IN (%s)"
			if p.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}
			plchldr := strings.Join(qms, ",")
			qb = qb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)
		}
	}

	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Delete").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
