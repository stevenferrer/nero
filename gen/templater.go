package gen

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/sf9v/nero"
)

func newTemplater(schema *nero.Schema, templater nero.Templater) (*bytes.Buffer, error) {
	tmpl, err := nero.ParseTemplater(templater)
	if err != nil {
		return nil, errors.Wrap(err, "parse templater")
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, schema)
	return buf, err
}
