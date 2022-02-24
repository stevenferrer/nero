package gen

import (
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
)

// Generate generates the repository code
func Generate(schema nero.Schema) ([]*File, error) {
	files := []*File{}
	file, err := newMetaFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "meta file")
	}
	files = append(files, file)

	file, err = newPredicateFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "predicate file")
	}
	files = append(files, file)

	file, err = newSortFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "sort file")
	}
	files = append(files, file)

	file, err = newAggregateFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "aggregate file")
	}
	files = append(files, file)

	file, err = newRepositoryFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "repository file")
	}
	files = append(files, file)

	for _, tmpl := range schema.Templates() {
		buf, err := newTemplate(schema, tmpl)
		if err != nil {
			return nil, errors.Wrap(err, "template file")
		}

		files = append(files, &File{name: tmpl.Filename(), buf: buf.Bytes()})
	}

	return files, nil
}
