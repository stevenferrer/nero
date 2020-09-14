package gen

import (
	"fmt"
	"go/format"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newRepositoryFile(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	buf, err := newRepositoryFile(schema)
	require.NoError(t, err)

	src, err := format.Source(buf.Bytes())
	require.NoError(t, err)

	fmt.Printf("%v\n", string(src))
}
