package postgres

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/aggregate"
	"github.com/sf9v/nero/predicate"
	"github.com/sf9v/nero/sort"
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

func newAggregateTxBlock() *jen.Statement {
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

			// quote table name
			g.Id("table").Op(":=").Qual("fmt", "Sprintf").
				Call(jen.Lit("%q"), jen.Id("a").Dot("collection"))

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("cols").Op("...")).
				Dot("From").Call(jen.Id("table")).
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

			// predicates
			g.Id("preds").Op(":=").Op("&").
				Qual(predPkg, "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("a").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("preds")))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("preds").Dot("All").Call())).
				Block(
					// switch block
					jen.Switch(jen.Id("p").Dot("Op")).
						BlockFunc(func(g *jen.Group) {
							for _, op := range predOps {
								var oprtr = "="
								switch op {
								case predicate.Eq:
									oprtr = "="
								case predicate.NotEq:
									oprtr = "<>"
								case predicate.Gt:
									oprtr = ">"
								case predicate.GtOrEq:
									oprtr = ">="
								case predicate.Lt:
									oprtr = "<"
								case predicate.LtOrEq:
									oprtr = "<="
								}

								g.Case(jen.Qual(pkgPath+"/predicate", op.String())).
									Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
										Call(
											jen.Qual("fmt", "Sprintf").Call(
												jen.Lit("%q "+oprtr+" ?"),
												jen.Id("p").Dot("Col"),
											),
											jen.Id("p").Dot("Val"),
										))
							}
						})).Line()

			// sorts
			g.Id("sorts").Op(":=").Op("&").
				Qual(sortPkg, "Sorts").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("sf")).
				Op(":=").Range().Id("a").Dot("sfs")).
				Block(jen.Id("sf").Call(jen.Id("sorts")))
			g.For(jen.List(jen.Id("_"), jen.Id("s").Op(":=").
				Range().Id("sorts").Dot("All").Call())).
				Block(
					jen.Id("col").Op(":=").Qual("fmt", "Sprintf").
						Call(jen.Lit("%q"), jen.Id("s").Dot("Col")),
					jen.Switch(jen.Id("s").Dot("Direction")).
						BlockFunc(func(g *jen.Group) {
							// ascending
							g.Case(jen.Qual(sortPkg, sort.Asc.String())).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("OrderBy").
									Call(jen.Id("col").Op("+").Lit(" ASC")))
							// descending
							g.Case(jen.Qual(sortPkg, sort.Desc.String())).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("OrderBy").
									Call(jen.Id("col").Op("+").Lit(" DESC")))
						})).Line()

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
			g.Id("dv").Op(":=").Qual("reflect", "ValueOf").
				Call(jen.Id("a").Dot("dest")).Dot("Elem").Call()
			g.Id("dt").Op(":=").Qual("reflect", "TypeOf").
				Call(jen.Id("dv").Dot("Interface").Call()).
				Dot("Elem").Call()
			// TODO: add more details to error message
			errMsg := "aggregate columns and destination struct field count should match"
			g.If(jen.Id("dt").Dot("NumField").Call().
				Op("!=").Len(jen.Id("cols"))).
				Block(jen.Return(jen.Qual(errPkg, "New").
					Call(jen.Lit(errMsg)))).Line()

			g.For(jen.Id("rows").Dot("Next").Call()).BlockFunc(func(g *jen.Group) {
				g.Id("de").Op(":=").Qual("reflect", "New").
					Call(jen.Id("dt")).Dot("Elem").Call()
				g.Id("dest").Op(":=").Make(
					jen.Op("[]").Interface(),
					jen.Id("de").Dot("NumField").Call(),
				)

				g.For(jen.Id("i").Op(":=").Lit(0),
					jen.Id("i").Op("<").Id("de").Dot("NumField").Call(),
					jen.Id("i").Op("++"),
				).Block(
					jen.Id("dest").Index(jen.Id("i")).Op("=").
						Id("de").Dot("Field").Call(jen.Id("i")).
						Dot("Addr").Call().Dot("Interface").Call(),
				).Line()

				g.Err().Op("=").Id("rows").Dot("Scan").
					Call(jen.Id("dest").Op("..."))
				g.Add(ifErr).Line()

				g.Id("dv").Dot("Set").Call(jen.Qual("reflect", "Append").
					Call(jen.Id("dv"), jen.Id("de")))
			}).Line()

			g.Return(jen.Nil())
		})
}
