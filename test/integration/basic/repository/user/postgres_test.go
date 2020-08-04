package user_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
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
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		group_res VARCHAR(20) NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	require.NoError(t, err)

	repo := user.NewPostgreSQLRepository(db).Debug(os.Stderr)
	randomAge := func() int {
		return rand.Intn(30-18) + 18
	}

	ctx := context.Background()
	t.Run("Create", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			now := time.Now()
			for i := 1; i <= 50; i++ {
				group := "human"
				if i%2 == 0 {
					group = "charr"
				} else if i%3 == 0 {
					group = "norn"
				} else if i%4 == 0 {
					group = "sylvari"
				}

				email := fmt.Sprintf("%s_%d@gg.io", group, i)
				name := fmt.Sprintf("%s_%d", group, i)
				age := randomAge()

				cr := user.NewCreator().
					Email(&email).Name(&name).
					Age(age).GroupRes(group).
					UpdatedAt(&now)

				id, err := repo.Create(ctx, cr)
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
			crs := []*user.Creator{}
			for i := 51; i <= 100; i++ {
				group := "human"
				if i%2 == 0 {
					group = "charr"
				} else if i%3 == 0 {
					group = "norn"
				} else if i%4 == 0 {
					group = "sylvari"
				}

				email := fmt.Sprintf("%s_%d_mm@gg.io", group, i)
				name := fmt.Sprintf("%s_%d_mm", group, i)
				age := randomAge()
				now := time.Now()

				cr := user.NewCreator().
					Email(&email).Name(&name).
					Age(age).UpdatedAt(&now)
				crs = append(crs, cr)
			}

			err = repo.CreateMany(ctx, crs...)
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
			assert.Len(t, users, 100)
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

			users, err = repo.Query(ctx, user.NewQueryer().
				Where(
					user.AgeEq(18), user.AgeNotEq(30),
					user.AgeGt(17), user.AgeGtOrEq(18),
					user.AgeLt(30), user.AgeLtOrEq(19),
				),
			)
			assert.NoError(t, err)
			assert.NotZero(t, len(users))

			users, err = repo.Query(ctx, user.NewQueryer().
				Where(
					user.GroupEq("norn"), user.GroupNotEq("human"),
					user.GroupGt("n"), user.GroupGtOrEq("norn"),
					user.GroupLt("nornn"), user.GroupLtOrEq("norn"),
				),
			)
			assert.NoError(t, err)
			assert.NotZero(t, len(users))

			// with sort
			// get last user
			users, err = repo.Query(ctx, user.NewQueryer().
				Sort(
					user.Desc(user.ColumnID),
					user.Asc(user.ColumnCreatedAt),
				),
			)
			assert.NoError(t, err)
			assert.Equal(t, "charr_100_mm", *users[0].Name)

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
				Where(user.IDEq("9999")))
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

	t.Run("Aggregate", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			type aggt struct {
				AvgAge   float64
				MinAge   float64
				MaxAge   float64
				CountAge float64
				SumAge   float64
			}

			agg := []aggt{}

			a := user.NewAggregator(&agg).
				Aggregate(
					user.Avg(user.ColumnAge),
					user.Min(user.ColumnAge),
					user.Max(user.ColumnAge),
					user.Count(user.ColumnAge),
					user.Sum(user.ColumnAge),
				).
				Where(user.AgeGt(18)).
				Group(user.ColumnGroup).
				Sort(user.Asc(user.ColumnGroup))

			err := repo.Aggregate(ctx, a)
			require.NoError(t, err)
			assert.Len(t, agg, 4)

			for _, ag := range agg {
				assert.NotZero(t, ag.AvgAge)
				assert.NotZero(t, ag.MinAge)
				assert.NotZero(t, ag.MaxAge)
			}
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
			rowsAffected, err := repo.Update(ctx,
				user.NewUpdater().
					Email(&email).
					Name(&name).
					UpdatedAt(&now).
					Where(preds...),
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
			assert.Equal(t, int64(99), rowsAffected)
		})

		t.Run("Error", func(t *testing.T) {
			cctx, cancel := context.WithCancel(ctx)
			cancel()
			_, err = repo.Delete(cctx, user.NewDeleter())
			assert.Error(t, err)
		})
	})
}
