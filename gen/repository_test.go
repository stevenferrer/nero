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
// Repository is a User repository
type Repository interface {
	// Tx returns a new transaction
	Tx(context.Context) (nero.Tx, error)
	// Create creates a User
	Create(context.Context, *Creator) (id int64, err error)
	// CreateTx creates a User inside transaction
	CreateTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	// CreateMany is a batch-create for User
	CreateMany(context.Context, ...*Creator) error
	// CreateManyTx is a batch-create for User inside transaction
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	// Query is used for querying many User
	Query(context.Context, *Queryer) ([]*example.User, error)
	// QueryTx is used for querying many User inside transaction
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*example.User, error)
	// QueryOne is used for querying a single User
	QueryOne(context.Context, *Queryer) (*example.User, error)
	// QueryOneTx is used for querying a single User inside transaction
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*example.User, error)
	// Update updates User
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates User inside transaction
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes User
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes User inside transaction
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate is used for doing aggregation
	Aggregate(context.Context, *Aggregator) error
	// Aggregate is used for doing aggregation inside transaction
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", stmt))
	assert.Equal(t, expect, got)
}
