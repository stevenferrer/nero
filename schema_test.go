package nero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn(t *testing.T) {
	cfg := NewColumn("id", int64(0)).
		Auto().Ident().Field("ID").Cfg()
	assert.Equal(t, "id", cfg.Name)
	_, ok := cfg.T.(int64)
	assert.True(t, ok)
	assert.True(t, cfg.Ident)
	assert.True(t, cfg.Auto)
	assert.Equal(t, "ID", cfg.Field)
}
