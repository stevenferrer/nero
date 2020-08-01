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
	for _, co := range ns.Columns {
		col := &Col{
			Name:      co.Name,
			FieldName: strcase.ToCamel(co.Name),
			Typ:       NewTyp(co.T),
			Ident:     co.IsIdent,
			Auto:      co.IsAuto,
		}

		if len(co.FieldName) > 0 {
			col.FieldName = co.FieldName
		}

		if co.IsIdent {
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
