package repository_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/comparison"
	"github.com/sf9v/nero/test/integration/repository"
	"github.com/sf9v/nero/test/integration/user"
)

func TestPredicates(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		pfs := []repository.PredFunc{
			repository.IDEq("1"), repository.IDNotEq("1"),
			repository.IDGt("1"), repository.IDGtOrEq("1"),
			repository.IDLt("1"), repository.IDLtOrEq("1"),
			repository.IDIn("1"), repository.IDNotIn("1"),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("id").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := `SELECT id FROM users WHERE "id" = ? AND "id" <> ? AND "id" > ? AND "id" >= ? AND "id" < ? AND "id" <= ? AND "id" IN (?) AND "id" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("UID", func(t *testing.T) {
		uid := ksuid.New()
		pfs := []repository.PredFunc{
			repository.UIDEq(uid), repository.UIDNotEq(uid),
			repository.UIDGt(uid), repository.UIDGtOrEq(uid),
			repository.UIDLt(uid), repository.UIDLtOrEq(uid),
			repository.UIDIn(uid), repository.UIDNotIn(uid),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("id").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := `SELECT id FROM users WHERE "uid" = ? AND "uid" <> ? AND "uid" > ? AND "uid" >= ? AND "uid" < ? AND "uid" <= ? AND "uid" IN (?) AND "uid" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("Name", func(t *testing.T) {
		name := "sf9v"
		pfs := []repository.PredFunc{
			repository.NameEq(name), repository.NameNotEq(name),
			repository.NameGt(name), repository.NameGtOrEq(name),
			repository.NameLt(name), repository.NameLtOrEq(name),
			repository.NameIn(name), repository.NameNotIn(name),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("name").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT name FROM users WHERE "name" = ? AND "name" <> ? AND "name" > ? AND "name" >= ? AND "name" < ? AND "name" <= ? AND "name" IN (?) AND "name" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("Age", func(t *testing.T) {
		age := 10
		pfs := []repository.PredFunc{
			repository.AgeEq(age), repository.AgeNotEq(age),
			repository.AgeGt(age), repository.AgeGtOrEq(age),
			repository.AgeLt(age), repository.AgeLtOrEq(age),
			repository.AgeIn(age), repository.AgeNotIn(age),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("age").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT age FROM users WHERE "age" = ? AND "age" <> ? AND "age" > ? AND "age" >= ? AND "age" < ? AND "age" <= ? AND "age" IN (?) AND "age" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("Email", func(t *testing.T) {
		email := "sf9v@gg.io"
		pfs := []repository.PredFunc{
			repository.EmailEq(email), repository.EmailNotEq(email),
			repository.EmailGt(email), repository.EmailGtOrEq(email),
			repository.EmailLt(email), repository.EmailLtOrEq(email),
			repository.EmailIn(email), repository.EmailNotIn(email),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("email").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT email FROM users WHERE "email" = ? AND "email" <> ? AND "email" > ? AND "email" >= ? AND "email" < ? AND "email" <= ? AND "email" IN (?) AND "email" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("Group", func(t *testing.T) {
		group := user.GroupNorn
		pfs := []repository.PredFunc{
			repository.GroupEq(group), repository.GroupNotEq(group),
			repository.GroupGt(group), repository.GroupGtOrEq(group),
			repository.GroupLt(group), repository.GroupLtOrEq(group),
			repository.GroupIn(group), repository.GroupNotIn(group),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("group").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT group FROM users WHERE "group" = ? AND "group" <> ? AND "group" > ? AND "group" >= ? AND "group" < ? AND "group" <= ? AND "group" IN (?) AND "group" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []repository.PredFunc{
			repository.CreatedAtEq(&now), repository.CreatedAtNotEq(&now),
			repository.CreatedAtGt(&now), repository.CreatedAtGtOrEq(&now),
			repository.CreatedAtLt(&now), repository.CreatedAtLtOrEq(&now),
			repository.CreatedAtIn(&now), repository.CreatedAtNotIn(&now),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("created_at").From("users")
		for _, p := range pb.All() {
			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT created_at FROM users WHERE "created_at" = ? AND "created_at" <> ? AND "created_at" > ? AND "created_at" >= ? AND "created_at" < ? AND "created_at" <= ? AND "created_at" IN (?) AND "created_at" NOT IN (?)`
		assert.Equal(t, expect, got)
	})

	t.Run("UpdatedAt", func(t *testing.T) {
		now := time.Now()
		pfs := []repository.PredFunc{
			repository.UpdatedAtEq(&now), repository.UpdatedAtNotEq(&now),
			repository.UpdatedAtGt(&now), repository.UpdatedAtGtOrEq(&now),
			repository.UpdatedAtLt(&now), repository.UpdatedAtLtOrEq(&now),
			repository.UpdatedAtIsNull(), repository.UpdatedAtIsNotNull(),
			repository.UpdatedAtIn(&now), repository.UpdatedAtNotIn(&now),
		}

		pb := &comparison.Predicates{}
		for _, pf := range pfs {
			pf(pb)
		}

		qb := sq.Select("updated_at").From("users")
		for _, p := range pb.All() {
			if p.Op == comparison.IsNull ||
				p.Op == comparison.IsNotNull {
				assert.Nil(t, p.Arg)
			}

			qb = addPred(qb, p)
		}

		got, _, err := qb.ToSql()
		assert.NoError(t, err)
		expect := `SELECT updated_at FROM users WHERE "updated_at" = ? AND "updated_at" <> ? AND "updated_at" > ? AND "updated_at" >= ? AND "updated_at" < ? AND "updated_at" <= ? AND "updated_at" IS NULL AND "updated_at" IS NOT NULL AND "updated_at" IN (?) AND "updated_at" NOT IN (?)`
		assert.Equal(t, expect, got)
	})
}

func addPred(sb sq.SelectBuilder,
	p *comparison.Predicate) sq.SelectBuilder {
	switch p.Op {
	case comparison.Eq:
		return sb.Where(fmt.Sprintf("%q = ?", p.Col), p.Arg)
	case comparison.NotEq:
		return sb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Arg)
	case comparison.Gt:
		return sb.Where(fmt.Sprintf("%q > ?", p.Col), p.Arg)
	case comparison.GtOrEq:
		return sb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Arg)
	case comparison.Lt:
		return sb.Where(fmt.Sprintf("%q < ?", p.Col), p.Arg)
	case comparison.LtOrEq:
		return sb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Arg)
	case comparison.IsNull:
		return sb.Where(fmt.Sprintf("%q IS NULL", p.Col))
	case comparison.IsNotNull:
		return sb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
	case comparison.In, comparison.NotIn:
		fmtStr := "%q IN (%s)"
		if p.Op == comparison.NotIn {
			fmtStr = "%q NOT IN (%s)"
		}
		args := p.Arg.([]interface{})
		qms := []string{}
		for range args {
			qms = append(qms, "?")
		}
		plchldr := strings.Join(qms, ",")
		return sb.Where(fmt.Sprintf(fmtStr, p.Col, plchldr), args...)

	}

	return sb
}
