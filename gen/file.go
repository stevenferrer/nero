package gen

import (
	"bytes"
	"os"
	"path"

	"github.com/dave/jennifer/jen"
)

// File is a generated file
type File struct {
	name string
	jf   *jen.File
	buff *bytes.Buffer
}

// Name returns the file name
func (fl *File) Name() string {
	return fl.name
}

// Bytes returns the []byte contents of file
func (fl *File) Bytes() []byte {
	return fl.buff.Bytes()
}

// Render renders the file to the base path
// TODO: auto create base path
func (fl *File) Render(basePath string) error {
	f, err := os.Create(path.Join(basePath, fl.name))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(fl.buff.Bytes())
	return err
}

// Files is a collection of File
type Files []*File

// Render renders all files to the base path
func (fls Files) Render(basePath string) error {
	for _, f := range fls {
		err := f.Render(basePath)
		if err != nil {
			return err
		}
	}

	return nil
}
