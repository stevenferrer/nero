package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

// File is a generated file
type File struct {
	name string
	buf  *bytes.Buffer
}

// Render renders the file to the specified path
func (fl *File) Render(basePath string) error {
	filePath := path.Join(basePath, fl.name)
	f, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "create base path")
	}
	defer f.Close()

	_, err = f.Write(fl.Bytes())
	if err != nil {
		return errors.Wrap(err, "write file")
	}

	return errors.Wrap(formatSource(filePath), "format source")
}

// FileName returns the filename
func (fl *File) FileName() string {
	return fl.name
}

// Bytes returns the bytes
func (fl *File) Bytes() []byte {
	return fl.buf.Bytes()
}

// formatSource removes unneeded imports from the given Go source file and runs gofmt on it.
// Copied from goa codebase https://github.com/goadesign/goa/blob/v3/codegen/file.go#L136
func formatSource(filePath string) error {
	// Make sure file parses and print content if it does not.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		content, _ := ioutil.ReadFile(filePath)
		var buf bytes.Buffer
		scanner.PrintError(&buf, err)
		return fmt.Errorf("%s\n========\nContent:\n%s", buf.String(), content)
	}

	// Clean unused imports
	imps := astutil.Imports(fset, file)
	for _, group := range imps {
		for _, imp := range group {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(file, path) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fset, file, imp.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, file, path)
				}
			}
		}
	}
	ast.SortImports(fset, file)
	w, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if err := format.Node(w, fset, file); err != nil {
		return err
	}
	w.Close()

	// Format code using goimport standard
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	opt := imports.Options{
		Comments:   true,
		FormatOnly: true,
	}
	bs, err = imports.Process(filePath, bs, &opt)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, bs, os.ModePerm)
}
