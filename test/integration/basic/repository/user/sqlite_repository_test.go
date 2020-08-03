package user_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/test/integration/basic/repository/user"
)

func TestSQLiteRepository(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer db.Close()

	// create users table
	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	require.NoError(t, err)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	repo := user.NewSQLiteRepository(db).Debug(os.Stderr)
	ctx := context.Background()
	t.Run("Create", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			now := time.Now()
			for i := 1; i <= 10; i++ {
				email := fmt.Sprintf("user%d@gg.io", i)
				name := fmt.Sprintf("user%d", i)
				id, err := repo.Create(ctx, user.NewCreator().
					Email(&email).Name(&name).UpdatedAt(&now))
				assert.NoError(t, err)
				assert.NotZero(t, id)
			}
		})

		t.Run("Error", func(t *testing.T) {
			id, err := repo.Create(ctx, user.NewCreator())
			assert.Error(t, err)
			assert.Zero(t, id)

			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.Create(cctx, user.NewCreator())
			assert.Error(t, err)
		})
	})

	t.Run("CreateMany", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			cs := []*user.Creator{}
			for i := 11; i <= 20; i++ {
				email := fmt.Sprintf("user_m%d@gg.io", i)
				name := fmt.Sprintf("user_m%d", i)
				now := time.Now()
				cs = append(cs, user.NewCreator().
					Email(&email).Name(&name).UpdatedAt(&now))
			}

			err = repo.CreateMany(ctx, cs...)
			assert.NoError(t, err)

			err = repo.CreateMany(ctx, []*user.Creator{}...)
			assert.NoError(t, err)
		})

		t.Run("Error", func(t *testing.T) {
			err := repo.CreateMany(ctx, user.NewCreator())
			assert.Error(t, err)

			cctx, cancel := context.WithCancel(ctx)
			cancel()
			err = repo.CreateMany(cctx, user.NewCreator())
			assert.Error(t, err)
		})
	})

	t.Run("Query", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			// all users
			users, err := repo.Query(ctx,
				user.NewQueryer())
			assert.NoError(t, err)
			assert.Len(t, users, 20)
			for _, u := range users {
				require.NotNil(t, u.Email)
				require.NotNil(t, u.Name)
				require.NotNil(t, u.UpdatedAt)
				require.NotNil(t, u.CreatedAt)
			}

			// with predicates
			users, err = repo.Query(ctx, user.NewQueryer().
				Where(user.IDEq("2"), user.IDNotEq("1"),
					user.IDGt("1"), user.IDGtOrEq("2"),
					user.IDLt("3"), user.IDLtOrEq("2")))
			assert.NoError(t, err)
			assert.Len(t, users, 1)

			// with sort
			// get last user
			users, err = repo.Query(ctx, user.NewQueryer().
				Sort(user.IDDesc(), user.CreatedAtAsc()),
			)
			assert.NoError(t, err)
			assert.Equal(t, "user_m20", *users[0].Name)

			// with limit and offset
			users, err = repo.Query(ctx, user.NewQueryer().
				Limit(1).Offset(1))
			assert.NoError(t, err)
			assert.Len(t, users, 1)
		})

		t.Run("Error", func(t *testing.T) {
			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.Query(cctx, user.NewQueryer())
			assert.Error(t, err)
		})
	})

	t.Run("QueryOne", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			usr, err := repo.QueryOne(ctx, user.NewQueryer())
			assert.NoError(t, err)
			assert.NotNil(t, usr)
			assert.Equal(t, "1", usr.ID)

			_, err = repo.QueryOne(ctx, user.NewQueryer().
				Where(user.IDEq("100")))
			assert.Error(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
		})

		t.Run("Error", func(t *testing.T) {
			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.QueryOne(cctx, user.NewQueryer())
			assert.Error(t, err)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			now := time.Now()
			preds := []user.PredFunc{
				user.IDEq("1"), user.IDNotEq("2"),
				user.IDGt("0"), user.IDGtOrEq("1"),
				user.IDLt("2"), user.IDLtOrEq("1"),
			}

			email := "sf9v@gg.io"
			name := "sf9v"
			rowsAffected, err := repo.Update(ctx, user.NewUpdater().
				Email(&email).Name(&name).
				UpdatedAt(&now).Where(preds...),
			)
			assert.NoError(t, err)
			assert.Equal(t, int64(1), rowsAffected)

			users, err := repo.Query(ctx, user.NewQueryer().
				Where(preds...))
			assert.NoError(t, err)
			assert.Len(t, users, 1)

			u := users[0]
			assert.Equal(t, "sf9v@gg.io", *u.Email)
			assert.Equal(t, "sf9v", *u.Name)
			assert.NotNil(t, u.UpdatedAt)
		})

		t.Run("Error", func(t *testing.T) {
			_, err = repo.Update(ctx, user.NewUpdater())
			assert.Error(t, err)

			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.Update(cctx, user.NewUpdater())
			assert.Error(t, err)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			preds := []user.PredFunc{
				user.IDEq("1"), user.IDNotEq("2"),
				user.IDGt("0"), user.IDGtOrEq("1"),
				user.IDLt("2"), user.IDLtOrEq("1"),
			}
			rowsAffected, err := repo.Delete(ctx,
				user.NewDeleter().Where(preds...))
			assert.NoError(t, err)
			assert.Equal(t, int64(1), rowsAffected)

			usr, err := repo.QueryOne(ctx,
				user.NewQueryer().Where(preds...))
			assert.Error(t, err, sql.ErrNoRows)
			assert.Nil(t, usr)

			// delete all
			rowsAffected, err = repo.Delete(ctx, user.NewDeleter())
			assert.NoError(t, err)
			assert.Equal(t, int64(19), rowsAffected)
		})

		t.Run("Error", func(t *testing.T) {
			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.Delete(cctx, user.NewDeleter())
			assert.Error(t, err)
		})
	})
}
