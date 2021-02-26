package nero

import (
	"testing"

	"github.com/sf9v/nero/comparison"
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
		expect = "(nero.myType{})"
		assert.Equal(t, expect, got)

		got = zeroFunc([1]int{1})
		expect = "([1]int{})"
		assert.Equal(t, expect, got)

		got = zeroFunc([1][1]*myType{})
		expect = "([1][1]*nero.myType{})"
		assert.Equal(t, expect, got)

		got = zeroFunc([0]interface{}{})
		expect = "([0]interface {}{})"
		assert.Equal(t, expect, got)

		got = zeroFunc("")
		expect = `""`
		assert.Equal(t, expect, got)
	})

	assert.True(t, isNullOp(comparison.IsNull))
	assert.True(t, isNullOp(comparison.IsNotNull))

	assert.True(t, isInOp(comparison.In))
	assert.True(t, isInOp(comparison.NotIn))

	assert.Len(t, prependToColumns(&Column{}, []*Column{}), 1)
}
