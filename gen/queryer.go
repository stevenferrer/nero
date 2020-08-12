package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
)

func newQueryer(schema *gen.Schema) *jen.Statement {
	queryerDoc := fmt.Sprintf("Query is the query builder for %s", schema.Type.Name())
	stmnt := jen.Comment(queryerDoc).Line().
		Type().Id("Queryer").Struct(
		jen.Id("limit").Uint64(),
		jen.Id("offset").Uint64(),
		jen.Id("pfs").Op("[]").Id("PredFunc"),
		jen.Id("sfs").Op("[]").Id("SortFunc"),
	).Line()

	factoryDoc := "NewQueryer returns a query builder"
	stmnt = stmnt.Comment(factoryDoc).Line().
		Func().Id("NewQueryer").Params().
		Params(jen.Op("*").Id("Queryer")).Block(
		jen.Return(jen.Op("&").Id("Queryer").Block())).
		Line().Line()

	rcvrParamsC := jen.Id("q").Op("*").Id("Queryer")
	retParamsC := jen.Op("*").Id("Queryer")
	retC := jen.Return(jen.Id("q"))

	// where
	whereDoc := "Where adds predicates to the query"
	stmnt = stmnt.Comment(whereDoc).Line().
		Func().Params(rcvrParamsC).Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(retParamsC).
		Block(jen.Id("q").Dot("pfs").Op("=").
			Append(
				jen.Id("q").Dot("pfs"),
				jen.Id("pfs").Op("..."),
			), retC).Line().Line()

	sortDoc := "Sort adds sort/order to the query"
	stmnt = stmnt.Comment(sortDoc).Line().
		Func().Params(rcvrParamsC).Id("Sort").
		Params(jen.Id("sfs").Op("...").Id("SortFunc")).
		Params(retParamsC).
		Block(jen.Id("q").Dot("sfs").Op("=").
			Append(
				jen.Id("q").Dot("sfs"),
				jen.Id("sfs").Op("..."),
			), retC).Line().Line()

	limitDoc := "Limit adds limit to the query"
	stmnt = stmnt.Comment(limitDoc).Line().
		Func().Params(rcvrParamsC).
		Id("Limit").Params(jen.Id("limit").Uint64()).
		Params(retParamsC).
		Block(
			jen.Id("q").Dot("limit").Op("=").Id("limit"),
			retC,
		).Line().Line()

	offsetDoc := "Offset adds offset to the query"
	stmnt = stmnt.Comment(offsetDoc).Line().
		Func().Params(rcvrParamsC).
		Id("Offset").Params(jen.Id("offset").Uint64()).
		Params(retParamsC).
		Block(
			jen.Id("q").Dot("offset").Op("=").Id("offset"),
			retC,
		).Line().Line()

	return stmnt
}
