package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func Test_newRepositoryFile(t *testing.T) {
	schema := (&example.User{}).Schema()
	buf, err := newRepositoryFile(schema)
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newRepositoryFile(nil)
	assert.Error(t, err)
}
