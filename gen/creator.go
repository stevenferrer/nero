package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newCreator(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("Creator").
		StructFunc(func(g *jen.Group) {
			g.Id("collection").String()
			g.Id("columns").Op("[]").String()

			for _, col := range schema.Cols {
				if col.Auto {
					continue
				}
				g.Id(col.LowerCamelName()).
					Add(gen.GetTypeC(col.Type))
			}
		}).Line()

	// factory
	stmnt = stmnt.Func().Id("NewCreator").Params().
		Params(jen.Op("*").Id("Creator")).Block(
		jen.Return(jen.Op("&").Id("Creator").Block(
			jen.Id("collection").Op(":").Id("collection").Op(","),
			jen.Id("columns").Op(":").Op("[]").String().
				ValuesFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						if col.Auto {
							continue
						}
						g.Lit(col.Name)
					}
				}).Op(","),
		))).Line().Line()

	rcvrParamsC := jen.Id("c").Op("*").Id("Creator")

	// methods
	for _, col := range schema.Cols {
		if col.Auto {
			continue
		}
		colLowerCamel := col.LowerCamelName()
		stmnt = stmnt.Func().Params(rcvrParamsC).Id(col.CamelName()).
			Params(jen.Id(colLowerCamel).
				Add(gen.GetTypeC(col.Type))).
			Params(jen.Op("*").Id("Creator")).Block(
			jen.Id("c").Dot(colLowerCamel).
				Op("=").Id(colLowerCamel),
			jen.Return(jen.Id("c"))).Line().Line()
	}

	return stmnt
}
