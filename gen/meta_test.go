package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/gen/internal"
)

func Test_newMetaFile(t *testing.T) {
	u := internal.User{}
	buf, err := newMetaFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newMetaFile(nil)
	assert.Error(t, err)
}
