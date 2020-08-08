package internal

import (
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/sf9v/mira"

	"github.com/sf9v/nero"
)

// Schema is an internal schema
type Schema struct {
	Collection string
	Type       *mira.Type
	Ident      *Col
	Cols       []*Col
	Pkg        string
}

// BuildSchema builds schema from a nero.Schemaer to Schema
func BuildSchema(s nero.Schemaer) (*Schema, error) {
	ns := s.Schema()
	schema := &Schema{
		Collection: ns.Collection,
		Type:       mira.NewType(s),
		Cols:       []*Col{},
		Pkg:        ns.Pkg,
	}

	identCnt := 0
	for _, column := range ns.Columns {
		cfg := column.Cfg()
		col := &Col{
			Name:        cfg.Name,
			StructField: strcase.ToCamel(cfg.Name),
			Type:        mira.NewType(cfg.T),
			Ident:       cfg.Ident,
			Auto:        cfg.Auto,
			Nullable:    cfg.Nullable,
		}

		if len(cfg.StructField) > 0 {
			col.StructField = cfg.StructField
		}

		if cfg.Ident {
			schema.Ident = col
			identCnt++
		}

		schema.Cols = append(schema.Cols, col)
	}

	if identCnt == 0 {
		return nil, errors.New("an identity column is required")
	}

	if identCnt > 1 {
		return nil, errors.New("only one identity column is allowed")
	}

	return schema, nil
}
