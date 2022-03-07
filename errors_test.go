package nero_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevenferrer/nero"
)

func TestErrRequiredField(t *testing.T) {
	err := nero.NewErrRequiredField("Name")
	expect := `Name field is required`
	assert.Equal(t, expect, err.Error())
}
