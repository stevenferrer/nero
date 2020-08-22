package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	jenx "github.com/sf9v/nero/x/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newCreator(schema *gen.Schema) *jen.Statement {
	creatorDoc := fmt.Sprintf("Creator is the create builder for %s", schema.Type.Name())
	stmnt := jen.Comment(creatorDoc).Line().
		Type().Id("Creator").
		StructFunc(func(g *jen.Group) {
			for _, col := range schema.Cols {
				if col.Auto {
					continue
				}
				colField := col.CamelName()
				if len(col.StructField) > 0 {
					colField = col.StructField
				}
				colField = lowCamel(colField)
				g.Id(colField).Add(jenx.Type(col.Type.V()))
			}
		}).Line()

	factoryDoc := "NewCreator returns a create builder"
	stmnt = stmnt.Comment(factoryDoc).Line().
		Func().Id("NewCreator").Params().
		Params(jen.Op("*").Id("Creator")).Block(
		jen.Return(jen.Op("&").Id("Creator").Block())).
		Line().Line()

	rcvrParamsC := jen.Id("c").Op("*").Id("Creator")

	// methods
	for _, col := range schema.Cols {
		if col.Auto {
			continue
		}
		methodID := col.CamelName()
		if len(col.StructField) > 0 {
			methodID = col.StructField
		}

		paramID := lowCamel(methodID)
		methodDoc := fmt.Sprintf("%s is the setter for %s", methodID, paramID)
		stmnt = stmnt.Comment(methodDoc).Line().
			Func().Params(rcvrParamsC).Id(methodID).
			Params(jen.Id(paramID).Add(jenx.Type(col.Type.V()))).
			Params(jen.Op("*").Id("Creator")).
			Block(
				jen.Id("c").Dot(paramID).
					Op("=").Id(paramID),
				jen.Return(jen.Id("c")),
			).Line().Line()
	}

	return stmnt
}
