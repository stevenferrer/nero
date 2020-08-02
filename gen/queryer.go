package gen

import (
	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
)

func newQueryer(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("Queryer").Struct(
		jen.Id("collection").String(),
		jen.Id("columns").Op("[]").String(),
		jen.Id("limit").Uint64(),
		jen.Id("offset").Uint64(),
		jen.Id("pfs").Op("[]").Id("PredFunc"),
		jen.Id("sfs").Op("[]").Id("SortFunc"),
	).Line()

	// factory
	stmnt = stmnt.Func().Id("NewQueryer").Params().
		Params(jen.Op("*").Id("Queryer")).Block(
		jen.Return(jen.Op("&").Id("Queryer").Block(
			jen.Id("collection").Op(":").Id("collection").Op(","),
			jen.Id("columns").Op(":").Op("[]").String().
				ValuesFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						g.Lit(col.Name)
					}
				}).Op(","),
		)),
	).Line().Line()

	rcvrParams := jen.Id("q").Op("*").Id("Queryer")
	retParams := jen.Op("*").Id("Queryer")
	ret := jen.Return(jen.Id("q"))

	// where
	stmnt = stmnt.Func().
		Params(rcvrParams).
		Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(retParams).
		Block(
			jen.Id("q").Dot("pfs").Op("=").
				Append(
					jen.Id("q").Dot("pfs"),
					jen.Id("pfs").Op("..."),
				),
			ret,
		).Line().Line()

	// sort
	stmnt = stmnt.Func().
		Params(rcvrParams).
		Id("Sort").
		Params(jen.Id("sfs").Op("...").Id("SortFunc")).
		Params(retParams).
		Block(
			jen.Id("q").Dot("sfs").Op("=").
				Append(
					jen.Id("q").Dot("sfs"),
					jen.Id("sfs").Op("..."),
				),
			ret,
		).Line().Line()

	// limit
	stmnt = stmnt.Func().
		Params(rcvrParams).
		Id("Limit").
		Params(jen.Id("limit").Uint64()).
		Params(retParams).
		Block(
			jen.Id("q").Dot("limit").Op("=").Id("limit"),
			ret,
		).
		Line().Line()

	// offset
	stmnt = stmnt.Func().
		Params(rcvrParams).
		Id("Offset").
		Params(jen.Id("offset").Uint64()).
		Params(retParams).
		Block(
			jen.Id("q").Dot("offset").Op("=").Id("offset"),
			ret,
		).
		Line().Line()

	return stmnt
}
