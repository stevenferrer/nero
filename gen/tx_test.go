package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTx(t *testing.T) {
	stmt := newTx()
	expect := strings.TrimSpace(`	
// Tx is a transaction type
type Tx interface {
	Commit() error
	Rollback() error
}
`)
	got := fmt.Sprintf("%#v", stmt)
	assert.Equal(t, expect, got)
}
