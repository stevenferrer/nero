package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDeleter(t *testing.T) {
	deleter := newDeleter()
	expect := strings.TrimSpace(`
type Deleter struct {
	collection string
	pfs        []PredicateFunc
}

func NewDeleter() *Deleter {
	return &Deleter{
		collection: collection,
	}
}

func (d *Deleter) Where(pfs ...PredicateFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", deleter))
	assert.Equal(t, expect, got)
}
