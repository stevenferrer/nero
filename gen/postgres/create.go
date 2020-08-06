package postgres

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/jenx"
)

func newCreateBlock(schema *gen.Schema) *jen.Statement {
	ident := schema.Ident
	identv := ident.Type.V()
	return jen.Func().Params(rcvrParamC).Id("Create").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("c").Op("*").Id("Creator"),
		).
		Params(jenx.Type(identv), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jenx.Zero(identv), jen.Err())).
				Line()

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("CreateTx").Call(
				ctxIDC,
				jen.Id("tx"),
				jen.Id("c"),
			)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jenx.Zero(identv), txRollbackC),
			).Line()

			g.Return(
				jen.Id(ident.LowerCamelName()),
				txCommitC,
			)
		})
}

func newCreateManyBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("CreateMany").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("cs").Op("...").Op("*").Id("Creator"),
		).
		Params(jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Add(rcvrIDC).Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err())).Line()
			g.List(jen.Err()).Op("=").Add(rcvrIDC).Dot("CreateManyTx").
				Call(ctxIDC, jen.Id("tx"), jen.Id("cs").Op("..."))
			g.If(jen.Err().Op("!=").Nil()).
				Block(jen.Return(txRollbackC)).Line()
			g.Return(txCommitC)
		})
}

func newCreateTxBlock(schema *gen.Schema) *jen.Statement {
	ident := schema.Ident
	identv := ident.Type.V()
	return jen.Func().Params(rcvrParamC).Id("CreateTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("c").Op("*").Id("Creator"),
		).
		Params(jenx.Type(identv), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(
				jen.Return(
					jenx.Zero(identv),
					jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
				)).Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jenx.Zero(identv), jen.Err()))

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Insert").
				Call(jen.Id("c").Dot("collection")).Op(".").Line().
				Id("Columns").Call(jen.Id("c").Dot("columns").
				Op("...")).Op(".").Line().Id("Values").
				CallFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						if col.Auto {
							continue
						}
						field := col.LowerCamelName()
						if len(col.StructField) > 0 {
							field = strcase.ToLowerCamel(col.StructField)
						}
						g.Id("c").Dot(field)
					}
				}).Op(".").Line().Id("Suffix").
				Call(jen.Lit(fmt.Sprintf("RETURNING %q", ident.Name))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar")).
				Op(".").Line().Id("RunWith").Call(jen.Id("txx"))
			// debug
			g.Add(newDebugLogBlock("Create")).Line().Line()

			g.Var().Id(ident.LowerCamelName()).Add(jenx.Type(identv))
			g.Err().Op(":=").Id("qb").Dot("QueryRowContext").
				Call(ctxIDC).Dot("Scan").Call(jen.Op("&").Id(ident.LowerCamelName()))
			g.Add(ifErr).Line()

			g.Return(jen.Id(ident.LowerCamelName()), jen.Nil())
		})
}

func newCreateManyTxBlock(schema *gen.Schema) *jen.Statement {
	ident := schema.Ident
	return jen.Func().Params(rcvrParamC).Id("CreateManyTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("cs").Op("...").Op("*").Id("Creator"),
		).
		Params(jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.If(jen.Len(jen.Id("cs")).Op("==").Lit(0)).Block(
				jen.Return(jen.Nil()),
			).Line()

			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(jen.Return(
				jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
			)).Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Insert").
				Call(jen.Id("cs").Index(jen.Lit(0)).
					Dot("collection")).Op(".").Line().
				Id("Columns").
				Call(jen.Id("cs").Index(jen.Lit(0)).
					Dot("columns").Op("..."))

			g.For(jen.List(jen.Id("_"), jen.Id("c")).
				Op(":=").Range().Id("cs"),
			).BlockFunc(func(g *jen.Group) {
				g.Id("qb").Op("=").Id("qb").Dot("Values").
					CallFunc(func(g *jen.Group) {
						for _, col := range schema.Cols {
							if col.Auto {
								continue
							}
							field := col.LowerCamelName()
							if len(col.StructField) > 0 {
								field = strcase.ToLowerCamel(col.StructField)
							}
							g.Id("c").Dot(field)
						}
					})
			}).Line()

			g.Id("qb").Op("=").Id("qb").Dot("Suffix").
				Call(jen.Lit(fmt.Sprintf("RETURNING %q", ident.Name))).
				Op(".").Line().Id("PlaceholderFormat").
				Call(jen.Qual(sqPkg, "Dollar"))

			// debug
			g.Add(newDebugLogBlock("CreateMany")).Line().Line()

			g.List(jen.Id("_"), jen.Err()).Op(":=").Id("qb").
				Dot("RunWith").Call(jen.Id("txx")).
				Dot("ExecContext").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err())).Line()

			g.Return(jen.Nil())
		})
}
