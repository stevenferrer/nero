package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newRepository(t *testing.T) {
	schema, err := buildSchema(&gen.Example{})
	require.NoError(t, err)
	stmt := newRepository(schema)
	expect := strings.TrimSpace(`
// Repository is the contract for storing Example
type Repository interface {
	Tx(context.Context) (nero.Tx, error)
	Create(context.Context, *Creator) (id int64, err error)
	Query(context.Context, *Queryer) ([]*internal.Example, error)
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	CreateTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*internal.Example, error)
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
}
`)
	got := fmt.Sprintf("%#v", stmt)
	assert.Equal(t, expect, got)
}
