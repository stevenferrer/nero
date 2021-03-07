package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/gen/internal"
)

func Test_newRepositoryFile(t *testing.T) {
	u := internal.User{}
	f, err := newRepositoryFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(f.Bytes())
	require.NoError(t, err)

	_, err = newRepositoryFile(nil)
	assert.Error(t, err)
}
