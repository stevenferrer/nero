package user_test

import (
	"fmt"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/comparison"
	"github.com/sf9v/nero/test/integration/basic/repository/user"
)

func TestPredicates(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		pfs := []user.PredFunc{
			user.IDEq("1"), user.IDNotEq("1"),
			user.IDGt("1"), user.IDGtOrEq("1"),
			user.IDLt("1"), user.IDLtOrEq("1"),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("id").From("users")
		for _, p := range pb.All() {
			require.Equal(t, "1", p.Val)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := `SELECT id FROM users WHERE "id" = ? AND "id" <> ? AND "id" > ? AND "id" >= ? AND "id" < ? AND "id" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("Name", func(t *testing.T) {
		name := "sf9v"
		pfs := []user.PredFunc{
			user.NameEq(name), user.NameNotEq(name),
			user.NameGt(name), user.NameGtOrEq(name),
			user.NameLt(name), user.NameLtOrEq(name),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("name").From("users")
		for _, p := range pb.All() {
			s, ok := p.Val.(string)
			assert.True(t, ok)
			assert.Equal(t, name, s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT name FROM users WHERE "name" = ? AND "name" <> ? AND "name" > ? AND "name" >= ? AND "name" < ? AND "name" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("Email", func(t *testing.T) {
		email := "sf9v@gg.io"
		pfs := []user.PredFunc{
			user.EmailEq(email), user.EmailNotEq(email),
			user.EmailGt(email), user.EmailGtOrEq(email),
			user.EmailLt(email), user.EmailLtOrEq(email),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("email").From("users")
		for _, p := range pb.All() {
			s, ok := p.Val.(string)
			assert.True(t, ok)
			assert.Equal(t, email, s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT email FROM users WHERE "email" = ? AND "email" <> ? AND "email" > ? AND "email" >= ? AND "email" < ? AND "email" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []user.PredFunc{
			user.CreatedAtEq(&now), user.CreatedAtNotEq(&now),
			user.CreatedAtGt(&now), user.CreatedAtGtOrEq(&now),
			user.CreatedAtLt(&now), user.CreatedAtLtOrEq(&now),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("created_at").From("users")
		for _, p := range pb.All() {
			assert.NotNil(t, p.Val)
			_, ok := p.Val.(*time.Time)
			assert.True(t, ok)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT created_at FROM users WHERE "created_at" = ? AND "created_at" <> ? AND "created_at" > ? AND "created_at" >= ? AND "created_at" < ? AND "created_at" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("UpdatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []user.PredFunc{
			user.UpdatedAtEq(&now), user.UpdatedAtNotEq(&now),
			user.UpdatedAtGt(&now), user.UpdatedAtGtOrEq(&now),
			user.UpdatedAtLt(&now), user.UpdatedAtLtOrEq(&now),
			user.UpdatedAtIsNull(), user.UpdatedAtIsNotNull(),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("updated_at").From("users")
		for _, p := range pb.All() {
			if p.Op == comparison.IsNull ||
				p.Op == comparison.IsNotNull {
				assert.Nil(t, p.Val)
			} else {
				assert.NotNil(t, p.Val)
				_, ok := p.Val.(*time.Time)
				assert.True(t, ok)
			}

			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT updated_at FROM users WHERE "updated_at" = ? AND "updated_at" <> ? AND "updated_at" > ? AND "updated_at" >= ? AND "updated_at" < ? AND "updated_at" <= ? AND "updated_at" IS NULL AND "updated_at" IS NOT NULL`
		assert.Equal(t, expect, got)
	})
}

func addPred(sb sq.SelectBuilder,
	p *comparison.Predicate) sq.SelectBuilder {
	switch p.Op {
	case comparison.Eq:
		return sb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
	case comparison.NotEq:
		return sb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
	case comparison.Gt:
		return sb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
	case comparison.GtOrEq:
		return sb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
	case comparison.Lt:
		return sb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
	case comparison.LtOrEq:
		return sb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
	case comparison.IsNull:
		return sb.Where(fmt.Sprintf("%q IS NULL", p.Col))
	case comparison.IsNotNull:
		return sb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
	}

	return sb
}
