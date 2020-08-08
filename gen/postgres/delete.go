package postgres

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/predicate"
)

func newDeleteBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Delete").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("d").Op("*").Id("Deleter"),
		).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err())).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("DeleteTx").
				Call(ctxIDC, jen.Id("tx"), jen.Id("d"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), txRollbackC),
			).Line()

			g.Return(jen.Id("rowsAffected"), txCommitC)
		})

}

func newDeleteTxBlock(schema *gen.Schema) *jen.Statement {
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

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("d").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Delete").
				Call(jen.Lit(fmt.Sprintf("%q", schema.Collection))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).
				Op(".").Line().Id("RunWith").
				Call(jen.Id("txx"))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("All").Call())).
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
						}))

			// debug
			g.Add(newDebugLogBlock("Delete")).Line().Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		})
}
