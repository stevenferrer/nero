package etc

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

// FmtSrc removes unneeded imports from the given Go source file and runs gofmt on it.
//
// See https://github.com/goadesign/goa/blob/v3/codegen/file.go#L136
func FmtSrc(path string) error {
	// Make sure file parses and print content if it does not.
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		content, _ := ioutil.ReadFile(path)
		buf := &bytes.Buffer{}
		scanner.PrintError(buf, err)
		return errors.Errorf("%s\n========\nContent:\n%s", buf.String(), content)
	}

	// Clean unused imports
	impss := astutil.Imports(fileSet, astFile)
	for _, imps := range impss {
		for _, imp := range imps {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(astFile, path) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fileSet, astFile, imp.Name.Name, path)
				} else {
					astutil.DeleteImport(fileSet, astFile, path)
				}
			}
		}
	}
	ast.SortImports(fileSet, astFile)
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if err := format.Node(w, fileSet, astFile); err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	// Format code using goimport standard
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	b, err = imports.Process(path, b, &imports.Options{
		Comments:   true,
		FormatOnly: true,
	})
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, os.ModePerm)
}
