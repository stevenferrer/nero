package gen

import (
	"github.com/pkg/errors"
	"github.com/sf9v/nero"
	gen "github.com/sf9v/nero/gen/internal"
)

// Generate generates the repository code
func Generate(schemaer nero.Schemaer) ([]*File, error) {
	schema, err := gen.BuildSchema(schemaer)
	if err != nil {
		return nil, errors.Wrap(err, "build schema")
	}

	files := []*File{}
	metaBuf, err := newMetaFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "meta file")
	}
	files = append(files, &File{
		name: "meta.go",
		buf:  metaBuf,
	})

	predsBuf, err := newPredicatesFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "predicates file")
	}
	files = append(files, &File{
		name: "predicates.go",
		buf:  predsBuf,
	})

	sortsBuf, err := newSortsFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "sorts file")
	}
	files = append(files, &File{
		name: "sorts.go",
		buf:  sortsBuf,
	})

	aggsBuf, err := newAggregatesFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "aggregates file")
	}
	files = append(files, &File{
		name: "aggregates.go",
		buf:  aggsBuf,
	})

	repoBuf, err := newRepositoryFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "repository file")
	}
	files = append(files, &File{
		name: "repository.go",
		buf:  repoBuf,
	})

	for _, tmpl := range schema.Templates {
		buff, err := newImplFile(schema, tmpl.Content())
		if err != nil {
			return nil, errors.Wrap(err, "repository implementation file")
		}

		files = append(files, &File{
			name: tmpl.Filename(),
			buf:  buff,
		})
	}

	return files, nil
}
