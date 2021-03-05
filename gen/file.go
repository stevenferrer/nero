package gen

import (
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/sf9v/nero/x/etc"
)

// File is a generated file
type File struct {
	name string
	buf  []byte
}

// Render renders the file to the specified path
func (f *File) Render(basePath string) error {
	filePath := path.Join(basePath, f.name)
	of, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "create base path")
	}
	defer of.Close()

	_, err = of.Write(f.buf)
	if err != nil {
		return errors.Wrap(err, "write file")
	}

	return errors.Wrap(etc.FmtSrc(filePath), "format source")
}

// Filename returns the filename
func (f *File) Filename() string {
	return f.name
}

// Bytes returns the bytes
func (f *File) Bytes() []byte {
	return f.buf[:]
}
