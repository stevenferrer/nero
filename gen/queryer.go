package gen

import (
	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
)

func newQueryer(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("Queryer").Struct(
		jen.Id("limit").Uint64(),
		jen.Id("offset").Uint64(),
		jen.Id("pfs").Op("[]").Id("PredFunc"),
		jen.Id("sfs").Op("[]").Id("SortFunc"),
	).Line()

	// factory
	stmnt = stmnt.Func().Id("NewQueryer").Params().
		Params(jen.Op("*").Id("Queryer")).Block(
		jen.Return(jen.Op("&").Id("Queryer").Block())).
		Line().Line()

	rcvrParamsC := jen.Id("q").Op("*").Id("Queryer")
	retParamsC := jen.Op("*").Id("Queryer")
	retC := jen.Return(jen.Id("q"))

	// where
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(retParamsC).
		Block(jen.Id("q").Dot("pfs").Op("=").
			Append(
				jen.Id("q").Dot("pfs"),
				jen.Id("pfs").Op("..."),
			), retC).Line().Line()

	// sort
	stmnt = stmnt.Func().Params(rcvrParamsC).Id("Sort").
		Params(jen.Id("sfs").Op("...").Id("SortFunc")).
		Params(retParamsC).
		Block(jen.Id("q").Dot("sfs").Op("=").
			Append(
				jen.Id("q").Dot("sfs"),
				jen.Id("sfs").Op("..."),
			), retC).Line().Line()

	// limit
	stmnt = stmnt.Func().Params(rcvrParamsC).
		Id("Limit").Params(jen.Id("limit").Uint64()).
		Params(retParamsC).
		Block(
			jen.Id("q").Dot("limit").Op("=").Id("limit"),
			retC,
		).Line().Line()

	// offset
	stmnt = stmnt.Func().Params(rcvrParamsC).
		Id("Offset").Params(jen.Id("offset").Uint64()).
		Params(retParamsC).
		Block(
			jen.Id("q").Dot("offset").Op("=").Id("offset"),
			retC,
		).Line().Line()

	return stmnt
}
