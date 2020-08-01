package internal

import (
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"

	"github.com/sf9v/nero"
)

// Schema is an internal schema
type Schema struct {
	Coln  string
	Typ   *Typ
	Ident *Col
	Cols  []*Col
	Pkg   string
}

// BuildSchema builds schema from a nero.Schemaer to Schema
func BuildSchema(s nero.Schemaer) (*Schema, error) {
	ns := s.Schema()
	schema := &Schema{
		Coln: ns.Collection,
		Typ:  NewTyp(s),
		Cols: []*Col{},
		Pkg:  ns.Pkg,
	}

	identCnt := 0
	for _, column := range ns.Columns {
		cfg := column.Cfg()
		col := &Col{
			Name:  cfg.Name,
			Field: strcase.ToCamel(cfg.Name),
			Typ:   NewTyp(cfg.T),
			Ident: cfg.Ident,
			Auto:  cfg.Auto,
		}

		if len(cfg.Field) > 0 {
			col.Field = cfg.Field
		}

		if cfg.Ident {
			schema.Ident = col
			identCnt++
		}

		schema.Cols = append(schema.Cols, col)
	}

	if identCnt == 0 {
		return nil, errors.New("at least one ident column is required")
	}

	if identCnt > 1 {
		return nil, errors.New("only one ident column is allowed")
	}

	return schema, nil
}
