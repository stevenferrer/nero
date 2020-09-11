package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPredicatesBlock(t *testing.T) {
	block := newPredicatesBlock()
	expect := strings.TrimSpace(`
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
