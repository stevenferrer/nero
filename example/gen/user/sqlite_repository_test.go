package user_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero/example/gen/user"
)

func TestRepository(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	assert.NoError(t, err)
	defer db.Close()

	err = createTable(db)
	assert.NoError(t, err)

	sqliteRepo := user.NewSQLiteRepository(db)

	// create
	for i := 1; i <= 10; i++ {
		id, err := sqliteRepo.Create(
			user.NewCreator().
				Email(fmt.Sprintf("user%d@gg.io", i)).
				Name(fmt.Sprintf("user%d", i)),
		)
		assert.NoError(t, err)
		assert.NotZero(t, id)
	}

	// query
	users, err := sqliteRepo.Query(
		user.NewQueryer().
			Where(user.IDGt(1)).
			Limit(4).Offset(0),
	)
	assert.NoError(t, err)
	assert.Len(t, users, 4)
	for _, u := range users {
		assert.NotNil(t, u.CreatedAt)
	}

	// update
	now := time.Now()
	rowsAffected, err := sqliteRepo.Update(
		user.NewUpdater().
			Email("sf9v@gg.io").
			Name("sf9v").
			UpdatedAt(&now).
			Where(user.IDGt(1), user.IDLt(3)),
	)
	assert.NoError(t, err)
	assert.NotZero(t, rowsAffected)

	users, err = sqliteRepo.Query(
		user.NewQueryer().
			Where(
				user.EmailEq("sf9v@gg.io"),
				user.NameEq("sf9v"),
				user.UpdatedAtEq(&now),
			),
	)
	assert.NoError(t, err)
	for _, u := range users {
		assert.NotNil(t, u.UpdatedAt)
	}

	users, err = sqliteRepo.Query(
		user.NewQueryer().
			Where(user.IDGt(1), user.IDLt(3)),
	)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, users[0].Email, "sf9v@gg.io")
	assert.Equal(t, users[0].Name, "sf9v")

	// delete
	rowsAffected, err = sqliteRepo.Delete(user.NewDeleter())
	assert.NoError(t, err)
	assert.Equal(t, int64(10), rowsAffected)

	users, err = sqliteRepo.Query(user.NewQueryer())
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func createTable(db *sql.DB) error {
	createUsersTable := `create table users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"email" TEXT UNIQUE,
		"name" TEXT NOT NULL,
		"updated_at" DATETIME,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	stmnt, err := db.Prepare(createUsersTable)
	if err != nil {
		return err
	}
	_, err = stmnt.Exec()
	if err != nil {
		return err
	}

	return nil
}
