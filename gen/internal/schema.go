package internal

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"

	"github.com/sf9v/nero"
)

// Schema is an internal schema
type Schema struct {
	Collection    string
	Type          *mira.Type
	Ident         *Col
	Cols          []*Col
	Pkg           string
	SchemaImports []string
	ColumnImports []string
}

// BuildSchema builds schema from a nero.Schemaer to Schema
func BuildSchema(s nero.Schemaer) (*Schema, error) {
	ns := s.Schema()
	st := mira.NewType(s)
	schema := &Schema{
		Pkg:           strings.ToLower(ns.Pkg),
		Collection:    ns.Collection,
		Type:          st,
		Cols:          []*Col{},
		SchemaImports: []string{st.PkgPath()},
	}

	colImportMap := map[string]int{}

	identCnt := 0
	for _, column := range ns.Columns {
		cfg := column.Cfg()
		col := &Col{
			Name:        cfg.Name,
			StructField: stringsx.ToCamel(cfg.Name),
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

		if col.Type.PkgPath() != "" {
			colImportMap[col.Type.PkgPath()] = 1
		}

		schema.Cols = append(schema.Cols, col)
	}

	columnImports := []string{}
	for k := range colImportMap {
		columnImports = append(columnImports, k)
	}

	schema.ColumnImports = columnImports

	if identCnt == 0 {
		return nil, errors.New("an identity column is required")
	}

	if identCnt > 1 {
		return nil, errors.New("only one identity column is allowed")
	}

	return schema, nil
}
