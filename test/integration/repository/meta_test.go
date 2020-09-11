package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero/test/integration/repository"
)

func TestColumnStringer(t *testing.T) {
	c := repository.Column(99)
	assert.Equal(t, "invalid", c.String())
}
