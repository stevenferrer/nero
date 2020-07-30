package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColMethods(t *testing.T) {
	col := Col{Name: "first_name"}
	expect := "FirstName"
	assert.Equal(t, expect, col.CamelName())
	expect = "firstName"
	assert.Equal(t, expect, col.LowerCamelName())
}
