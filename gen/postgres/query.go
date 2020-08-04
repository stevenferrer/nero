package postgres

import (
	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/sort"
)

func newQueryBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Query").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(
			jen.Op("[]").Op("*").Qual(
				schema.Type.PkgPath(),
				schema.Type.Name(),
			),
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err())).Line()

			g.List(jen.Id("list"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("QueryTx").Call(ctxIDC, jen.Id("tx"), jen.Id("q"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), txRollbackC),
			).Line()

			g.Return(
				jen.Id("list"),
				txCommitC,
			)
		})
}

func newQueryOneBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("QueryOne").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(
			jen.Op("*").Qual(schema.Type.PkgPath(), schema.Type.Name()),
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err())).Line()

			g.List(jen.Id("item"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("QueryOneTx").Call(ctxIDC, jen.Id("tx"), jen.Id("q"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), txRollbackC),
			).Line()

			g.Return(jen.Id("item"), txCommitC)
		})
}

func newQueryTxBlock(schema *gen.Schema) *jen.Statement {
	retTypeC := jen.Op("[]").Op("*").
		Qual(schema.Type.PkgPath(), schema.Type.Name())
	return jen.Func().Params(rcvrParamC).Id("QueryTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(retTypeC, jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(jen.Return(jen.Nil(),
				jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")))).
				Line()

			g.Id("qb").Op(":=").Add(rcvrIDC).Dot("buildSelect").Call(jen.Id("q"))

			// debug
			g.Add(newDebugLogBlock("Query")).Line().Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()))

			g.List(jen.Id("rows"), jen.Err()).Op(":=").
				Id("qb").Dot("RunWith").Call(jen.Id("txx")).
				Dot("QueryContext").Call(ctxIDC)
			g.Add(ifErr)
			g.Defer().Id("rows").Dot("Close").Call().Line()

			g.Id("list").Op(":=").Add(retTypeC).Block()
			g.For(jen.Id("rows").Dot("Next").Call()).BlockFunc(func(g *jen.Group) {
				g.Var().Id("item").Qual(schema.Type.PkgPath(), schema.Type.Name())
				g.Err().Op("=").Id("rows").Dot("Scan").CallFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						g.Line().Op("&").Id("item").Dot(col.StructField)
					}
					g.Line()
				})
				g.Add(ifErr).Line()

				g.Id("list").Op("=").Append(jen.Id("list"), jen.Op("&").Id("item"))
			}).Line()

			g.Return(jen.Id("list"), jen.Nil())
		})

}

func newQueryOneTxBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("QueryOneTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(
			jen.Op("*").Qual(schema.Type.PkgPath(), schema.Type.Name()),
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(
				jen.Return(jen.Nil(), jen.Qual(errPkg, "New").
					Call(jen.Lit("expecting tx to be *sql.Tx")))).Line()

			g.Id("qb").Op(":=").Add(rcvrIDC).Dot("buildSelect").Call(jen.Id("q"))

			// debug
			g.Add(newDebugLogBlock("One")).Line().Line()

			g.Var().Id("item").Qual(schema.Type.PkgPath(), schema.Type.Name())
			g.Err().Op(":=").Id("qb").Dot("RunWith").
				Call(jen.Id("txx")).Op(".").Line().Id("QueryRowContext").
				Call(ctxIDC).Op(".").Line().
				Id("Scan").CallFunc(func(g *jen.Group) {
				for _, col := range schema.Cols {
					g.Line().Op("&").Id("item").Dot(col.StructField)
				}
				g.Line()
			})
			g.If(jen.Err().Op("!=").Nil()).
				Block(jen.Return(jen.Nil(), jen.Err())).Line()

			g.Return(
				jen.Op("&").Id("item"),
				jen.Nil(),
			)
		})
}

func newSelectBuilderBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("buildSelect").
		Params(jen.Id("q").Op("*").Id("Queryer")).
		Params(jen.Qual(sqPkg, "SelectBuilder")).
		BlockFunc(func(g *jen.Group) {
			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("q").Dot("columns").Op("...")).
				Op(".").Line().Id("From").
				Call(jen.Id("q").Dot("collection")).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).Line()

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("q").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb")))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("All").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range predOps {
							opStr := op.String()
							g.Case(jen.Qual(pkgPath+"/predicate", opStr)).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
									Call(jen.Qual(sqPkg, opStr).
										Block(
											jen.Id("p").Dot("Col").Op(":").
												Id("p").Dot("Val").Op(",")),
									))
						}
					})).Line()

			// sorts
			g.Id("sorts").Op(":=").Op("&").
				Qual(sortPkg, "Sorts").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("sf")).
				Op(":=").Range().Id("q").Dot("sfs")).
				Block(jen.Id("sf").Call(jen.Id("sorts")))
			g.For(jen.List(jen.Id("_"), jen.Id("s").Op(":=").
				Range().Id("sorts").Dot("All").Call())).
				Block(
					jen.Id("col").Op(":=").Id("s").Dot("Col"),
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

			// limit
			g.If(jen.Id("q").Dot("limit").Op(">").Lit(0)).Block(
				jen.Id("qb").Op("=").Id("qb").Dot("Limit").Call(
					jen.Id("q").Dot("limit"),
				)).Line()

			// offset
			g.If(jen.Id("q").Dot("offset").Op(">").Lit(0)).
				Block(jen.Id("qb").Op("=").Id("qb").
					Dot("Offset").Call(jen.Id("q").Dot("offset")),
				).Line()

			g.Return(jen.Id("qb"))
		})
}
