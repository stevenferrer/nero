package postgres

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/aggregate"
	gen "github.com/sf9v/nero/gen/internal"
)

func newAggregateBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Aggregate").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("a").Op("*").Id("Aggregator"),
		).
		Params(jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).
				Block(jen.Return(jen.Err())).Line()

			g.Err().Op("=").Add(rcvrIDC).Dot("AggregateTx").
				Call(ctxIDC, jen.Id("tx"), jen.Id("a"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(txRollbackC),
			).Line()

			g.Return(txCommitC)
		})
}

func newAggregateTxBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("AggregateTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("a").Op("*").Id("Aggregator"),
		).
		Params(jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(jen.Return(
				jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
			)).Line()

			g.Id("aggs").Op(":=").Op("&").Qual(aggPkg, "Aggregates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("aggf")).
				Op(":=").Range().Id("a").Dot("aggfs"),
			).Block(jen.Id("aggf").Call(jen.Id("aggs")))

			g.Id("cols").Op(":=").Op("[]").String().Block()
			g.For(jen.List(jen.Id("_"), jen.Id("agg")).
				Op(":=").Range().Id("aggs").Dot("All").Call(),
			).Block(
				jen.Id("col").Op(":=").Id("agg").Dot("Col"),
				// quoted column
				jen.Id("qcol").Op(":=").Qual("fmt", "Sprintf").
					Call(jen.Lit("%q"), jen.Id("col")),
				jen.Switch(jen.Id("agg").Dot("Fn").
					BlockFunc(func(g *jen.Group) {
						// switch block
						for _, aggFn := range aggFns {
							if aggFn == aggregate.None {
								g.Case(jen.Qual(aggPkg, aggFn.String())).Block(
									jen.Id("cols").Op("=").Append(
										jen.Id("cols"),
										jen.Id("qcol"),
									),
								)
								continue
							}

							fnUp := strings.ToUpper(aggFn.String())
							fnLow := strings.ToLower(aggFn.String())
							g.Case(jen.Qual(aggPkg, aggFn.String())).
								Block(
									jen.Id("cols").Op("=").Append(
										jen.Id("cols"),
										jen.Lit(fnUp+"(").Op("+").Id("qcol").Op("+").
											Lit(") "+fnLow+"_").Op("+").Id("col"),
									),
								)
						}
					}))).Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("cols").Op("...")).
				Dot("From").Call(jen.Lit(fmt.Sprintf("%q", schema.Collection))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).Line()

			g.Id("groups").Op(":=").Op("[]").String().Block()
			g.For(jen.List(jen.Id("_"), jen.Id("group")).
				Op(":=").Range().Id("a").Dot("groups")).
				Block(
					jen.Id("groups").Op("=").Append(
						jen.Id("groups"),
						// quote group clause columns
						jen.Qual("fmt", "Sprintf").
							Call(jen.Lit("%q"), jen.Id("group").Dot("String").Call()),
					))

			g.Id("qb").Op("=").Id("qb").Dot("GroupBy").
				Call(jen.Id("groups").Op("...")).Line()

			g.Id("pfs").Op(":=").Id("a").Dot("pfs")
			g.Add(newPredicatesBlock()).Line()

			g.Id("sfs").Op(":=").Id("a").Dot("sfs")
			g.Add(newSortsBlock()).Line()

			// debug
			g.Add(newDebugLogBlock("Aggregate")).Line().Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err()))

			// run query
			g.List(jen.Id("rows"), jen.Err()).Op(":=").Id("qb").
				Dot("RunWith").Call(jen.Id("txx")).
				Dot("QueryContext").Call(ctxIDC)
			g.Add(ifErr)
			g.Defer().Id("rows").Dot("Close").Call().Line()

			// inspect aggregate destination
			g.Id("v").Op(":=").Qual("reflect", "ValueOf").
				Call(jen.Id("a").Dot("v")).Dot("Elem").Call()
			g.Id("t").Op(":=").Qual("reflect", "TypeOf").
				Call(jen.Id("v").Dot("Interface").Call()).
				Dot("Elem").Call()
			// TODO: add more details to error message
			errMsg := "aggregate columns and destination struct field count should match"
			g.If(jen.Id("t").Dot("NumField").Call().
				Op("!=").Len(jen.Id("cols"))).
				Block(jen.Return(jen.Qual(errPkg, "New").
					Call(jen.Lit(errMsg)))).Line()

			g.For(jen.Id("rows").Dot("Next").Call()).
				BlockFunc(func(g *jen.Group) {
					g.Id("ve").Op(":=").Qual("reflect", "New").
						Call(jen.Id("t")).Dot("Elem").Call()
					g.Id("dest").Op(":=").Make(
						jen.Op("[]").Interface(),
						jen.Id("ve").Dot("NumField").Call(),
					)

					g.For(jen.Id("i").Op(":=").Lit(0),
						jen.Id("i").Op("<").Id("ve").Dot("NumField").Call(),
						jen.Id("i").Op("++"),
					).Block(
						jen.Id("dest").Index(jen.Id("i")).Op("=").
							Id("ve").Dot("Field").Call(jen.Id("i")).
							Dot("Addr").Call().Dot("Interface").Call(),
					).Line()

					g.Err().Op("=").Id("rows").Dot("Scan").
						Call(jen.Id("dest").Op("..."))
					g.Add(ifErr).Line()

					g.Id("v").Dot("Set").Call(jen.Qual("reflect", "Append").
						Call(jen.Id("v"), jen.Id("ve")))
				}).Line()

			g.Return(jen.Nil())
		})
}
