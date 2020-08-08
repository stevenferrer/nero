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
	pfs []PredFunc
}

func NewDeleter() *Deleter {
	return &Deleter{}
}

func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", deleter))
	assert.Equal(t, expect, got)
}
