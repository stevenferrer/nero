package nero_test

import (
	"testing"

	"github.com/sf9v/nero"
	"github.com/stretchr/testify/assert"
)

func TestErrRequiredField(t *testing.T) {
	err := nero.NewErrRequiredField("Name")
	expect := `Name field is required`
	assert.Equal(t, expect, err.Error())
}
