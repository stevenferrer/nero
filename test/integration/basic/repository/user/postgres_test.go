package user_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nero "github.com/sf9v/nero"
	example "github.com/sf9v/nero/example"
	userr "github.com/sf9v/nero/test/integration/basic/repository/user"
	user "github.com/sf9v/nero/test/integration/basic/user"
)

func TestPostgreSQLRepository(t *testing.T) {
	var dsnFmt = "postgres://postgres:postgres@localhost:%s?sslmode=disable"

	// postgres setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12-alpine",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=postgres",
		},
	}

	t.Run("Non-Tx", func(t *testing.T) {
		resource, err := pool.RunWithOptions(&opts)
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}
		defer func() {
			require.NoError(t, pool.Purge(resource))
		}()

		var db *sql.DB
		dsn := fmt.Sprintf(dsnFmt, resource.GetPort("5432/tcp"))
		if err = pool.Retry(func() error {
			db, err = sql.Open("postgres", dsn)
			if err != nil {
				return err
			}
			return db.Ping()
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		require.NoError(t, createTable(db))
		repo := userr.NewPostgreSQLRepository(db)
		newRepoTestRunner(repo)(t)
		require.NoError(t, dropTable(db))
	})

	t.Run("Tx", func(t *testing.T) {
		resource, err := pool.RunWithOptions(&opts)
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}
		defer func() {
			require.NoError(t, pool.Purge(resource))
		}()

		var db *sql.DB
		dsn := fmt.Sprintf(dsnFmt, resource.GetPort("5432/tcp"))
		if err = pool.Retry(func() error {
			db, err = sql.Open("postgres", dsn)
			if err != nil {
				return err
			}
			return db.Ping()
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		require.NoError(t, createTable(db))
		repo := userr.NewPostgreSQLRepository(db)
		newRepoTestRunnerTx(repo)(t)
		require.NoError(t, dropTable(db))
	})
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE users(
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
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table users`)
	return err
}

func newRepoTestRunner(repo userr.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()

		uids := []ksuid.KSUID{}
		kv := example.Map{"asdf": "ghjk", "qwert": "yuio", "zxcv": "bnml"}
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
						Uid(ksuid.New()).
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

func newRepoTestRunnerTx(repo userr.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()
		newTx := func(ctx context.Context, t *testing.T) nero.Tx {
			tx, err := repo.Tx(ctx)
			assert.NoError(t, err)
			return tx
		}
		uids := []ksuid.KSUID{}
		kv := example.Map{"asdf": "ghjk", "qwert": "yuio", "zxcv": "bnml"}
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

					tx := newTx(ctx, t)
					id, err := repo.CreateTx(ctx, tx,
						userr.NewCreator().UID(uid).
							Email(email).Name(name).
							Age(age).Group(group).
							Kv(kv).UpdatedAt(&now))
					assert.NoError(t, err)
					assert.NotZero(t, id)
					assert.NoError(t, tx.Commit())
				}
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				id, err := repo.CreateTx(ctx, tx, userr.NewCreator())
				assert.Error(t, err)
				assert.Zero(t, id)
				assert.Error(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.CreateTx(cctx, tx, userr.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
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
				tx := newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, crs...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, []*userr.Creator{}...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				err := repo.CreateManyTx(ctx, tx, userr.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				err = repo.CreateManyTx(cctx, tx, userr.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("Query", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				users, err := repo.QueryTx(ctx, tx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.NotNil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
				}
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
				}
				assert.NoError(t, tx.Commit())

				// with predicates
				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, userr.NewQueryer().
					Where(userr.IDEq("2"), userr.IDNotEq("1"),
						userr.IDGt("1"), userr.IDGtOrEq("2"),
						userr.IDLt("3"), userr.IDLtOrEq("2")))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.AgeEq(18), userr.AgeNotEq(30),
						userr.AgeGt(17), userr.AgeGtOrEq(18),
						userr.AgeLt(30), userr.AgeLtOrEq(19),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.GroupEq(user.Norn), userr.GroupNotEq(user.Human),
						userr.GroupGt("n"), userr.GroupGtOrEq(user.Norn),
						userr.GroupLt("nornn"), userr.GroupLtOrEq(user.Norn),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(
						userr.UIDEq(uids[0]), userr.UIDNotEq(uids[1]),
						userr.UIDGtOrEq(uids[0]), userr.UIDLtOrEq(uids[0]),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Where(userr.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				// with sort
				// get last user
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().
					Sort(
						userr.Desc(userr.ColumnID),
						userr.Asc(userr.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(users))
				assert.Equal(t, "charr_100_mm", users[0].Name)
				assert.NoError(t, tx.Commit())

				// with limit and offset
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, userr.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryTx(cctx, tx, userr.NewQueryer())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("QueryOne", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				usr, err := repo.QueryOneTx(ctx, tx, userr.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				_, err = repo.QueryOne(ctx, userr.NewQueryer().Where(userr.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryOneTx(cctx, tx, userr.NewQueryer())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
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

				tx := newTx(ctx, t)
				err := repo.AggregateTx(ctx, tx, a)
				require.NoError(t, err)
				assert.Len(t, agg, 3)
				assert.NoError(t, tx.Commit())

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

				tx := newTx(ctx, t)
				rowsAffected, err := repo.UpdateTx(ctx, tx,
					userr.NewUpdater().
						Uid(ksuid.New()).
						Email(email).
						Name(name).
						Age(age).
						Group(user.Outcast).
						UpdatedAt(&now).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx, userr.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				assert.Equal(t, email, usr.Email)
				assert.Equal(t, name, usr.Name)
				assert.Equal(t, age, usr.Age)
				assert.NotNil(t, usr.UpdatedAt)
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				_, err = repo.UpdateTx(ctx, tx, userr.NewUpdater())
				assert.Error(t, err)
				assert.NoError(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.UpdateTx(cctx, tx, userr.NewUpdater())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
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
				tx := newTx(ctx, t)
				rowsAffected, err := repo.DeleteTx(ctx, tx,
					userr.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx,
					userr.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)
				assert.NoError(t, tx.Commit())

				// delete all
				tx = newTx(ctx, t)
				rowsAffected, err = repo.Delete(ctx, userr.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(ctx, t)
				cancel()
				_, err = repo.DeleteTx(cctx, tx, userr.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func randAge() int {
	return rand.Intn(30-18) + 18
}
