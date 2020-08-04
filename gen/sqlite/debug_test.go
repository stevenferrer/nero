package sqlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDebugBlock(t *testing.T) {
	block := newDebugBlock()
	expect := strings.TrimSpace(`
func (sqlr *SQLiteRepository) Debug(out io.Writer) *SQLiteRepository {
	lg := zerolog.New(out).With().Timestamp().Logger()
	return &SQLiteRepository{
		db:  sqlr.db,
		log: &lg,
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
