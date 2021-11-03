package gen

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
)

func newTemplate(schema *nero.Schema, template nero.Template) (*bytes.Buffer, error) {
	tmpl, err := nero.ParseTemplate(template)
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, schema)
	return buf, err
}
