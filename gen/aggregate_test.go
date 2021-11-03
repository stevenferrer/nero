package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero/gen/internal"
)

func Test_newAggregateFile(t *testing.T) {
	u := internal.User{}
	f, err := newAggregateFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(f.Bytes())
	require.NoError(t, err)

	_, err = newAggregateFile(nil)
	assert.Error(t, err)
}
