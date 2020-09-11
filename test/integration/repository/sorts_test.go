package repository_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/sort"
	"github.com/sf9v/nero/test/integration/repository"
)

func TestSorts(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		sfs := []repository.SortFunc{
			repository.Asc(repository.ColumnID),
			repository.Desc(repository.ColumnID),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("id").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT id FROM users ORDER BY id ASC, id DESC"
		assert.Equal(t, expect, got)
	})

	t.Run("UID", func(t *testing.T) {
		sfs := []repository.SortFunc{
			repository.Asc(repository.ColumnUID),
			repository.Desc(repository.ColumnUID),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("uid").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT uid FROM users ORDER BY uid ASC, uid DESC"
		assert.Equal(t, expect, got)
	})

	t.Run("Email", func(t *testing.T) {
		sfs := []repository.SortFunc{
			repository.Asc(repository.ColumnEmail),
			repository.Desc(repository.ColumnEmail),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("email").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT email FROM users ORDER BY email ASC, email DESC"
		assert.Equal(t, expect, got)
	})

	t.Run("Name", func(t *testing.T) {
		sfs := []repository.SortFunc{
			repository.Asc(repository.ColumnName),
			repository.Desc(repository.ColumnName),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("name").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT name FROM users ORDER BY name ASC, name DESC"
		assert.Equal(t, expect, got)
	})

	t.Run("UpdatedAt", func(t *testing.T) {
		sfs := []repository.SortFunc{
			repository.Asc(repository.ColumnUpdatedAt),
			repository.Desc(repository.ColumnUpdatedAt),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("updated_at").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT updated_at FROM users ORDER BY updated_at ASC, updated_at DESC"
		assert.Equal(t, expect, got)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		sfs := []repository.SortFunc{
			// user.CreatedAtAsc(), user.CreatedAtDesc(),
			repository.Asc(repository.ColumnCreatedAt),
			repository.Desc(repository.ColumnCreatedAt),
		}

		sb := &sort.Sorts{}
		for _, sf := range sfs {
			sf(sb)
		}

		qb := sq.Select("created_at").From("users")
		for _, s := range sb.All() {
			qb = addSorts(qb, s)
		}

		got, _, err := qb.ToSql()
		require.NoError(t, err)
		expect := "SELECT created_at FROM users ORDER BY created_at ASC, created_at DESC"
		assert.Equal(t, expect, got)
	})
}

func addSorts(sb sq.SelectBuilder, s *sort.Sort) sq.SelectBuilder {
	switch s.Direction {
	case sort.Asc:
		return sb.OrderBy(fmt.Sprintf("%s ASC", s.Col))
	case sort.Desc:
		return sb.OrderBy(fmt.Sprintf("%s DESC", s.Col))
	}

	return sb
}
