package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleSchema(t *testing.T) {
	e := new(Example)
	s := e.Schema()
	assert.Equal(t, "example", s.Pkg)
	assert.Equal(t, "examples", s.Collection)
	assert.Len(t, s.Columns, 4)
}

func TestExample2Schema(t *testing.T) {
	e := new(Example2)
	s := e.Schema()
	assert.Equal(t, "example2", s.Pkg)
	assert.Equal(t, "examples", s.Collection)
	assert.Len(t, s.Columns, 4)
}
