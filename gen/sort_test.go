package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/gen/internal"
)

func Test_newSortFile(t *testing.T) {
	u := internal.User{}
	f, err := newSortFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(f.Bytes())
	require.NoError(t, err)

	_, err = newSortFile(nil)
	assert.Error(t, err)
}
