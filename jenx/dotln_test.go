package jenx_test

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/jenx"
	"github.com/stretchr/testify/assert"
)

func TestDotln(t *testing.T) {
	got := fmt.Sprintf("%#v", jenx.Dotln("name"))
	expect := fmt.Sprintf("%#v", jen.Op(".").Line().Id("name"))
	assert.Equal(t, expect, got)
}
