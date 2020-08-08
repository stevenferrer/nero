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

func Test_newUpdater(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	updater := newUpdater(schema)
	expect := strings.TrimSpace(`
type Updater struct {
	name      string
	group     string
	updatedAt *time.Time
	pfs       []PredFunc
}

func NewUpdater() *Updater {
	return &Updater{}
}

func (u *Updater) Name(name string) *Updater {
	u.name = name
	return u
}

func (u *Updater) Group(group string) *Updater {
	u.group = group
	return u
}

func (u *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	u.updatedAt = updatedAt
	return u
}

func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", updater))
	assert.Equal(t, expect, got)
}
