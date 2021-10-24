package nero

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type myType struct{}

func TestFuncs(t *testing.T) {
	t.Run("typeFunc", func(t *testing.T) {
		got := typeFunc(1)
		expect := "int"
		assert.Equal(t, expect, got)

		got = typeFunc(&myType{})
		expect = "nero.myType"
		assert.Equal(t, expect, got)
	})

	t.Run("rawTypeFunc", func(t *testing.T) {
		got := rawTypeFunc(1)
		expect := "int"
		assert.Equal(t, expect, got)

		got = rawTypeFunc(&myType{})
		expect = "*nero.myType"
		assert.Equal(t, expect, got)
	})

	t.Run("zeroFunc", func(t *testing.T) {
		got := zeroValueFunc(0)
		expect := "0"
		assert.Equal(t, expect, got)

		got = zeroValueFunc([]string{})
		expect = "nil"
		assert.Equal(t, expect, got)

		got = zeroValueFunc(true)
		expect = "false"
		assert.Equal(t, expect, got)

		got = zeroValueFunc(myType{})
		expect = "(nero.myType{})"
		assert.Equal(t, expect, got)

		got = zeroValueFunc([1]int{1})
		expect = "([1]int{})"
		assert.Equal(t, expect, got)

		got = zeroValueFunc([1][1]*myType{})
		expect = "([1][1]*nero.myType{})"
		assert.Equal(t, expect, got)

		got = zeroValueFunc([0]interface{}{})
		expect = "([0]interface {}{})"
		assert.Equal(t, expect, got)

		got = zeroValueFunc("")
		expect = `""`
		assert.Equal(t, expect, got)
	})

	t.Run("isType", func(t *testing.T) {
		now := time.Now()
		assert.True(t, isTypeFunc(now, "time.Time"))
		assert.True(t, isTypeFunc(&now, "time.Time"))
	})
	assert.Len(t, prependToFields(&Field{}, []*Field{}), 1)
	assert.NotEmpty(t, fileHeadersFunc())
}
