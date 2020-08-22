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
	// Tx returns a transaction
	Tx(context.Context) (nero.Tx, error)
	// Create creates a User
	Create(context.Context, *Creator) (id int64, err error)
	// CreateTx creates a User in a transaction
	CreateTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	// CreateMany creates Users
	CreateMany(context.Context, ...*Creator) error
	// CreateManyTx creates Users in a transaction
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries Users
	Query(context.Context, *Queryer) ([]*example.User, error)
	// QueryTx queries Users in a transaction
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*example.User, error)
	// QueryOne queries a User
	QueryOne(context.Context, *Queryer) (*example.User, error)
	// QueryOneTx queries a User in a transaction
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*example.User, error)
	// Update updates Users
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates Users in a transaction
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes Users
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes User in a transaction
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate aggregates Users
	Aggregate(context.Context, *Aggregator) error
	// Aggregate aggregates Users in a transaction
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", stmt))
	assert.Equal(t, expect, got)
}
