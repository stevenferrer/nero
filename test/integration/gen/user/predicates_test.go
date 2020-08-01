package user_test

import (
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/predicate"
	"github.com/sf9v/nero/test/integration/gen/user"
)

func TestPredicates(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		pfs := []user.PredicateFunc{
			user.IDEq(1), user.IDNotEq(1),
			user.IDGt(1), user.IDGtOrEq(1),
			user.IDLt(1), user.IDLtOrEq(1),
		}

		pb := &predicate.Builder{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("id").From("users")
		for _, p := range pb.Predicates() {
			require.Equal(t, int64(1), p.Val)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT id FROM users WHERE id = ? AND id <> ? AND id > ? AND id >= ? AND id < ? AND id <= ?"
		assert.Equal(t, expect, got)
	})

	t.Run("Name", func(t *testing.T) {
		name := "sf9v"
		pfs := []user.PredicateFunc{
			user.NameEq(&name), user.NameNotEq(&name),
			user.NameGt(&name), user.NameGtOrEq(&name),
			user.NameLt(&name), user.NameLtOrEq(&name),
		}

		pb := &predicate.Builder{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("name").From("users")
		for _, p := range pb.Predicates() {
			s, ok := p.Val.(*string)
			require.True(t, ok)
			require.Equal(t, name, *s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT name FROM users WHERE name = ? AND name <> ? AND name > ? AND name >= ? AND name < ? AND name <= ?"
		assert.Equal(t, expect, got)
	})

	t.Run("Email", func(t *testing.T) {
		email := "sf9v@gg.io"
		pfs := []user.PredicateFunc{
			user.EmailEq(&email), user.EmailNotEq(&email),
			user.EmailGt(&email), user.EmailGtOrEq(&email),
			user.EmailLt(&email), user.EmailLtOrEq(&email),
		}

		pb := &predicate.Builder{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("email").From("users")
		for _, p := range pb.Predicates() {
			s, ok := p.Val.(*string)
			require.True(t, ok)
			require.Equal(t, email, *s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT email FROM users WHERE email = ? AND email <> ? AND email > ? AND email >= ? AND email < ? AND email <= ?"
		assert.Equal(t, expect, got)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []user.PredicateFunc{
			user.CreatedAtEq(&now), user.CreatedAtNotEq(&now),
			user.CreatedAtGt(&now), user.CreatedAtGtOrEq(&now),
			user.CreatedAtLt(&now), user.CreatedAtLtOrEq(&now),
		}

		pb := &predicate.Builder{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("created_at").From("users")
		for _, p := range pb.Predicates() {
			require.NotNil(t, p.Val)
			_, ok := p.Val.(*time.Time)
			require.True(t, ok)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT created_at FROM users WHERE created_at = ? AND created_at <> ? AND created_at > ? AND created_at >= ? AND created_at < ? AND created_at <= ?"
		assert.Equal(t, expect, got)
	})

	t.Run("UpdatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []user.PredicateFunc{
			user.UpdatedAtEq(&now), user.UpdatedAtNotEq(&now),
			user.UpdatedAtGt(&now), user.UpdatedAtGtOrEq(&now),
			user.UpdatedAtLt(&now), user.UpdatedAtLtOrEq(&now),
		}

		pb := &predicate.Builder{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("updated_at").From("users")
		for _, p := range pb.Predicates() {
			require.NotNil(t, p.Val)
			_, ok := p.Val.(*time.Time)
			require.True(t, ok)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT updated_at FROM users WHERE updated_at = ? AND updated_at <> ? AND updated_at > ? AND updated_at >= ? AND updated_at < ? AND updated_at <= ?"
		assert.Equal(t, expect, got)
	})
}

func addPred(sb sq.SelectBuilder,
	p *predicate.Predicate) sq.SelectBuilder {
	switch p.Op {
	case predicate.Eq:
		return sb.Where(sq.Eq{
			p.Field: p.Val,
		})
	case predicate.NotEq:
		return sb.Where(sq.NotEq{
			p.Field: p.Val,
		})
	case predicate.Gt:
		return sb.Where(sq.Gt{
			p.Field: p.Val,
		})
	case predicate.GtOrEq:
		return sb.Where(sq.GtOrEq{
			p.Field: p.Val,
		})
	case predicate.Lt:
		return sb.Where(sq.Lt{
			p.Field: p.Val,
		})
	case predicate.LtOrEq:
		return sb.Where(sq.LtOrEq{
			p.Field: p.Val,
		})
	}

	return sb
}
