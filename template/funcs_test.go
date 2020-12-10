package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type myType struct{}

func TestFuncs(t *testing.T) {
	t.Run("typeFunc", func(t *testing.T) {
		got := typeFunc(1)
		expect := "int"
		assert.Equal(t, expect, got)

		got = typeFunc(&myType{})
		expect = "template.myType"
		assert.Equal(t, expect, got)
	})

	t.Run("zeroFunc", func(t *testing.T) {
		got := zeroFunc(0)
		expect := "0"
		assert.Equal(t, expect, got)

		got = zeroFunc([]string{})
		expect = "nil"
		assert.Equal(t, expect, got)

		got = zeroFunc(true)
		expect = "false"
		assert.Equal(t, expect, got)

		got = zeroFunc(myType{})
		expect = "(template.myType{})"
		assert.Equal(t, expect, got)

		got = zeroFunc([1]int{1})
		expect = "([1]int{})"
		assert.Equal(t, expect, got)

		got = zeroFunc([1][1]*myType{})
		expect = "([1][1]*template.myType{})"
		assert.Equal(t, expect, got)

		got = zeroFunc("")
		expect = `""`
		assert.Equal(t, expect, got)
	})

	t.Run("pluralFunc", func(t *testing.T) {
		got := pluralFunc("user")
		expect := "users"
		assert.Equal(t, got, expect)

		got = pluralFunc("matrix")
		expect = "matrices"
		assert.Equal(t, got, expect)
	})

	t.Run("lowerCamelFunc", func(t *testing.T) {
		got := lowerCamelFunc("TheMatrix")
		expect := "theMatrix"
		assert.Equal(t, got, expect)
	})
}
