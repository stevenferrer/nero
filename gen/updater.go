package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/jenx"
)

func newUpdater(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("Updater").StructFunc(func(g *jen.Group) {
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

	// factory
	stmnt = stmnt.Func().Id("NewUpdater").Params().
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
			methodID = lowCamel(col.StructField)
		}

		paramID := lowCamel(methodID)

		colv := col.Type.V()
		stmnt = stmnt.Func().Params(rcvrParamsC).Id(camel(methodID)).
			Params(jen.Id(paramID).Add(jenx.Type(colv))).
			Params(retParamsC).
			Block(
				jen.Id("u").Dot(paramID).Op("=").Id(paramID),
				retIDC,
			).Line().Line()
	}

	// where
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Where").
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
