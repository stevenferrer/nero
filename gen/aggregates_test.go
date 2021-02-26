package gen

import (
	"go/format"
	"testing"

	"github.com/sf9v/nero/example"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newAggregatesFile(t *testing.T) {
	schema := (&example.User{}).Schema()
	require.NotNil(t, schema)

	buf, err := newAggregatesFile(schema)
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newAggregatesFile(nil)
	assert.Error(t, err)
}
