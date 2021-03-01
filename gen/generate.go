package gen

import (
	"github.com/pkg/errors"
	"github.com/sf9v/nero"
)

// Generate generates the repository code
func Generate(schema *nero.Schema) ([]*File, error) {
	files := []*File{}
	buf, err := newMetaFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "meta file")
	}
	files = append(files, &File{name: "meta.go", buf: buf.Bytes()})

	buf, err = newPredicatesFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "predicates file")
	}
	files = append(files, &File{name: "predicates.go", buf: buf.Bytes()})

	buf, err = newSortsFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "sorts file")
	}
	files = append(files, &File{name: "sorts.go", buf: buf.Bytes()})

	buf, err = newAggregatesFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "aggregates file")
	}
	files = append(files, &File{
		name: "aggregates.go",
		buf:  buf.Bytes(),
	})

	buf, err = newRepositoryFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "repository file")
	}
	files = append(files, &File{name: "repository.go", buf: buf.Bytes()})

	for _, tmpl := range schema.Templaters() {
		buf, err = newTemplater(schema, tmpl)
		if err != nil {
			return nil, errors.Wrap(err, "templater file")
		}

		files = append(files, &File{name: tmpl.Filename(), buf: buf.Bytes()})
	}

	return files, nil
}
