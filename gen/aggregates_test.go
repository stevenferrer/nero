package gen

import (
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

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)
}
