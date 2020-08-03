package gen

import (
	"bytes"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func TestRenderFiles(t *testing.T) {
	files := Files{
		{
			name: "test.txt",
			buff: bytes.NewBuffer([]byte("test")),
			jf:   &jen.File{},
		},
	}

	err := files.Render("asdf")
	assert.Error(t, err)
}
