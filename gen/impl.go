package gen

import (
	"bytes"

	"github.com/pkg/errors"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/template"
)

func newImplFile(schema *gen.Schema, tmpl string) (*bytes.Buffer, error) {
	tmplt, err := template.ParseTemplate(tmpl)
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	buf := new(bytes.Buffer)
	err = tmplt.Execute(buf, schema)
	return buf, err
}
