package sqlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newCreateBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newCreateBlock(schema)
	expect := strings.TrimSpace(`
func (sl *SQLiteRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	tx, err := sl.Tx(ctx)
	if err != nil {
		return 0, err
	}

	id, err := sl.CreateTx(ctx, tx, c)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return id, tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newCreateManyBlock(t *testing.T) {
	block := newCreateManyBlock()
	expect := strings.TrimSpace(`
func (sl *SQLiteRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	tx, err := sl.Tx(ctx)
	if err != nil {
		return err
	}

	err = sl.CreateManyTx(ctx, tx, cs...)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newCreateTxBlock(t *testing.T) {
	t.Run("schema with integer id", func(t *testing.T) {
		schema, err := gen.BuildSchema(new(example.User))
		require.NoError(t, err)
		require.NotNil(t, schema)

		block := newCreateTxBlock(schema)
		expect := strings.TrimSpace(`
func (sl *SQLiteRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.group, c.updatedAt).
		RunWith(txx)
	if log := sl.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Create").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", block))
		assert.Equal(t, expect, got)
	})

	t.Run("schema with string id", func(t *testing.T) {
		schema, err := gen.BuildSchema(new(example.Group))
		require.NoError(t, err)
		require.NotNil(t, schema)

		block := newCreateTxBlock(schema)
		expect := strings.TrimSpace(`
func (sl *SQLiteRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (string, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return "", errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(c.collection).
		Columns(c.columns...).
		Values(c.name, c.updatedAt).
		RunWith(txx)
	if log := sl.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "Create").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.ExecContext(ctx)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", block))
		assert.Equal(t, expect, got)
	})
}

func Test_newCreateManyTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newCreateManyTxBlock(schema)
	expect := strings.TrimSpace(`
func (sl *SQLiteRepository) CreateManyTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Insert(cs[0].collection).
		Columns(cs[0].columns...)
	for _, c := range cs {
		qb = qb.Values(c.name, c.group, c.updatedAt)
	}
	if log := sl.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("op", "CreateMany").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	_, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
