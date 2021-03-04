package gen

import (
	"go/format"
	"testing"

	"github.com/sf9v/nero/example"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newAggregateFile(t *testing.T) {
	schema := (&example.User{}).Schema()
	require.NotNil(t, schema)

	buf, err := newAggregateFile(schema)
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newAggregateFile(nil)
	assert.Error(t, err)
}
