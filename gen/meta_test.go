package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/gen/internal"
)

func Test_newMetaFile(t *testing.T) {
	u := internal.User{}
	f, err := newMetaFile(u.Schema())
	require.NoError(t, err)

	_, err = format.Source(f.Bytes())
	require.NoError(t, err)

	_, err = newMetaFile(nero.Schema{})
	assert.Error(t, err)
}
