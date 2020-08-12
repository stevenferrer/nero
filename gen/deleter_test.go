package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newDeleter(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	deleter := newDeleter(schema)
	expect := strings.TrimSpace(`
// Deleter is the delete builder for User
type Deleter struct {
	pfs []PredFunc
}

// NewDeleter returns a delete builder
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where adds predicates to the query
func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", deleter))
	assert.Equal(t, expect, got)
}
