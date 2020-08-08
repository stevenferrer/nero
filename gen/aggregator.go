package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newAggregator(schema *gen.Schema) *jen.Statement {
	typeID := "Aggregator"
	stmnt := jen.Type().Id(typeID).
		StructFunc(func(g *jen.Group) {
			g.Id("v").Interface()
			g.Id("aggfs").Op("[]").Id("AggFunc")
			g.Id("pfs").Op("[]").Id("PredFunc")
			g.Id("sfs").Op("[]").Id("SortFunc")
			g.Id("groups").Op("[]").Id("Column")
		}).Line()

	// factory
	stmnt = stmnt.Func().Id("NewAggregator").
		Params(jen.Id("v").Interface()).
		Params(jen.Op("*").Id(typeID)).Block(
		jen.Return(jen.Op("&").Id(typeID).Block(
			jen.Id("v").Op(":").Id("v").Op(","),
		))).Line().Line()

	rcvrID := "a"
	rcvrParamsC := jen.Id(rcvrID).Op("*").Id(typeID)
	retParamsC := jen.Op("*").Id(typeID)
	retC := jen.Return(jen.Id(rcvrID))

	// aggregate
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Aggregate").
		Params(jen.Id("aggfs").Op("...").Id("AggFunc")).
		Params(retParamsC).Block(jen.Id(rcvrID).Dot("aggfs").Op("=").
		Append(jen.Id(rcvrID).Dot("aggfs"), jen.Id("aggfs").Op("...")), retC).
		Line().Line()

	// where
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(retParamsC).Block(
		jen.Id(rcvrID).Dot("pfs").Op("=").
			Append(
				jen.Id(rcvrID).Dot("pfs"),
				jen.Id("pfs").Op("..."),
			), retC).Line().Line()

	// sort
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Sort").
		Params(jen.Id("sfs").Op("...").Id("SortFunc")).
		Params(retParamsC).Block(
		jen.Id(rcvrID).Dot("sfs").Op("=").
			Append(
				jen.Id(rcvrID).Dot("sfs"),
				jen.Id("sfs").Op("..."),
			), retC).Line().Line()

	// group
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Group").
		Params(jen.Id("cols").Op("...").Id("Column")).
		Params(retParamsC).Block(
		jen.Id(rcvrID).Dot("groups").Op("=").
			Append(
				jen.Id(rcvrID).Dot("groups"),
				jen.Id("cols").Op("..."),
			), retC).Line().Line()

	return stmnt
}
