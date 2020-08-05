package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero/test/integration/basic/repository/user"
)

func TestColumnStringer(t *testing.T) {
	c := user.Column(99)
	assert.Equal(t, "invalid", c.String())
}
