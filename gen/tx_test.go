package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTx(t *testing.T) {
	stmnt := newTx()
	expect := strings.TrimSpace(`
func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}`)
	got := fmt.Sprintf("%#v", stmnt)
	assert.Equal(t, expect, got)
}
