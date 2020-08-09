package user_test

import (
	"fmt"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/comparison"
	userr "github.com/sf9v/nero/test/integration/basic/repository/user"
	"github.com/sf9v/nero/test/integration/basic/user"
)

func TestPredicates(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		pfs := []userr.PredFunc{
			userr.IDEq("1"), userr.IDNotEq("1"),
			userr.IDGt("1"), userr.IDGtOrEq("1"),
			userr.IDLt("1"), userr.IDLtOrEq("1"),
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

	t.Run("UID", func(t *testing.T) {
		uid := ksuid.New()
		pfs := []userr.PredFunc{
			userr.UIDEq(uid), userr.UIDNotEq(uid),
			userr.UIDGt(uid), userr.UIDGtOrEq(uid),
			userr.UIDLt(uid), userr.UIDLtOrEq(uid),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("id").From("users")
		for _, p := range pb.All() {
			require.Equal(t, uid, p.Val)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := `SELECT id FROM users WHERE "uid" = ? AND "uid" <> ? AND "uid" > ? AND "uid" >= ? AND "uid" < ? AND "uid" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("Name", func(t *testing.T) {
		name := "sf9v"
		pfs := []userr.PredFunc{
			userr.NameEq(name), userr.NameNotEq(name),
			userr.NameGt(name), userr.NameGtOrEq(name),
			userr.NameLt(name), userr.NameLtOrEq(name),
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

	t.Run("Age", func(t *testing.T) {
		age := 10
		pfs := []userr.PredFunc{
			userr.AgeEq(age), userr.AgeNotEq(age),
			userr.AgeGt(age), userr.AgeGtOrEq(age),
			userr.AgeLt(age), userr.AgeLtOrEq(age),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("age").From("users")
		for _, p := range pb.All() {
			s, ok := p.Val.(int)
			assert.True(t, ok)
			assert.Equal(t, age, s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT age FROM users WHERE "age" = ? AND "age" <> ? AND "age" > ? AND "age" >= ? AND "age" < ? AND "age" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("Email", func(t *testing.T) {
		email := "sf9v@gg.io"
		pfs := []userr.PredFunc{
			userr.EmailEq(email), userr.EmailNotEq(email),
			userr.EmailGt(email), userr.EmailGtOrEq(email),
			userr.EmailLt(email), userr.EmailLtOrEq(email),
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

	t.Run("Group", func(t *testing.T) {
		group := user.Norn
		pfs := []userr.PredFunc{
			userr.GroupEq(group), userr.GroupNotEq(group),
			userr.GroupGt(group), userr.GroupGtOrEq(group),
			userr.GroupLt(group), userr.GroupLtOrEq(group),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("group").From("users")
		for _, p := range pb.All() {
			s, ok := p.Val.(user.Group)
			assert.True(t, ok)
			assert.Equal(t, group, s)
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT group FROM users WHERE "group" = ? AND "group" <> ? AND "group" > ? AND "group" >= ? AND "group" < ? AND "group" <= ?`
		assert.Equal(t, expect, got)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []userr.PredFunc{
			userr.CreatedAtEq(&now), userr.CreatedAtNotEq(&now),
			userr.CreatedAtGt(&now), userr.CreatedAtGtOrEq(&now),
			userr.CreatedAtLt(&now), userr.CreatedAtLtOrEq(&now),
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
		pfs := []userr.PredFunc{
			userr.UpdatedAtEq(&now), userr.UpdatedAtNotEq(&now),
			userr.UpdatedAtGt(&now), userr.UpdatedAtGtOrEq(&now),
			userr.UpdatedAtLt(&now), userr.UpdatedAtLtOrEq(&now),
			userr.UpdatedAtIsNull(), userr.UpdatedAtIsNotNull(),
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
