package postgres

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
)

func newDeleteBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Delete").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("d").Op("*").Id("Deleter"),
		).
		Params(jen.Int64(), jen.Error()).
		Block(jen.Return(jen.Id(rcvrID).Dot("delete").Call(
			jen.Id("ctx"),
			jen.Id(rcvrID).Dot("db"),
			jen.Id("d"),
		)))
}

func newDeleteTxBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("DeleteTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("d").Op("*").Id("Deleter"),
		).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(
				jen.Return(
					jen.Lit(0),
					jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
				),
			).Line()

			g.Return(jen.Id(rcvrID).Dot("delete").Call(
				jen.Id("ctx"),
				jen.Id("txx"),
				jen.Id("d"),
			))
		})
}

func newDeleteRunnerBlock(schema *gen.Schema) *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("delete").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("runner").Add(runnerC),
			jen.Id("d").Op("*").Id("Deleter"),
		).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Delete").
				Call(jen.Lit(fmt.Sprintf("%q", schema.Collection))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).Line()

			g.Id("pfs").Op(":=").Id("d").Dot("pfs")
			g.Add(newPredicatesBlock()).Line()

			// debug
			g.Add(newDebugLogBlock("Delete")).Line().Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("RunWith").Call(jen.Id("runner")).
				Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		})
}
