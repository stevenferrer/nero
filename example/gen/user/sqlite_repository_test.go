package user

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSQLiteRepo(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer db.Close()

	err = createTable(db)
	require.NoError(t, err)

	sqliteRepo := NewSQLiteRepository(db)

	ctx := context.Background()
	id, err := sqliteRepo.Create(ctx,
		NewCreator().
			Email("sf9v@gg.io").
			Name("sf9v"),
	)
	require.NoError(t, err)
	assert.NotZero(t, id)

	users, err := sqliteRepo.Query(ctx, NewQueryer().Where(IDEq(id)))
	require.NoError(t, err)
	assert.Len(t, users, 1)

	_, err = sqliteRepo.Update(ctx, NewUpdater().Name("steve").Where(IDEq(id)))
	require.NoError(t, err)

	users, err = sqliteRepo.Query(ctx, NewQueryer().Where(IDEq(id)))
	require.NoError(t, err)
	assert.Len(t, users, 1)

	for _, user := range users {
		if user.ID == id {
			assert.Equal(t, "steve", user.Name)
		}
	}
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	updated_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`)
	return err
}
