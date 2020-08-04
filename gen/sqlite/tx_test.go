package sqlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTxBlock(t *testing.T) {
	block := newTxBlock()
	expect := strings.TrimSpace(`
func (sqlr *SQLiteRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return sqlr.db.BeginTx(ctx, nil)
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
