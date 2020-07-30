package gen

import "github.com/dave/jennifer/jen"

func newDeleter() *jen.Statement {
	stmnt := jen.Type().Id("Deleter").Struct(
		jen.Id("collection").String(),
		jen.Id("pfs").Op("[]").Id("PredicateFunc"),
	).Line()

	// factory
	stmnt = stmnt.Func().Id("NewDeleter").Params().
		Params(jen.Op("*").Id("Deleter")).Block(
		jen.Return(jen.Op("&").Id("Deleter").Block(
			jen.Id("collection").Op(":").
				Id("collection").Op(","),
		)),
	).Line().Line()

	// where
	stmnt = stmnt.Func().
		Params(jen.Id("d").Op("*").Id("Deleter")).
		Id("Where").
		Params(jen.Id("pfs").Op("...").Id("PredicateFunc")).
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
