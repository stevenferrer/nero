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

func Test_newUpdateBlock(t *testing.T) {
	block := newUpdateBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	return pg.update(ctx, pg.db, u)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newUpdateTxBlock(t *testing.T) {
	block := newUpdateTxBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.update(ctx, txx, u)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newUpdateRunnerBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newUpdateRunnerBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) update(ctx context.Context, runner nero.SqlRunner, u *Updater) (int64, error) {
	qb := squirrel.Update("\"users\"").PlaceholderFormat(squirrel.Dollar)
	if u.name != "" {
		qb = qb.Set("\"name\"", u.name)
	}
	if u.group != "" {
		qb = qb.Set("\"group_res\"", u.group)
	}
	if u.updatedAt != nil {
		qb = qb.Set("\"updated_at\"", u.updatedAt)
	}

	pfs := u.pfs
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
		log.Debug().Str("method", "Update").Str("stmnt", sql).
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
