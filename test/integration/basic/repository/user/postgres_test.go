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
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	example "github.com/sf9v/nero/example"
	userr "github.com/sf9v/nero/test/integration/basic/repository/user"
	user "github.com/sf9v/nero/test/integration/basic/user"
)

func TestPostgreSQLRepository(t *testing.T) {
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
		Tag:        "12-alpine",
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
		id bigint GENERATED always AS IDENTITY PRIMARY KEY,
		uid varchar(27) NOT NULL UNIQUE,
		email VARCHAR(255) UNIQUE NOT NULL,
		"name" VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		"group" VARCHAR(20) NOT NULL,
		kv jsonb NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	require.NoError(t, err)

	repo := userr.NewPostgreSQLRepository(db).Debug(os.Stderr)
	newRepoTestRunner(repo)(t)
}

func newRepoTestRunner(repo userr.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()

		uids := []ksuid.KSUID{}
		kv := example.Map{"asdf": "jklm"}
		t.Run("Create", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				for i := 1; i <= 50; i++ {
					group := user.Human
					if i%2 == 0 {
						group = user.Charr
					} else if i%3 == 0 {
						group = user.Norn
					} else if i%4 == 0 {
						group = user.Sylvari
					}

					email := fmt.Sprintf("%s_%d@gg.io", group, i)
					name := fmt.Sprintf("%s_%d", group, i)
					age := randAge()

					uid := ksuid.New()
					uids = append(uids, uid)

					cr := userr.NewCreator().
						UID(uid).
						Email(email).
						Name(name).
						Age(age).
						Group(group).
						Kv(kv).
						UpdatedAt(&now)

					id, err := repo.Create(ctx, cr)
					assert.NoError(t, err)
					assert.NotZero(t, id)
				}
			})

			t.Run("Error", func(t *testing.T) {
				id, err := repo.Create(ctx, userr.NewCreator())
				assert.Error(t, err)
				assert.Zero(t, id)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Create(cctx, userr.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("CreateMany", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				crs := []*userr.Creator{}
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
					age := randAge()
					uid := ksuid.New()
					uids = append(uids, uid)

					cr := userr.NewCreator().
						UID(uid).
						Email(email).
						Name(name).
						Age(age).
						Kv(kv)
					crs = append(crs, cr)
				}

				err = repo.CreateMany(ctx, crs...)
				assert.NoError(t, err)

				err = repo.CreateMany(ctx, []*userr.Creator{}...)
				assert.NoError(t, err)
			})

			t.Run("Error", func(t *testing.T) {
				err := repo.CreateMany(ctx, userr.NewCreator())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				err = repo.CreateMany(cctx, userr.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("Query", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				users, err := repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.NotNil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
				}

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
				}

				// with predicates
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.IDEq("2"), userr.IDNotEq("1"),
						userr.IDGt("1"), userr.IDGtOrEq("2"),
						userr.IDLt("3"), userr.IDLtOrEq("2")))
				assert.NoError(t, err)
				assert.Len(t, users, 1)

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.AgeEq(18), userr.AgeNotEq(30),
						userr.AgeGt(17), userr.AgeGtOrEq(18),
						userr.AgeLt(30), userr.AgeLtOrEq(19),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.GroupEq(user.Norn), userr.GroupNotEq(user.Human),
						userr.GroupGt("n"), userr.GroupGtOrEq(user.Norn),
						userr.GroupLt("nornn"), userr.GroupLtOrEq(user.Norn),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.UIDEq(uids[0]), userr.UIDNotEq(uids[1]),
						userr.UIDGtOrEq(uids[0]), userr.UIDLtOrEq(uids[0]),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)

				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)

				// with sort
				// get last user
				users, err = repo.Query(ctx, userr.NewQueryer().
					Sort(
						userr.Desc(userr.ColumnID),
						userr.Asc(userr.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(users))
				assert.Equal(t, "charr_100_mm", users[0].Name)

				// with limit and offset
				users, err = repo.Query(ctx, userr.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Query(cctx, userr.NewQueryer())
				assert.Error(t, err)
			})
		})

		t.Run("QueryOne", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				usr, err := repo.QueryOne(ctx, userr.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)

				_, err = repo.QueryOne(ctx, userr.NewQueryer().Where(userr.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.QueryOne(cctx, userr.NewQueryer())
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
					Group    string
				}

				agg := []aggt{}

				a := userr.NewAggregator(&agg).
					Aggregate(
						userr.Avg(userr.ColumnAge),
						userr.Min(userr.ColumnAge),
						userr.Max(userr.ColumnAge),
						userr.Count(userr.ColumnAge),
						userr.Sum(userr.ColumnAge),
						userr.None(userr.ColumnGroup),
					).
					Where(userr.AgeGt(18), userr.GroupNotEq("")).
					Group(userr.ColumnGroup).
					Sort(userr.Asc(userr.ColumnGroup))

				err := repo.Aggregate(ctx, a)
				require.NoError(t, err)
				assert.Len(t, agg, 3)

				for _, ag := range agg {
					assert.NotZero(t, ag.AvgAge)
					assert.NotZero(t, ag.MinAge)
					assert.NotZero(t, ag.MaxAge)
					assert.NotZero(t, ag.CountAge)
					assert.NotZero(t, ag.SumAge)
					assert.NotEmpty(t, ag.Group)
				}
			})
		})

		t.Run("Update", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				preds := []userr.PredFunc{
					userr.IDEq("1"), userr.IDNotEq("2"),
					userr.IDGt("0"), userr.IDGtOrEq("1"),
					userr.IDLt("2"), userr.IDLtOrEq("1"),
				}

				email := "outcast@gg.io"
				name := "outcastn"
				age := 300
				rowsAffected, err := repo.Update(ctx,
					userr.NewUpdater().
						UID(ksuid.New()).
						Email(email).
						Name(name).
						Age(age).
						Group(user.Outcast).
						UpdatedAt(&now).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				usr, err := repo.QueryOne(ctx, userr.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)

				assert.Equal(t, email, usr.Email)
				assert.Equal(t, name, usr.Name)
				assert.Equal(t, age, usr.Age)
				assert.NotNil(t, usr.UpdatedAt)
			})

			t.Run("Error", func(t *testing.T) {
				_, err = repo.Update(ctx, userr.NewUpdater())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Update(cctx, userr.NewUpdater())
				assert.Error(t, err)
			})
		})

		t.Run("Delete", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				preds := []userr.PredFunc{
					userr.IDEq("1"), userr.IDNotEq("2"),
					userr.IDGt("0"), userr.IDGtOrEq("1"),
					userr.IDLt("2"), userr.IDLtOrEq("1"),
				}
				// delete one
				rowsAffected, err := repo.Delete(ctx,
					userr.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				usr, err := repo.QueryOne(ctx,
					userr.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)

				// delete all
				rowsAffected, err = repo.Delete(ctx, userr.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Delete(cctx, userr.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func randAge() int {
	return rand.Intn(30-18) + 18
}
