package repository_test

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nero "github.com/sf9v/nero"
	example "github.com/sf9v/nero/example"
	"github.com/sf9v/nero/test/integration/repository"
	user "github.com/sf9v/nero/test/integration/user"
)

func TestPostgreSQLRepository(t *testing.T) {
	const dsn = "postgres://postgres:postgres@localhost:5432?sslmode=disable"

	// regular methods
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	require.NoError(t, createTable(db))

	buf := &bytes.Buffer{}
	logger := log.New(buf, "", 0)
	repo := repository.NewPostgresRepository(db).Debug().WithLogger(logger)
	newRepoTestRunner(repo)(t)
	require.NoError(t, dropTable(db))

	// tx methods
	require.NoError(t, createTable(db))
	repo = repository.NewPostgresRepository(db).Debug().WithLogger(logger)
	newRepoTestRunnerTx(repo)(t)
	require.NoError(t, dropTable(db))
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
		tags varchar(64)[] NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table users`)
	return err
}

func newRepoTestRunner(repo repository.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()

		uids := []ksuid.KSUID{}
		kv := example.Map{"asdf": "ghjk", "qwert": "yuio", "zxcv": "bnml"}
		tags := []string{"one", "two", "three"}
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

					cr := repository.NewCreator().
						UID(uid).
						Email(email).
						Name(name).
						Age(age).
						Group(group).
						Kv(kv).
						Tags(tags).
						UpdatedAt(&now)

					id, err := repo.Create(ctx, cr)
					assert.NoError(t, err)
					assert.NotZero(t, id)
				}
			})

			t.Run("Error", func(t *testing.T) {
				id, err := repo.Create(ctx, repository.NewCreator())
				assert.Error(t, err)
				assert.Zero(t, id)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Create(cctx, repository.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("CreateMany", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				crs := []*repository.Creator{}
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

					cr := repository.NewCreator().
						UID(uid).
						Email(email).
						Name(name).
						Age(age).
						Kv(kv).
						Tags(tags)
					crs = append(crs, cr)
				}

				err = repo.CreateMany(ctx, crs...)
				assert.NoError(t, err)

				err = repo.CreateMany(ctx, []*repository.Creator{}...)
				assert.NoError(t, err)
			})

			t.Run("Error", func(t *testing.T) {
				err := repo.CreateMany(ctx, repository.NewCreator())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				err = repo.CreateMany(cctx, repository.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("Query", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				users, err := repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.NotNil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Tags, 3)
				}

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Tags, 3)
				}

				// with predicates
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.IDEq("2"), repository.IDNotEq("1"),
						repository.IDGt("1"), repository.IDGtOrEq("2"),
						repository.IDLt("3"), repository.IDLtOrEq("2"),
						repository.IDIn("2"), repository.IDNotIn("1"),
					),
				)
				assert.NoError(t, err)
				assert.Len(t, users, 1)

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.AgeEq(18), repository.AgeNotEq(30),
						repository.AgeGt(17), repository.AgeGtOrEq(18),
						repository.AgeLt(30), repository.AgeLtOrEq(19),
						repository.AgeIn(18), repository.AgeNotIn(30),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.GroupEq(user.Norn), repository.GroupNotEq(user.Human),
						repository.GroupGt("n"), repository.GroupGtOrEq(user.Norn),
						repository.GroupLt("nornn"), repository.GroupLtOrEq(user.Norn),
						repository.GroupIn(user.Norn), repository.GroupNotIn(user.Human),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.UIDEq(uids[0]), repository.UIDNotEq(uids[1]),
						repository.UIDGtOrEq(uids[0]), repository.UIDLtOrEq(uids[0]),
						repository.UIDIn(uids[0]), repository.UIDNotIn(uids[1]),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)

				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)

				// with sort
				// get last user
				users, err = repo.Query(ctx, repository.NewQueryer().
					Sort(
						repository.Desc(repository.ColumnID),
						repository.Asc(repository.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(users))
				assert.Equal(t, "charr_100_mm", users[0].Name)

				// with limit and offset
				users, err = repo.Query(ctx, repository.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Query(cctx, repository.NewQueryer())
				assert.Error(t, err)
			})
		})

		t.Run("QueryOne", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				usr, err := repo.QueryOne(ctx, repository.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)
				assert.Len(t, usr.Tags, 3)

				_, err = repo.QueryOne(ctx, repository.NewQueryer().Where(repository.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.QueryOne(cctx, repository.NewQueryer())
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

				a := repository.NewAggregator(&agg).
					Aggregate(
						repository.Avg(repository.ColumnAge),
						repository.Min(repository.ColumnAge),
						repository.Max(repository.ColumnAge),
						repository.Count(repository.ColumnAge),
						repository.Sum(repository.ColumnAge),
						repository.None(repository.ColumnGroup),
					).
					Where(repository.AgeGt(18), repository.GroupNotEq("")).
					Group(repository.ColumnGroup).
					Sort(repository.Asc(repository.ColumnGroup))

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

		newTags := []string{"five"}

		t.Run("Update", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				preds := []repository.PredFunc{
					repository.IDEq("1"), repository.IDNotEq("2"),
					repository.IDGt("0"), repository.IDGtOrEq("1"),
					repository.IDLt("2"), repository.IDLtOrEq("1"),
				}

				email := "outcast@gg.io"
				name := "outcastn"
				age := 300
				rowsAffected, err := repo.Update(ctx,
					repository.NewUpdater().
						UID(ksuid.New()).
						Email(email).
						Name(name).
						Age(age).
						Group(user.Outcast).
						Tags(newTags).
						UpdatedAt(&now).
						Kv(example.Map{"abc": "def"}).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				usr, err := repo.QueryOne(ctx, repository.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)

				assert.Equal(t, email, usr.Email)
				assert.Equal(t, name, usr.Name)
				assert.Equal(t, age, usr.Age)
				assert.NotNil(t, usr.UpdatedAt)
				assert.Len(t, usr.Tags, 1)
			})

			t.Run("Error", func(t *testing.T) {
				_, err = repo.Update(ctx, repository.NewUpdater())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Update(cctx, repository.NewUpdater())
				assert.Error(t, err)
			})
		})

		t.Run("Delete", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				preds := []repository.PredFunc{
					repository.IDEq("1"), repository.IDNotEq("2"),
					repository.IDGt("0"), repository.IDGtOrEq("1"),
					repository.IDLt("2"), repository.IDLtOrEq("1"),
				}
				// delete one
				rowsAffected, err := repo.Delete(ctx,
					repository.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				usr, err := repo.QueryOne(ctx,
					repository.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)

				// delete all
				rowsAffected, err = repo.Delete(ctx, repository.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Delete(cctx, repository.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func newRepoTestRunnerTx(repo repository.Repository) func(t *testing.T) {
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
		tags := []string{"one", "two", "three"}
		t.Run("CreateTx", func(t *testing.T) {
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
						repository.NewCreator().UID(uid).
							Email(email).Name(name).
							Age(age).Group(group).
							Kv(kv).
							Tags(tags).
							UpdatedAt(&now))
					assert.NoError(t, err)
					assert.NotZero(t, id)
					assert.NoError(t, tx.Commit())
				}
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				id, err := repo.CreateTx(ctx, tx, repository.NewCreator())
				assert.Error(t, err)
				assert.Zero(t, id)
				assert.Error(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.CreateTx(cctx, tx, repository.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("CreateManyTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				crs := []*repository.Creator{}
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

					cr := repository.NewCreator().
						UID(uid).
						Email(email).
						Name(name).
						Age(age).
						Kv(kv).
						Tags(tags)
					crs = append(crs, cr)
				}
				tx := newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, crs...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, []*repository.Creator{}...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				err := repo.CreateManyTx(ctx, tx, repository.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				err = repo.CreateManyTx(cctx, tx, repository.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("QueryTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				users, err := repo.QueryTx(ctx, tx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.NotNil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Tags, 3)
				}
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Tags, 3)
				}
				assert.NoError(t, tx.Commit())

				// with predicates
				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, repository.NewQueryer().
					Where(repository.IDEq("2"), repository.IDNotEq("1"),
						repository.IDGt("1"), repository.IDGtOrEq("2"),
						repository.IDLt("3"), repository.IDLtOrEq("2")))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.AgeEq(18), repository.AgeNotEq(30),
						repository.AgeGt(17), repository.AgeGtOrEq(18),
						repository.AgeLt(30), repository.AgeLtOrEq(19),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.GroupEq(user.Norn), repository.GroupNotEq(user.Human),
						repository.GroupGt("n"), repository.GroupGtOrEq(user.Norn),
						repository.GroupLt("nornn"), repository.GroupLtOrEq(user.Norn),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(
						repository.UIDEq(uids[0]), repository.UIDNotEq(uids[1]),
						repository.UIDGtOrEq(uids[0]), repository.UIDLtOrEq(uids[0]),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Where(repository.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				// with sort
				// get last user
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().
					Sort(
						repository.Desc(repository.ColumnID),
						repository.Asc(repository.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(users))
				assert.Equal(t, "charr_100_mm", users[0].Name)
				assert.NoError(t, tx.Commit())

				// with limit and offset
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, repository.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryTx(cctx, tx, repository.NewQueryer())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("QueryOneTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				usr, err := repo.QueryOneTx(ctx, tx, repository.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)
				assert.Len(t, usr.Tags, 3)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				_, err = repo.QueryOne(ctx, repository.NewQueryer().Where(repository.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryOneTx(cctx, tx, repository.NewQueryer())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("AggregateTx", func(t *testing.T) {
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

				a := repository.NewAggregator(&agg).
					Aggregate(
						repository.Avg(repository.ColumnAge),
						repository.Min(repository.ColumnAge),
						repository.Max(repository.ColumnAge),
						repository.Count(repository.ColumnAge),
						repository.Sum(repository.ColumnAge),
						repository.None(repository.ColumnGroup),
					).
					Where(repository.AgeGt(18), repository.GroupNotEq("")).
					Group(repository.ColumnGroup).
					Sort(repository.Asc(repository.ColumnGroup))

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

		newTags := []string{"five"}

		t.Run("UpdateTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				preds := []repository.PredFunc{
					repository.IDEq("1"), repository.IDNotEq("2"),
					repository.IDGt("0"), repository.IDGtOrEq("1"),
					repository.IDLt("2"), repository.IDLtOrEq("1"),
				}

				email := "outcast@gg.io"
				name := "outcastn"
				age := 300

				tx := newTx(ctx, t)
				rowsAffected, err := repo.UpdateTx(ctx, tx,
					repository.NewUpdater().
						UID(ksuid.New()).
						Email(email).
						Name(name).
						Age(age).
						Group(user.Outcast).
						Tags(newTags).
						UpdatedAt(&now).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx, repository.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				assert.Equal(t, email, usr.Email)
				assert.Equal(t, name, usr.Name)
				assert.Equal(t, age, usr.Age)
				assert.NotNil(t, usr.UpdatedAt)
				assert.Len(t, usr.Tags, 1)
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				_, err = repo.UpdateTx(ctx, tx, repository.NewUpdater())
				assert.Error(t, err)
				assert.NoError(t, tx.Commit())

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.UpdateTx(cctx, tx, repository.NewUpdater())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("DeleteTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				preds := []repository.PredFunc{
					repository.IDEq("1"), repository.IDNotEq("2"),
					repository.IDGt("0"), repository.IDGtOrEq("1"),
					repository.IDLt("2"), repository.IDLtOrEq("1"),
				}
				// delete one
				tx := newTx(ctx, t)
				rowsAffected, err := repo.DeleteTx(ctx, tx,
					repository.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx,
					repository.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)
				assert.NoError(t, tx.Commit())

				// delete all
				tx = newTx(ctx, t)
				rowsAffected, err = repo.Delete(ctx, repository.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(ctx, t)
				cancel()
				_, err = repo.DeleteTx(cctx, tx, repository.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func randAge() int {
	return rand.Intn(30-18) + 18
}
