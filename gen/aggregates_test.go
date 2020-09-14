package gen

import (
	"fmt"
	"go/format"
	"testing"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/require"
)

func Test_newAggregatesFile(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	buf, err := newAggregatesFile(schema)
	require.NoError(t, err)

	src, err := format.Source(buf.Bytes())
	require.NoError(t, err)

	fmt.Printf("%v\n", string(src))
}
