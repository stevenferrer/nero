package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newUpdater(t *testing.T) {
	schema, err := buildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	updater := newUpdater(schema)
	expect := strings.TrimSpace(`
type Updater struct {
	collection string
	columns    []string
	name       string
	updatedAt  *time.Time
	pfs        []PredicateFunc
}

func NewUpdater() *Updater {
	return &Updater{
		collection: collection,
		columns:    []string{"name", "updated_at"},
	}
}

func (u *Updater) Name(name string) *Updater {
	u.name = name
	return u
}

func (u *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	u.updatedAt = updatedAt
	return u
}

func (u *Updater) Where(pfs ...PredicateFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", updater))
	assert.Equal(t, expect, got)
}
