package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newRepository(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	stmt := newRepository(schema)
	expect := strings.TrimSpace(`
// Repository is the contract for storing User
type Repository interface {
	Tx(context.Context) (nero.Tx, error)
	Create(context.Context, *Creator) (id int64, err error)
	CreateMany(context.Context, ...*Creator) error
	CreateTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	Query(context.Context, *Queryer) ([]*example.User, error)
	QueryOne(context.Context, *Queryer) (*example.User, error)
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*example.User, error)
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*example.User, error)
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	Aggregate(context.Context, *Aggregator) error
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", stmt))
	assert.Equal(t, expect, got)
}
