package user_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/test/integration/gen/user"
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
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
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
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(50) NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	require.NoError(t, err)

	repo := user.NewPGRepository(db)

	ctx := context.Background()
	t.Run("Create", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			for i := 1; i <= 10; i++ {
				email := fmt.Sprintf("user%d@gg.io", i)
				name := fmt.Sprintf("user%d", i)
				id, err := repo.Create(ctx, user.NewCreator().
					Email(&email).Name(&name))
				assert.NoError(t, err)
				assert.NotZero(t, id)
			}
		})

		t.Run("Error", func(t *testing.T) {
			id, err := repo.Create(ctx, user.NewCreator())
			assert.Error(t, err)
			assert.Zero(t, id)
		})
	})

	t.Run("CreateM", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			cs := []*user.Creator{}
			for i := 1; i <= 10; i++ {
				email := fmt.Sprintf("user_m%d@gg.io", i)
				name := fmt.Sprintf("user_%d", i)
				now := time.Now()
				cs = append(cs, user.NewCreator().
					Email(&email).Name(&name).UpdatedAt(&now))
			}

			err = repo.CreateM(ctx, cs...)
			assert.NoError(t, err)

			err = repo.CreateM(ctx, []*user.Creator{}...)
			assert.NoError(t, err)
		})

		t.Run("Error", func(t *testing.T) {
			err := repo.CreateM(ctx, user.NewCreator())
			assert.Error(t, err)
		})
	})

	t.Run("Query", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			// all users
			users, err := repo.Query(ctx, user.NewQueryer())
			assert.NoError(t, err)
			assert.Len(t, users, 20)

			// with predicates
			users, err = repo.Query(ctx, user.NewQueryer().
				Where(user.IDEq("2"), user.IDNotEq("1"),
					user.IDGt("1"), user.IDGtOrEq("2"),
					user.IDLt("3"), user.IDLtOrEq("2")))
			assert.NoError(t, err)
			assert.Len(t, users, 1)

			// with limit and offset
			users, err = repo.Query(ctx, user.NewQueryer().
				Limit(1).Offset(1))
			assert.NoError(t, err)
			assert.Len(t, users, 1)
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

			users, err := repo.Query(ctx,
				user.NewQueryer().Where(preds...))
			assert.NoError(t, err)
			assert.Len(t, users, 0)
		})
	})
}
