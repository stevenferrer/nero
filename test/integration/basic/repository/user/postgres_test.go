package user_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/test/integration/basic/repository/user"
)

func TestPGRepository(t *testing.T) {
	var (
		usr    = "postgres"
		pwd    = "postgres"
		dbName = "postgres"
		port   = "5432"
		dsnFmt = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	)

	// postgres setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + usr,
			"POSTGRES_PASSWORD=" + pwd,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {{HostIP: "localhost", HostPort: port}},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	defer func() {
		require.NoError(t, pool.Purge(resource))
	}()

	var db *sql.DB
	dsn := fmt.Sprintf(dsnFmt, usr, pwd, port, dbName)
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE users(
		id bigserial PRIMARY KEY,
		uid UUID NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		group_res VARCHAR(20) NOT NULL,
		kv jsonb NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	require.NoError(t, err)

	repo := user.NewPostgreSQLRepository(db).Debug(os.Stderr)
	newRepoTestRunner(repo)(t)
}
