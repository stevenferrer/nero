package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

const (
	sqPkg = "github.com/Masterminds/squirrel"
)

func newSQLiteRepo(schema *gen.Schema) *jen.Statement {
	ident := schema.Ident
	stmnt := jen.Type().Id("SQLiteRepository").
		Struct(jen.Id("db").Op("*").Qual("database/sql", "DB")).
		Line()

	stmnt = stmnt.Var().Id("_").Op("=").Id("Repository").Call(
		jen.Op("&").Id("SQLiteRepository").Block(),
	).Line()

	// factory
	stmnt = stmnt.Func().Id("NewSQLiteRepository").
		Params(jen.Id("db").Op("*").Qual("database/sql", "DB")).
		Params(jen.Op("*").Id("SQLiteRepository")).
		Block(jen.Return(jen.Op("&").Id("SQLiteRepository").Block(
			jen.Id("db").Op(":").Id("db").Op(","),
		))).Line().Line()

	rcvrParam := jen.Id("s").Op("*").Id("SQLiteRepository")

	// create method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Create").
		Params(jen.Id("c").Op("*").Id("Creator")).
		Params(gen.GetTypeC(ident.Typ), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(gen.GetZeroValC(ident.Typ), jen.Err()))

			// sql builder
			g.Id("sb").Op(":=").
				Qual(sqPkg, "Insert").
				Call(jen.Id("c").Dot("collection")).
				Op(".").Line().Id("Columns").
				Call(jen.Id("c").Dot("columns").Op("...")).
				Op(".").Line().Id("Values").CallFunc(func(g *jen.Group) {
				for _, col := range schema.Cols {
					if col.Auto {
						continue
					}
					g.Id("c").Dot(col.LowerCamelName())
				}
			}).Line()

			g.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
				Op(":=").Id("sb").Dot("ToSql").Call()
			g.Add(ifErr).Line()

			g.List(jen.Id("stmnt"), jen.Err()).Op(":=").
				Id("s").Dot("db").Dot("Prepare").Call(jen.Id("sql"))
			g.Add(ifErr).Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("stmnt").Dot("Exec").Call(jen.Id("args").Op("..."))
			g.Add(ifErr).Line()

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).Op(":=").
				Id("res").Dot("LastInsertId").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id(ident.LowerCamelName()), jen.Nil())
		}).Line().Line()

	// query method
	retTyp := jen.Op("[]").Op("*").
		Qual(schema.Typ.PkgPath, schema.Typ.Name)
	stmnt = stmnt.Func().Params(rcvrParam).Id("Query").
		Params(jen.Id("q").Op("*").Id("Queryer")).
		Params(retTyp, jen.Error()).
		BlockFunc(func(g *jen.Group) {
			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Builder").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("q").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// sql builder
			g.Id("sb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("q").Dot("columns").Op("...")).
				Dot("From").Call(jen.Id("q").Dot("collection"))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("Predicates").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range ops {
							g.Case(jen.Qual(pkgPath+"/predicate", op)).
								Block(jen.Id("sb").Op("=").Id("sb").Dot("Where").
									Call(jen.Qual(sqPkg, op).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			// limit
			g.If(jen.Id("q").Dot("limit").Op(">").Lit(0)).Block(
				jen.Id("sb").Op("=").Id("sb").Dot("Limit").Call(
					jen.Id("q").Dot("limit"),
				),
			).Line()

			// offset
			g.If(jen.Id("q").Dot("offset").Op(">").Lit(0)).Block(
				jen.Id("sb").Op("=").Id("sb").Dot("Offset").Call(
					jen.Id("q").Dot("offset"),
				),
			).Line()

			g.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
				Op(":=").Id("sb").Dot("ToSql").Call()
			g.Add(ifErr).Line()

			g.List(jen.Id("stmnt"), jen.Err()).Op(":=").
				Id("s").Dot("db").Dot("Prepare").Call(jen.Id("sql"))
			g.Add(ifErr).Line()

			g.List(jen.Id("rows"), jen.Err()).Op(":=").
				Id("stmnt").Dot("Query").Call(jen.Id("args").Id("..."))
			g.Add(ifErr)
			g.Defer().Id("rows").Dot("Close").Call().Line()

			g.Id("list").Op(":=").Add(retTyp).Block()
			g.For(jen.Id("rows").Dot("Next").Call()).BlockFunc(func(g *jen.Group) {
				g.Var().Id("item").Qual(schema.Typ.PkgPath, schema.Typ.Name)
				g.Err().Op("=").Id("rows").Dot("Scan").CallFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						g.Line().Op("&").Id("item").Dot(col.FieldName)
					}
					g.Line()
				})
				g.Add(ifErr).Line()

				g.Id("list").Op("=").Append(jen.Id("list"), jen.Op("&").Id("item"))
			}).Line()

			g.Return(jen.Id("list"), jen.Nil())
		}).Line().Line()

	// update method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Update").
		Params(jen.Id("u").Op("*").Id("Updater")).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Builder").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("u").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// sql builder

			g.Id("sb").Op(":=").Qual(sqPkg, "Update").
				Call(jen.Id("u").Dot("collection")).Op(".").Line()
			for i, col := range schema.Cols {
				if col.Auto {
					continue
				}

				s := jen.Id("Set").Call(jen.Lit(col.Name),
					jen.Id("u").Dot(col.LowerCamelName()))
				if i+i < len(schema.Cols) {
					s.Add(jen.Op("."))
				}

				g.Add(s)

			}

			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("Predicates").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range ops {
							g.Case(jen.Qual(pkgPath+"/predicate", op)).
								Block(jen.Id("sb").Op("=").Id("sb").Dot("Where").
									Call(jen.Qual(sqPkg, op).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			g.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
				Op(":=").Id("sb").Dot("ToSql").Call()
			g.Add(ifErr).Line()

			g.List(jen.Id("stmnt"), jen.Err()).Op(":=").
				Id("s").Dot("db").Dot("Prepare").Call(jen.Id("sql"))
			g.Add(ifErr).Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("stmnt").Dot("Exec").Call(jen.Id("args").Op("..."))
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		}).Line().Line()

	// delete method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Delete").
		Params(jen.Id("d").Op("*").Id("Deleter")).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Builder").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("d").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// sql builder
			g.Id("sb").Op(":=").Qual(sqPkg, "Delete").
				Call(jen.Id("d").Dot("collection"))

			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("Predicates").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range ops {
							g.Case(jen.Qual(pkgPath+"/predicate", op)).
								Block(jen.Id("sb").Op("=").Id("sb").Dot("Where").
									Call(jen.Qual(sqPkg, op).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			g.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
				Op(":=").Id("sb").Dot("ToSql").Call()
			g.Add(ifErr).Line()

			g.List(jen.Id("stmnt"), jen.Err()).Op(":=").
				Id("s").Dot("db").Dot("Prepare").Call(jen.Id("sql"))
			g.Add(ifErr).Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("stmnt").Dot("Exec").Call(jen.Id("args").Op("..."))
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		})

	return stmnt
}
