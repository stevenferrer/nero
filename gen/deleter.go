package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	gen "github.com/sf9v/nero/gen/internal"
)

func newDeleter(schema *gen.Schema) *jen.Statement {
	deleterDoc := fmt.Sprintf("Deleter is the delete builder for %s", schema.Type.Name())
	stmnt := jen.Comment(deleterDoc).Line().
		Type().Id("Deleter").Struct(
		jen.Id("pfs").Op("[]").Id("PredFunc"),
	).Line()

	factoryDoc := "NewDeleter returns a delete builder"
	stmnt = stmnt.Comment(factoryDoc).Line().
		Func().Id("NewDeleter").Params().
		Params(jen.Op("*").Id("Deleter")).Block(
		jen.Return(jen.Op("&").Id("Deleter").Block())).
		Line().Line()

	whereDoc := "Where adds predicates to the query"
	stmnt = stmnt.Comment(whereDoc).Line().
		Func().Params(jen.Id("d").Op("*").Id("Deleter")).
		Id("Where").Params(jen.Id("pfs").Op("...").Id("PredFunc")).
		Params(jen.Op("*").Id("Deleter")).
		Block(
			jen.Id("d").Dot("pfs").Op("=").
				Append(
					jen.Id("d").Dot("pfs"),
					jen.Id("pfs").Op("..."),
				),
			jen.Return(jen.Id("d")),
		).Line().Line()

	return stmnt
}
