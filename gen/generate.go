package gen

import (
	"bytes"
	"go/format"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/sf9v/nero"
	gen "github.com/sf9v/nero/gen/internal"
)

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

	pgBuf, err := newPostgresFile(schema)
	if err != nil {
		return nil, errors.Wrap(err, "postgres file")
	}
	files = append(files, &File{
		name: "postgres.go",
		buf:  pgBuf,
	})

	return files, nil
}

type File struct {
	name string
	buf  *bytes.Buffer
}

func (fl *File) Render(basePath string) error {
	f, err := os.Create(path.Join(basePath, fl.name))
	if err != nil {
		return errors.Wrap(err, "create base path")
	}
	defer f.Close()
	src, err := format.Source(fl.buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "format source")
	}

	_, err = f.Write(src)
	return errors.Wrap(err, "write file")
}

func (fl *File) FileName() string {
	return fl.name
}

func (fl *File) Bytes() []byte {
	return fl.buf.Bytes()
}
