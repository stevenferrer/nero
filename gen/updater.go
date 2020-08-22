package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	jenx "github.com/sf9v/nero/x/jen"
)

func newUpdater(schema *gen.Schema) *jen.Statement {
	updaterDoc := fmt.Sprintf("Updater is the update builder for %s", schema.Type.Name())
	stmnt := jen.Comment(updaterDoc).Line().
		Type().Id("Updater").StructFunc(func(g *jen.Group) {
		for _, col := range schema.Cols {
			if col.Auto {
				continue
			}
			field := col.LowerCamelName()
			if len(col.StructField) > 0 {
				field = lowCamel(col.StructField)
			}
			colv := col.Type.V()
			g.Id(field).Add(jenx.Type(colv))
		}

		g.Id("pfs").Op("[]").Id("PredFunc")
	}).Line()

	factoryDoc := "NewUpdater returns an update builder"
	stmnt = stmnt.Comment(factoryDoc).Line().
		Func().Id("NewUpdater").Params().
		Params(jen.Op("*").Id("Updater")).Block(
		jen.Return(jen.Op("&").Id("Updater").Block())).
		Line().Line()

	rcvrParamsC := jen.Id("u").Op("*").Id("Updater")
	retParamsC := jen.Op("*").Id("Updater")
	retIDC := jen.Return(jen.Id("u"))

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
			Func().Params(rcvrParamsC).Id(camel(methodID)).
			Params(jen.Id(paramID).Add(jenx.Type(col.Type.V()))).
			Params(retParamsC).
			Block(
				jen.Id("u").Dot(paramID).Op("=").Id(paramID),
				retIDC,
			).Line().Line()
	}

	// where
	whereDoc := "Where adds predicates to the query"
	stmnt = stmnt.Comment(whereDoc).Line().
		Func().Params(rcvrParamsC).Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(retParamsC).
		Block(
			jen.Id("u").Dot("pfs").Op("=").
				Append(
					jen.Id("u").Dot("pfs"),
					jen.Id("pfs").Op("..."),
				),
			retIDC,
		).Line().Line()

	return stmnt
}
