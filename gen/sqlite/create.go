package sqlite

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
	jenx "github.com/sf9v/nero/x/jen"
	stringsx "github.com/sf9v/nero/x/strings"
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
			g.If(jen.Err().Op("!=").Nil()).Block(jen.Return(
				jenx.Zero(identv), jen.Err())).Line()

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).
				Op(":=").Add(rcvrIDC).Dot("CreateTx").
				Call(
					ctxIDC,
					jen.Id("tx"),
					jen.Id("c"),
				)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jenx.Zero(identv), txRollbackC),
			).Line()

			g.Return(jen.Id(ident.LowerCamelName()), txCommitC)
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
			g.If(jen.Op("!").Id("ok")).Block(jen.Return(
				jenx.Zero(identv),
				jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
			)).Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jenx.Zero(identv), jen.Err()))

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Insert").
				Call(jen.Id("c").Dot("collection")).Op(".").Line().
				Id("Columns").
				Call(jen.Id("c").Dot("columns").Op("...")).Op(".").Line().
				Id("Values").
				CallFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						if col.Auto {
							continue
						}

						field := col.LowerCamelName()
						if len(col.StructField) > 0 {
							field = stringsx.ToLowerCamel(col.StructField)
						}
						g.Id("c").Dot(field)
					}
				}).Op(".").Line().
				Id("RunWith").Call(jen.Id("txx"))

			// debug
			g.Add(newDebugLogBlock("Create")).Line().Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).Op(":=").
				Id("res").Dot("LastInsertId").Call()
			g.Add(ifErr).Line()

			// string ids
			if ident.Type.T().Kind() == reflect.String {
				g.Return(
					jen.Qual("strconv", "FormatInt").
						Call(jen.Id(ident.Name), jen.Lit(10)),
					jen.Nil(),
				)
				return
			}

			g.Return(jen.Id(ident.LowerCamelName()), jen.Nil())
		})
}

func newCreateManyTxBlock(schema *gen.Schema) *jen.Statement {
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
								field = stringsx.ToLowerCamel(col.StructField)
							}
							g.Id("c").Dot(field)
						}
					})
			})

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
