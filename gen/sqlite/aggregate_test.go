package sqlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gen "github.com/sf9v/nero/gen/internal"

	"github.com/sf9v/nero/example"
)

func Test_newAggregateBlock(t *testing.T) {
	aggBlocks := newAggregateBlock()
	expect := strings.TrimSpace(`
func (sl *SQLiteRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	tx, err := sl.Tx(ctx)
	if err != nil {
		return err
	}

	err = sl.AggregateTx(ctx, tx, a)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", aggBlocks))
	assert.Equal(t, expect, got)
}

func Test_newAggregateTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	aggBlocks := newAggregateTxBlock()
	expect := strings.TrimSpace(`
func (sl *SQLiteRepository) AggregateTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
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
		case aggregate.None:
			cols = append(cols, col)
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

	if log := sl.log; log != nil {
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

	got := strings.TrimSpace(fmt.Sprintf("%#v", aggBlocks))
	assert.Equal(t, expect, got)
}
