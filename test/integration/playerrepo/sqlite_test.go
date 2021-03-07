// +build integration

package playerrepo_test

import (
	"bytes"
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sf9v/nero/test/integration/playerrepo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSQLiteRepository(t *testing.T) {
	t.Parallel()

	const dsn = "file:test.db?mode=memory&cache=shared"
	db, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	defer db.Close()

	// create table
	err = createSqliteTable(db)
	assert.NoError(t, err)

	// initialize a new repo
	repo := playerrepo.NewSQLiteRepository(db).Debug().
		WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunner(repo)(t)
	// cleanup
	require.NoError(t, dropTable(db))

	// Tx methods
	// re-create table
	err = createSqliteTable(db)
	assert.NoError(t, err)

	// initialize a new repo
	repo = playerrepo.NewSQLiteRepository(db).Debug().
		WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunnerTx(repo)(t)
	require.NoError(t, dropTable(db))
}

func createSqliteTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE players (
		id INTEGER PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		"name" TEXT NOT NULL,
		age INTEGER NOT NULL,
		race TEXT NOT NULL,
		updated_at DATETIME NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	return err
}
