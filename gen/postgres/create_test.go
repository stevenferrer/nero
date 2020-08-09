package postgres

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
func (pg *PostgreSQLRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	return pg.create(ctx, pg.db, c)
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
func (pg *PostgreSQLRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return pg.create(ctx, txx, c)
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
func (pg *PostgreSQLRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (string, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return "", errors.New("expecting tx to be *sql.Tx")
	}

	return pg.create(ctx, txx, c)
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", block))
		assert.Equal(t, expect, got)
	})
}

func Test_newCreateRunnerBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newCreateRunnerBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) create(ctx context.Context, runner nero.SqlRunner, c *Creator) (int64, error) {
	columns := []string{}
	values := []interface{}{}
	if c.name != "" {
		columns = append(columns, "\"name\"")
		values = append(values, c.name)
	}
	if c.group != "" {
		columns = append(columns, "\"group_res\"")
		values = append(values, c.group)
	}
	if c.updatedAt != nil {
		columns = append(columns, "\"updated_at\"")
		values = append(values, c.updatedAt)
	}

	qb := squirrel.Insert("\"users\"").
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(runner)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Create").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	var id int64
	err := qb.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newCreateManyBlock(t *testing.T) {
	block := newCreateManyBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	return pg.createMany(ctx, pg.db, cs...)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newCreateManyTxBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newCreateManyTxBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) CreateManyTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return pg.createMany(ctx, txx, cs...)
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newCreateManyRunnerBlock(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	block := newCreateManyRunnerBlock(schema)
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) createMany(ctx context.Context, runner nero.SqlRunner, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	columns := []string{"\"name\"", "\"group_res\"", "\"updated_at\""}
	qb := squirrel.Insert("\"users\"").Columns(columns...)
	for _, c := range cs {
		qb = qb.Values(c.name, c.group, c.updatedAt)
	}

	qb = qb.Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "CreateMany").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	_, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
