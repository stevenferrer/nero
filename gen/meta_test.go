package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func Test_newMetaFile(t *testing.T) {
	schema := (&example.User{}).Schema()
	require.NotNil(t, schema)

	buf, err := newMetaFile(schema)
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newMetaFile(nil)
	assert.Error(t, err)
}
