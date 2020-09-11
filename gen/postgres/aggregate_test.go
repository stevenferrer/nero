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

func Test_newAggregateBlock(t *testing.T) {
	aggBlocks := newAggregateBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	return pg.aggregate(ctx, pg.db, a)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", aggBlocks))
	assert.Equal(t, expect, got)
}

func Test_newAggregateTxBlock(t *testing.T) {
	aggBlocks := newAggregateTxBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) AggregateTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return pg.aggregate(ctx, txx, a)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", aggBlocks))
	assert.Equal(t, expect, got)
}

func Test_newAggregateRunnerBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	aggBlocks := newAggregateRunnerBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) aggregate(ctx context.Context, runner nero.SqlRunner, a *Aggregator) error {
	aggs := &aggregate.Aggregates{}
	for _, aggf := range a.aggfs {
		aggf(aggs)
	}
	cols := []string{}
	for _, agg := range aggs.All() {
		col := agg.Col
		qcol := fmt.Sprintf("%q", col)
		switch agg.Fn {
		case aggregate.Avg:
			cols = append(cols, "AVG("+qcol+") avg_"+col)
		case aggregate.Count:
			cols = append(cols, "COUNT("+qcol+") count_"+col)
		case aggregate.Max:
			cols = append(cols, "MAX("+qcol+") max_"+col)
		case aggregate.Min:
			cols = append(cols, "MIN("+qcol+") min_"+col)
		case aggregate.Sum:
			cols = append(cols, "SUM("+qcol+") sum_"+col)
		case aggregate.None:
			cols = append(cols, qcol)
		}
	}

	qb := squirrel.Select(cols...).From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	groups := []string{}
	for _, group := range a.groups {
		groups = append(groups, fmt.Sprintf("%q", group.String()))
	}
	qb = qb.GroupBy(groups...)

	pfs := a.pfs
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

	sfs := a.sfs
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

	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Aggregate").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	v := goreflect.ValueOf(a.v).Elem()
	t := goreflect.TypeOf(v.Interface()).Elem()
	if t.NumField() != len(cols) {
		return errors.New("aggregate columns and destination struct field count should match")
	}

	for rows.Next() {
		ve := goreflect.New(t).Elem()
		dest := make([]interface{}, ve.NumField())
		for i := 0; i < ve.NumField(); i++ {
			dest[i] = ve.Field(i).Addr().Interface()
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		v.Set(goreflect.Append(v, ve))
	}

	return nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", aggBlocks))
	assert.Equal(t, expect, got)
}
