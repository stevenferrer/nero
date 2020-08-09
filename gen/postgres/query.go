package postgres

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
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
		Block(jen.Return(jen.Id(rcvrID).Dot("query").Call(
			jen.Id("ctx"),
			jen.Id(rcvrID).Dot("db"),
			jen.Id("q"),
		)))
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

			g.Return(jen.Id(rcvrID).Dot("query").Call(
				jen.Id("ctx"),
				jen.Id("txx"),
				jen.Id("q"),
			))
		})

}

func newQueryRunnerBlock(schema *gen.Schema) *jen.Statement {
	retTypeC := jen.Op("[]").Op("*").
		Qual(schema.Type.PkgPath(), schema.Type.Name())
	return jen.Func().Params(rcvrParamC).Id("query").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("runner").Add(runnerC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(retTypeC, jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.Id("qb").Op(":=").Add(rcvrIDC).Dot("buildSelect").Call(jen.Id("q"))

			// debug
			g.Add(newDebugLogBlock("Query")).Line().Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()))

			g.List(jen.Id("rows"), jen.Err()).Op(":=").
				Id("qb").Dot("RunWith").Call(jen.Id("runner")).
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
		Block(jen.Return(jen.Id(rcvrID).Dot("queryOne").Call(
			jen.Id("ctx"),
			jen.Id(rcvrID).Dot("db"),
			jen.Id("q"),
		)))
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

			g.Return(
				jen.Id(rcvrID).Dot("queryOne").Call(
					jen.Id("ctx"),
					jen.Id("txx"),
					jen.Id("q"),
				),
			)
		})
}

func newQueryOneRunnerBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("queryOne").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("runner").Add(runnerC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(
			jen.Op("*").Qual(schema.Type.PkgPath(), schema.Type.Name()),
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			g.Id("qb").Op(":=").Add(rcvrIDC).
				Dot("buildSelect").Call(jen.Id("q"))

			// debug
			g.Add(newDebugLogBlock("QueryOne")).Line().Line()

			g.Var().Id("item").Qual(schema.Type.PkgPath(), schema.Type.Name())
			g.Err().Op(":=").Id("qb").Dot("RunWith").
				Call(jen.Id("runner")).Op(".").Line().Id("QueryRowContext").
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

func newBuildSelectBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("buildSelect").
		Params(jen.Id("q").Op("*").Id("Queryer")).
		Params(jen.Qual(sqPkg, "SelectBuilder")).
		BlockFunc(func(g *jen.Group) {

			g.Id("columns").Op(":=").Index().String().
				ValuesFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						g.Lit(fmt.Sprintf("%q", col.Name))
					}
				})

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("columns").Op("...")).
				Op(".").Line().Id("From").
				Call(jen.Lit(fmt.Sprintf("%q", schema.Collection))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).Line()

			g.Id("pfs").Op(":=").Id("q").Dot("pfs")
			g.Add(newPredicatesBlock()).Line()

			g.Id("sfs").Op(":=").Id("q").Dot("sfs")
			g.Add(newSortsBlock()).Line()

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
