package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDebugBlock(t *testing.T) {
	block := newDebugBlock()
	expect := strings.TrimSpace(`
func (pg *PostgreSQLRepository) Debug(out io.Writer) *PostgreSQLRepository {
	lg := zerolog.New(out).With().Timestamp().Logger()
	return &PostgreSQLRepository{
		db:  pg.db,
		log: &lg,
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}
