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
	Create(*Creator) (id int64, err error)
	Query(*Queryer) ([]*internal.Example, error)
	Update(*Updater) (rowsAffected int64, err error)
	Delete(*Deleter) (rowsAffected int64, err error)
}

`)
	got := fmt.Sprintf("%#v", stmt)
	assert.Equal(t, expect, got)
}
