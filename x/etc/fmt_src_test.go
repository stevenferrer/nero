package etc_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/nero/x/etc"
)

const src = `
package main

import ("fmt"
"os"
nero "github.com/stevenferrer/nero"
)
func main() {
fmt.Println("Hello, world!")
}
`

func TestFmtSrc(t *testing.T) {
	filename := "temp.go"
	filepath := path.Join(os.TempDir(), filename)
	err := ioutil.WriteFile(filepath, []byte(src), 0644)
	require.NoError(t, err)

	err = etc.FmtSrc(filepath)
	assert.NoError(t, err)

	// cleanup
	assert.NoError(t, os.Remove(filepath))

	err = etc.FmtSrc(filepath)
	assert.Error(t, err)
}
