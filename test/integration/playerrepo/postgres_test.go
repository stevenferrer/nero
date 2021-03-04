package playerrepo_test

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/comparison"
	"github.com/sf9v/nero/test/integration/player"
	"github.com/sf9v/nero/test/integration/playerrepo"
)

func TestPostgresRepository(t *testing.T) {
	const dsn = "postgres://postgres:postgres@localhost:5432?sslmode=disable"

	// regular methods
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())

	// create table
	require.NoError(t, createTable(db))
	repo := playerrepo.NewPostgresRepository(db).Debug().
		WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunner(repo)(t)
	require.NoError(t, dropTable(db))

	// tx methods
	require.NoError(t, createTable(db))
	repo = playerrepo.NewPostgresRepository(db).Debug().
		WithLogger(log.New(&bytes.Buffer{}, "", 0))
	newRepoTestRunnerTx(repo)(t)
	require.NoError(t, dropTable(db))
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE players(
		id bigint GENERATED always AS IDENTITY PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		"name" VARCHAR(50) NOT NULL,
		age INTEGER NOT NULL,
		"race" VARCHAR(20) NOT NULL,
		interests varchar(64)[] NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table players`)
	return err
}

func newRepoTestRunner(repo playerrepo.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()

		interests := []string{"pizza", "dogs", "cats"}

		t.Run("Create", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				for i := 1; i <= 50; i++ {
					race := player.RaceHuman
					if i%2 == 0 {
						race = player.RaceCharr
					} else if i%3 == 0 {
						race = player.RaceNorn
					} else if i%4 == 0 {
						race = player.RaceSylvari
					}

					email := fmt.Sprintf("%s_%d@gg.io", race, i)
					name := fmt.Sprintf("%s_%d", race, i)
					age := randAge()

					id, err := repo.Create(ctx, playerrepo.NewCreator().
						Email(email).Name(name).Age(age).Race(race).
						Interests(interests))
					assert.NoError(t, err)
					assert.NotZero(t, id)
				}
			})

			t.Run("Error", func(t *testing.T) {
				_, err := repo.Create(ctx, playerrepo.NewCreator())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Create(cctx, playerrepo.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("CreateMany", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				crs := []*playerrepo.Creator{}
				for i := 51; i <= 100; i++ {
					race := player.RaceHuman
					if i%2 == 0 {
						race = player.RaceCharr
					} else if i%3 == 0 {
						race = player.RaceNorn
					} else if i%4 == 0 {
						race = player.RaceSylvari
					}

					email := fmt.Sprintf("%s_%d_mm@gg.io", race, i)
					name := fmt.Sprintf("%s_%d_mm", race, i)
					age := randAge()

					crs = append(crs, playerrepo.NewCreator().
						Email(email).Name(name).Age(age).
						Interests(interests).Race(race))
				}

				err = repo.CreateMany(ctx, crs...)
				assert.NoError(t, err)

				err = repo.CreateMany(ctx, []*playerrepo.Creator{}...)
				assert.NoError(t, err)
			})

			t.Run("Error", func(t *testing.T) {
				err := repo.CreateMany(ctx, playerrepo.NewCreator())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				cancel()
				err = repo.CreateMany(cctx, playerrepo.NewCreator())
				assert.Error(t, err)
			})
		})

		t.Run("Query", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				players, err := repo.Query(ctx, playerrepo.NewQueryer().
					Where(
						playerrepo.UpdatedAtIsNull(),
						playerrepo.CreatedAtIsNotNull()),
				)
				assert.NoError(t, err)
				require.Len(t, players, 100)
				for _, u := range players {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Interests, 3)
				}

				// with predicates
				players, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(playerrepo.IDEq("2"), playerrepo.IDNotEq("1"),
						playerrepo.IDGt("1"), playerrepo.IDGtOrEq("2"),
						playerrepo.IDLt("3"), playerrepo.IDLtOrEq("2"),
						playerrepo.IDIn("2"), playerrepo.IDNotIn("1"),
					),
				)
				assert.NoError(t, err)
				assert.Len(t, players, 1)

				players, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(
						playerrepo.AgeEq(18), playerrepo.AgeNotEq(30),
						playerrepo.AgeGt(17), playerrepo.AgeGtOrEq(18),
						playerrepo.AgeLt(30), playerrepo.AgeLtOrEq(19),
						playerrepo.AgeIn(18), playerrepo.AgeNotIn(30),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(players))

				players, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(
						playerrepo.RaceEq(player.RaceNorn),
						playerrepo.RaceNotEq(player.RaceHuman),
						playerrepo.RaceIn(player.RaceNorn),
						playerrepo.RaceNotIn(player.RaceHuman),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(players))

				// with sort
				// get last user
				players, err = repo.Query(ctx, playerrepo.NewQueryer().
					Sort(
						playerrepo.Desc(playerrepo.ColumnID),
						playerrepo.Asc(playerrepo.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(players))
				assert.Equal(t, "charr_100_mm", players[0].Name)

				// with limit and offset
				players, err = repo.Query(ctx, playerrepo.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, players, 1)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Query(cctx, playerrepo.NewQueryer())
				assert.Error(t, err)
			})
		})

		t.Run("QueryOne", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				usr, err := repo.QueryOne(ctx, playerrepo.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)
				assert.Len(t, usr.Interests, 3)

				_, err = repo.QueryOne(ctx, playerrepo.NewQueryer().Where(playerrepo.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.QueryOne(cctx, playerrepo.NewQueryer())
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

				a := playerrepo.NewAggregator(&agg).
					Aggregate(
						playerrepo.Avg(playerrepo.ColumnAge),
						playerrepo.Min(playerrepo.ColumnAge),
						playerrepo.Max(playerrepo.ColumnAge),
						playerrepo.Count(playerrepo.ColumnAge),
						playerrepo.Sum(playerrepo.ColumnAge),
						playerrepo.None(playerrepo.ColumnRace),
					).
					Where(playerrepo.AgeGt(18), playerrepo.RaceNotEq(player.RaceHuman)).
					GroupBy(playerrepo.ColumnRace).
					Sort(playerrepo.Asc(playerrepo.ColumnRace))

				err := repo.Aggregate(ctx, a)
				require.NoError(t, err)
				assert.Len(t, agg, 2)

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

		newInterests := []string{"anime", "mangga", "samyang noddles"}

		t.Run("Update", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				preds := []comparison.PredFunc{
					playerrepo.IDEq("1"), playerrepo.IDNotEq("2"),
					playerrepo.IDGt("0"), playerrepo.IDGtOrEq("1"),
					playerrepo.IDLt("2"), playerrepo.IDLtOrEq("1"),
				}

				email := "titan@gg.io"
				name := "titan"
				age := 300
				rowsAffected, err := repo.Update(ctx,
					playerrepo.NewUpdater().
						Email(email).
						Name(name).
						Age(age).
						Race(player.RaceTitan).
						Interests(newInterests).
						UpdatedAt(&now).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				playr, err := repo.QueryOne(ctx, playerrepo.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)

				assert.Equal(t, email, playr.Email)
				assert.Equal(t, name, playr.Name)
				assert.Equal(t, age, playr.Age)
				assert.NotNil(t, playr.UpdatedAt)
				assert.Len(t, playr.Interests, 3)
			})

			_, err = repo.Update(ctx, playerrepo.NewUpdater())
			assert.NoError(t, err)

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Update(cctx, playerrepo.NewUpdater().Age(1))
				assert.Error(t, err)
			})
		})

		t.Run("Delete", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				preds := []comparison.PredFunc{
					playerrepo.IDEq("1"), playerrepo.IDNotEq("2"),
					playerrepo.IDGt("0"), playerrepo.IDGtOrEq("1"),
					playerrepo.IDLt("2"), playerrepo.IDLtOrEq("1"),
				}
				// delete one
				rowsAffected, err := repo.Delete(ctx,
					playerrepo.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)

				usr, err := repo.QueryOne(ctx,
					playerrepo.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)

				// delete all
				rowsAffected, err = repo.Delete(ctx, playerrepo.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				cancel()
				_, err = repo.Delete(cctx, playerrepo.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func newRepoTestRunnerTx(repo playerrepo.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		var err error
		ctx := context.Background()

		newTx := func(ctx context.Context, t *testing.T) nero.Tx {
			tx, err := repo.Tx(ctx)
			assert.NoError(t, err)
			return tx
		}

		interests := []string{"pizza", "cats", "dogs"}
		t.Run("CreateTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				now := time.Now()
				for i := 1; i <= 50; i++ {
					race := player.RaceHuman
					if i%2 == 0 {
						race = player.RaceCharr
					} else if i%3 == 0 {
						race = player.RaceNorn
					} else if i%4 == 0 {
						race = player.RaceSylvari
					}

					email := fmt.Sprintf("%s_%d@gg.io", race, i)
					name := fmt.Sprintf("%s_%d", race, i)
					age := randAge()

					tx := newTx(ctx, t)
					id, err := repo.CreateTx(ctx, tx, playerrepo.NewCreator().
						Email(email).Name(name).
						Age(age).Race(race).
						Interests(interests).
						UpdatedAt(&now))
					assert.NoError(t, err)
					assert.NotZero(t, id)
					assert.NoError(t, tx.Commit())
				}
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				id, err := repo.CreateTx(ctx, tx, playerrepo.NewCreator())
				assert.Error(t, err)
				assert.Zero(t, id)

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.CreateTx(cctx, tx, playerrepo.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("CreateManyTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				crs := []*playerrepo.Creator{}
				for i := 51; i <= 100; i++ {
					race := player.RaceHuman
					if i%2 == 0 {
						race = player.RaceCharr
					} else if i%3 == 0 {
						race = player.RaceNorn
					} else if i%4 == 0 {
						race = player.RaceSylvari
					}

					email := fmt.Sprintf("%s_%d_mm@gg.io", race, i)
					name := fmt.Sprintf("%s_%d_mm", race, i)
					age := randAge()

					crs = append(crs, playerrepo.NewCreator().
						Email(email).
						Name(name).
						Age(age).
						Interests(interests).
						Race(race))
				}
				tx := newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, crs...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				err = repo.CreateManyTx(ctx, tx, []*playerrepo.Creator{}...)
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				tx := newTx(ctx, t)
				err := repo.CreateManyTx(ctx, tx, playerrepo.NewCreator())
				assert.Error(t, err)

				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				err = repo.CreateManyTx(cctx, tx, playerrepo.NewCreator())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("QueryTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				users, err := repo.QueryTx(ctx, tx, playerrepo.NewQueryer().
					Where(playerrepo.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.NotNil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Interests, 3)
				}
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, playerrepo.NewQueryer().
					Where(playerrepo.UpdatedAtIsNull()))
				assert.NoError(t, err)
				require.Len(t, users, 50)
				for _, u := range users {
					assert.NotNil(t, u.Email)
					assert.NotNil(t, u.Name)
					assert.Nil(t, u.UpdatedAt)
					assert.NotNil(t, u.CreatedAt)
					assert.Len(t, u.Interests, 3)
				}
				assert.NoError(t, tx.Commit())

				// with predicates
				tx = newTx(ctx, t)
				users, err = repo.QueryTx(ctx, tx, playerrepo.NewQueryer().
					Where(playerrepo.IDEq("2"), playerrepo.IDNotEq("1"),
						playerrepo.IDGt("1"), playerrepo.IDGtOrEq("2"),
						playerrepo.IDLt("3"), playerrepo.IDLtOrEq("2")))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(
						playerrepo.AgeEq(18), playerrepo.AgeNotEq(30),
						playerrepo.AgeGt(17), playerrepo.AgeGtOrEq(18),
						playerrepo.AgeLt(30), playerrepo.AgeLtOrEq(19),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(
						playerrepo.RaceEq(player.RaceNorn),
						playerrepo.RaceNotEq(player.RaceHuman),
					),
				)
				assert.NoError(t, err)
				assert.NotZero(t, len(users))
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(playerrepo.UpdatedAtIsNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().
					Where(playerrepo.UpdatedAtIsNotNull()))
				assert.NoError(t, err)
				assert.Len(t, users, 50)
				assert.NoError(t, tx.Commit())

				// with sort
				// get last user
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().
					Sort(
						playerrepo.Desc(playerrepo.ColumnID),
						playerrepo.Asc(playerrepo.ColumnCreatedAt),
					),
				)
				assert.NoError(t, err)
				require.NotZero(t, len(users))
				assert.Equal(t, "charr_100_mm", users[0].Name)
				assert.NoError(t, tx.Commit())

				// with limit and offset
				tx = newTx(ctx, t)
				users, err = repo.Query(ctx, playerrepo.NewQueryer().Limit(1).Offset(1))
				assert.NoError(t, err)
				assert.Len(t, users, 1)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryTx(cctx, tx, playerrepo.NewQueryer())
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("QueryOneTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				tx := newTx(ctx, t)
				usr, err := repo.QueryOneTx(ctx, tx, playerrepo.NewQueryer())
				assert.NoError(t, err)
				assert.NotNil(t, usr)
				assert.Equal(t, "1", usr.ID)
				assert.Len(t, usr.Interests, 3)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				_, err = repo.QueryOne(ctx, playerrepo.NewQueryer().Where(playerrepo.IDEq("9999")))
				assert.Error(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(cctx, t)
				cancel()
				_, err = repo.QueryOneTx(cctx, tx, playerrepo.NewQueryer())
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

				a := playerrepo.NewAggregator(&agg).
					Aggregate(
						playerrepo.Avg(playerrepo.ColumnAge),
						playerrepo.Min(playerrepo.ColumnAge),
						playerrepo.Max(playerrepo.ColumnAge),
						playerrepo.Count(playerrepo.ColumnAge),
						playerrepo.Sum(playerrepo.ColumnAge),
						playerrepo.None(playerrepo.ColumnRace),
					).
					Where(
						playerrepo.AgeGt(18),
						playerrepo.RaceNotEq(player.RaceTitan),
					).
					GroupBy(playerrepo.ColumnRace).
					Sort(playerrepo.Asc(playerrepo.ColumnRace))

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
				preds := []comparison.PredFunc{
					playerrepo.IDEq("1"), playerrepo.IDNotEq("2"),
					playerrepo.IDGt("0"), playerrepo.IDGtOrEq("1"),
					playerrepo.IDLt("2"), playerrepo.IDLtOrEq("1"),
				}

				email := "titan@gg.io"
				name := "titan"
				age := 300

				tx := newTx(ctx, t)
				rowsAffected, err := repo.UpdateTx(ctx, tx,
					playerrepo.NewUpdater().
						Email(email).
						Name(name).
						Age(age).
						Race(player.RaceTitan).
						Interests(newTags).
						UpdatedAt(&now).
						Where(preds...),
				)
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx, playerrepo.NewQueryer().
					Where(preds...))
				assert.NoError(t, err)
				assert.NoError(t, tx.Commit())

				assert.Equal(t, email, usr.Email)
				assert.Equal(t, name, usr.Name)
				assert.Equal(t, age, usr.Age)
				assert.NotNil(t, usr.UpdatedAt)
				assert.Len(t, usr.Interests, 1)
			})

			tx := newTx(ctx, t)
			_, err = repo.UpdateTx(ctx, tx, playerrepo.NewUpdater())
			assert.NoError(t, err)

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx = newTx(cctx, t)
				cancel()
				_, err = repo.UpdateTx(cctx, tx, playerrepo.NewUpdater().Age(1))
				assert.Error(t, err)
				assert.Error(t, tx.Commit())
			})
		})

		t.Run("DeleteTx", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				preds := []comparison.PredFunc{
					playerrepo.IDEq("1"), playerrepo.IDNotEq("2"),
					playerrepo.IDGt("0"), playerrepo.IDGtOrEq("1"),
					playerrepo.IDLt("2"), playerrepo.IDLtOrEq("1"),
				}
				// delete one
				tx := newTx(ctx, t)
				rowsAffected, err := repo.DeleteTx(ctx, tx,
					playerrepo.NewDeleter().Where(preds...))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), rowsAffected)
				assert.NoError(t, tx.Commit())

				tx = newTx(ctx, t)
				usr, err := repo.QueryOne(ctx,
					playerrepo.NewQueryer().Where(preds...))
				assert.Error(t, err, sql.ErrNoRows)
				assert.Nil(t, usr)
				assert.NoError(t, tx.Commit())

				// delete all
				tx = newTx(ctx, t)
				rowsAffected, err = repo.Delete(ctx, playerrepo.NewDeleter())
				assert.NoError(t, err)
				assert.Equal(t, int64(99), rowsAffected)
				assert.NoError(t, tx.Commit())
			})

			t.Run("Error", func(t *testing.T) {
				cctx, cancel := context.WithCancel(ctx)
				tx := newTx(ctx, t)
				cancel()
				_, err = repo.DeleteTx(cctx, tx, playerrepo.NewDeleter())
				assert.Error(t, err)
			})
		})
	}
}

func randAge() int {
	return rand.Intn(30-18) + 18
}
