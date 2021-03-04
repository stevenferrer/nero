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

	buf, err = newPredicateFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "predicate file")
	}
	files = append(files, &File{name: "predicate.go", buf: buf.Bytes()})

	buf, err = newSortFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "sort file")
	}
	files = append(files, &File{name: "sort.go", buf: buf.Bytes()})

	buf, err = newAggregateFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "aggregate file")
	}
	files = append(files, &File{
		name: "aggregate.go",
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
