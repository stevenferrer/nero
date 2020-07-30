package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newUpdater(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("Updater").StructFunc(func(g *jen.Group) {
		g.Id("collection").String()
		g.Id("columns").Op("[]").String()

		for _, col := range schema.Cols {
			if col.Auto {
				continue
			}
			g.Id(col.LowerCamelName()).Add(gen.GetTypeC(col.Typ))
		}

		g.Id("pfs").Op("[]").Id("PredicateFunc")
	}).Line()

	// factory
	stmnt = stmnt.Func().Id("NewUpdater").Params().
		Params(jen.Op("*").Id("Updater")).Block(
		jen.Return(jen.Op("&").Id("Updater").Block(
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

	rcvrParams := jen.Id("u").Op("*").Id("Updater")
	retParams := jen.Op("*").Id("Updater")
	ret := jen.Return(jen.Id("u"))

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
			Params(retParams).
			Block(
				jen.Id("u").Dot(col.LowerCamelName()).
					Op("=").Id(col.LowerCamelName()),
				ret,
			).Line().Line()
	}

	// where
	stmnt = stmnt.Func().
		Params(rcvrParams).
		Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredicateFunc")).
		Params(retParams).
		Block(
			jen.Id("u").Dot("pfs").Op("=").
				Append(
					jen.Id("u").Dot("pfs"),
					jen.Id("pfs").Op("..."),
				),
			ret,
		).Line().Line()

	return stmnt
}
