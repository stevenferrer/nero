package sqlite

import (
	"reflect"

	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/predicate"
)

const (
	pkgPath = "github.com/sf9v/nero"
	sqPkg   = "github.com/Masterminds/squirrel"
	errPkg  = "github.com/pkg/errors"
)

// NewSQLiteRepoC generates an sqlite repository implementation
func NewSQLiteRepoC(schema *gen.Schema) *jen.Statement {
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
	ctxC := jen.Qual("context", "Context")
	ctxIDC := jen.Id("ctx")
	txC := jen.Qual(pkgPath, "Tx")

	// tx method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Tx").
		Params(jen.Id("ctx").Add(ctxC)).
		Params(txC, jen.Error()).Block(
		jen.Return(jen.Id("s").Dot("db").Dot("BeginTx").
			Call(ctxIDC, jen.Nil()))).
		Line().Line()

	txCommit := jen.Id("tx").Dot("Commit").Call()
	txRollback := jen.Id("rollback").Call(jen.Id("tx"), jen.Err())

	// create method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Create").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("c").Op("*").Id("Creator"),
		).
		Params(gen.GetTypeC(ident.Typ), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Id("s").Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(gen.GetZeroValC(ident.Typ), jen.Err()))

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).Op(":=").
				Id("s").Dot("CreateTx").Call(
				ctxIDC,
				jen.Id("tx"),
				jen.Id("c"),
			)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(gen.GetZeroValC(ident.Typ), txRollback),
			).Line()

			g.Return(
				jen.Id(ident.LowerCamelName()),
				txCommit,
			)
		}).Line().Line()

	// query method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Query").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(
			jen.Op("[]").Op("*").
				Qual(schema.Typ.PkgPath, schema.Typ.Name),
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Id("s").Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err())).Line()

			g.List(jen.Id("list"), jen.Err()).Op(":=").
				Id("s").Dot("QueryTx").Call(ctxIDC, jen.Id("tx"), jen.Id("q"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), txRollback),
			).Line()

			g.Return(
				jen.Id("list"),
				txCommit,
			)
		}).Line().Line()

	// update method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Update").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("u").Op("*").Id("Updater"),
		).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Id("s").Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err())).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("s").Dot("UpdateTx").Call(
				ctxIDC, jen.Id("tx"), jen.Id("u"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), txRollback),
			).Line()

			g.Return(jen.Id("rowsAffected"), txCommit)
		}).Line().Line()

	// delete method
	stmnt = stmnt.Func().Params(rcvrParam).Id("Delete").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("d").Op("*").Id("Deleter"),
		).
		Params(jen.Int64(), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("tx"), jen.Err()).Op(":=").
				Id("s").Dot("Tx").Call(ctxIDC)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err())).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("s").Dot("DeleteTx").Call(
				ctxIDC,
				jen.Id("tx"),
				jen.Id("d"),
			)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), txRollback),
			).Line()

			g.Return(jen.Id("rowsAffected"), txCommit)
		}).Line().Line()

	// create tx method
	stmnt = stmnt.Func().Params(rcvrParam).Id("CreateTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("c").Op("*").Id("Creator"),
		).
		Params(gen.GetTypeC(ident.Typ), jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(
				jen.Return(
					gen.GetZeroValC(ident.Typ),
					jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
				),
			).Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(gen.GetZeroValC(ident.Typ), jen.Err()))

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
						g.Id("c").Dot(col.LowerCamelName())
					}
				}).Op(".").Line().
				Id("RunWith").Call(jen.Id("txx"))
			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id(ident.LowerCamelName()), jen.Err()).Op(":=").
				Id("res").Dot("LastInsertId").Call()
			g.Add(ifErr).Line()

			if ident.Typ.T.Kind() == reflect.String {
				g.Return(
					jen.Qual("strconv", "FormatInt").
						Call(jen.Id(ident.Name), jen.Lit(10)),
					jen.Nil(),
				)
				return
			}

			g.Return(jen.Id(ident.LowerCamelName()), jen.Nil())
		}).Line().Line()

	// query tx method
	queryRetTyp := jen.Op("[]").Op("*").
		Qual(schema.Typ.PkgPath, schema.Typ.Name)
	stmnt = stmnt.Func().Params(rcvrParam).Id("QueryTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("q").Op("*").Id("Queryer"),
		).
		Params(queryRetTyp, jen.Error()).
		BlockFunc(func(g *jen.Group) {
			// assert tx type
			g.List(jen.Id("txx"), jen.Id("ok")).Op(":=").
				Id("tx").Assert(jen.Op("*").Qual("database/sql", "Tx"))
			g.If(jen.Op("!").Id("ok")).Block(
				jen.Return(
					jen.Nil(),
					jen.Qual(errPkg, "New").Call(jen.Lit("expecting tx to be *sql.Tx")),
				),
			).Line()

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("q").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Select").
				Call(jen.Id("q").Dot("columns").Op("...")).Op(".").Line().
				Id("From").Call(jen.Id("q").Dot("collection")).Op(".").Line().
				Id("RunWith").Call(jen.Id("txx"))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("All").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range predicate.Ops {
							opStr := string(op)
							g.Case(jen.Qual(pkgPath+"/predicate", opStr)).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
									Call(jen.Qual(sqPkg, opStr).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			// limit
			g.If(jen.Id("q").Dot("limit").Op(">").Lit(0)).Block(
				jen.Id("qb").Op("=").Id("qb").Dot("Limit").Call(
					jen.Id("q").Dot("limit"),
				),
			).Line()

			// offset
			g.If(jen.Id("q").Dot("offset").Op(">").Lit(0)).Block(
				jen.Id("qb").Op("=").Id("qb").Dot("Offset").Call(
					jen.Id("q").Dot("offset"),
				),
			).Line()

			g.List(jen.Id("rows"), jen.Err()).Op(":=").
				Id("qb").Dot("QueryContext").Call(ctxIDC)
			g.Add(ifErr)
			g.Defer().Id("rows").Dot("Close").Call().Line()

			g.Id("list").Op(":=").Add(queryRetTyp).Block()
			g.For(jen.Id("rows").Dot("Next").Call()).BlockFunc(func(g *jen.Group) {
				g.Var().Id("item").Qual(schema.Typ.PkgPath, schema.Typ.Name)
				g.Err().Op("=").Id("rows").Dot("Scan").CallFunc(func(g *jen.Group) {
					for _, col := range schema.Cols {
						g.Line().Op("&").Id("item").Dot(col.Field)
					}
					g.Line()
				})
				g.Add(ifErr).Line()

				g.Id("list").Op("=").Append(jen.Id("list"), jen.Op("&").Id("item"))
			}).Line()

			g.Return(jen.Id("list"), jen.Nil())
		}).Line().Line()

	// update tx method
	stmnt = stmnt.Func().Params(rcvrParam).Id("UpdateTx").
		Params(
			jen.Id("ctx").Add(ctxC),
			jen.Id("tx").Add(txC),
			jen.Id("u").Op("*").Id("Updater"),
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

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("u").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Update").
				Call(jen.Id("u").Dot("collection")).Op(".").Line().
				Do(func(s *jen.Statement) {
					colCnt := 0
					for _, col := range schema.Cols {
						if col.Auto {
							continue
						}
						colCnt++
					}

					for i, col := range schema.Cols {
						if col.Auto {
							continue
						}
						s.Id("Set").Call(jen.Lit(col.Name),
							jen.Id("u").Dot(col.LowerCamelName()))
						// add dot
						if i < colCnt {
							s.Op(".").Line()
						}
					}
				}).Op(".").Line().
				Id("RunWith").Call(jen.Id("txx"))

			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("All").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range predicate.Ops {
							opStr := string(op)
							g.Case(jen.Qual(pkgPath+"/predicate", opStr)).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
									Call(jen.Qual(sqPkg, opStr).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		}).Line().Line()

	// delete tx method
	stmnt = stmnt.Func().Params(rcvrParam).Id("DeleteTx").
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

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Lit(0), jen.Err()))

			// predicates
			g.Id("pb").Op(":=").Op("&").
				Qual(pkgPath+"/predicate", "Predicates").Block()
			g.For(jen.List(jen.Id("_"), jen.Id("pf")).
				Op(":=").Range().Id("d").Dot("pfs")).
				Block(jen.Id("pf").Call(jen.Id("pb"))).
				Line()

			// query builder
			g.Id("qb").Op(":=").Qual(sqPkg, "Delete").
				Call(jen.Id("d").Dot("collection")).Op(".").Line().
				Id("RunWith").Call(jen.Id("txx"))
			g.For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
				Range().Id("pb").Dot("All").Call())).
				Block(jen.Switch(jen.Id("p").Dot("Op")).
					BlockFunc(func(g *jen.Group) {
						for _, op := range predicate.Ops {
							opStr := string(op)
							g.Case(jen.Qual(pkgPath+"/predicate", opStr)).
								Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
									Call(jen.Qual(sqPkg, opStr).Block(
										jen.Id("p").Dot("Field").Op(":").
											Id("p").Dot("Val").Op(",")),
									),
								)
						}
					}),
				).Line()

			g.List(jen.Id("res"), jen.Err()).Op(":=").
				Id("qb").Dot("ExecContext").Call(ctxIDC)
			g.Add(ifErr).Line()

			g.List(jen.Id("rowsAffected"), jen.Err()).Op(":=").
				Id("res").Dot("RowsAffected").Call()
			g.Add(ifErr).Line()

			g.Return(jen.Id("rowsAffected"), jen.Nil())
		})

	return stmnt
}
