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
					Add(gen.GetTypeC(col.Typ))
			}
		}).Line()

	// factory
	stmnt = stmnt.Func().Id("NewCreator").Params().
		Params(jen.Op("*").Id("Creator")).Block(
		jen.Return(jen.Op("&").Id("Creator").Block(
			jen.Id("collection").Op(":").Id("collection").Op(","),
			jen.Id("columns").Op(":").Op("[]").String().ValuesFunc(func(g *jen.Group) {
				for _, col := range schema.Cols {
					if col.Auto {
						continue
					}
					g.Lit(col.Name)
				}
			}).Op(","),
		)),
	).Line().Line()

	rcvrParams := jen.Id("c").Op("*").Id("Creator")

	// methods
	for _, col := range schema.Cols {
		if col.Auto {
			continue
		}
		stmnt = stmnt.Func().
			Params(rcvrParams).
			Id(col.CamelName()).
			Params(jen.Id(col.LowerCamelName()).
				Add(gen.GetTypeC(col.Typ))).
			Params(jen.Op("*").Id("Creator")).
			Block(
				jen.Id("c").Dot(col.LowerCamelName()).
					Op("=").Id(col.LowerCamelName()),
				jen.Return(jen.Id("c")),
			).Line().Line()
	}

	return stmnt
}
