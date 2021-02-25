package internal

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/template"
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
	Templates     []nero.Templater
}

// BuildSchema builds schema from a nero.Schemaer to Schema
func BuildSchema(schemaer nero.Schemaer) (*Schema, error) {
	schema := schemaer.Schema()
	st := mira.NewType(schemaer)

	tmpls := schema.Templates
	if len(tmpls) == 0 {
		// default templates
		tmpls = []nero.Templater{
			template.NewPostgresTemplate(),
		}
	}
	internalSchema := &Schema{
		Pkg:           strings.ToLower(schema.PkgName),
		Collection:    schema.Collection,
		Type:          st,
		Cols:          []*Col{},
		SchemaImports: []string{st.PkgPath()},
		Templates:     tmpls,
	}

	colImportMap := map[string]int{}

	identCnt := 0
	for _, c := range schema.Columns {
		col := &Col{
			Name:             c.Name,
			StructField:      stringsx.ToCamel(c.Name),
			Type:             mira.NewType(c.T),
			Ident:            c.Identity,
			Auto:             c.Auto,
			Optional:         c.Optional,
			ColumnComparable: c.ColumnComparable,
		}

		if len(c.StructField) > 0 {
			col.StructField = c.StructField
		}

		if c.Identity {
			internalSchema.Ident = col
			identCnt++
		}

		if col.Type.PkgPath() != "" {
			colImportMap[col.Type.PkgPath()] = 1
		}

		internalSchema.Cols = append(internalSchema.Cols, col)
	}

	columnImports := []string{}
	for k := range colImportMap {
		columnImports = append(columnImports, k)
	}

	internalSchema.ColumnImports = columnImports

	if identCnt == 0 {
		return nil, errors.New("an identity column is required")
	}

	if identCnt > 1 {
		return nil, errors.New("only one identity column is allowed")
	}

	return internalSchema, nil
}
