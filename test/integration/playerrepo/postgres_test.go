//go:build integration

package playerrepo_test

import (
	"bytes"
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero/test/integration/playerrepo"
)

func TestPostgresRepository(t *testing.T) {
	const dsn = "postgres://postgres:postgres@localhost:5432?sslmode=disable"

	// regular methods
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	defer db.Close()

	// create table
	require.NoError(t, createPgTable(db))

	// initialize a new repo
	repo := playerrepo.NewPostgresRepository(db).Debug()
	// WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunner(repo)(t)
	require.NoError(t, dropTable(db))

	// tx methods
	require.NoError(t, createPgTable(db))
	repo = playerrepo.NewPostgresRepository(db).Debug().
		WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunnerTx(repo)(t)
	require.NoError(t, dropTable(db))
}

func createPgTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE players (
		id bigint GENERATED always AS IDENTITY PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		"name" VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		"race" VARCHAR(20) NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table players`)
	return err
}
