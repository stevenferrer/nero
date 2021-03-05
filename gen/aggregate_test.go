package gen

import (
	"go/format"
	"testing"

	"github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newAggregateFile(t *testing.T) {
	u := internal.User{}
	buf, err := newAggregateFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newAggregateFile(nil)
	assert.Error(t, err)
}
