package sqlite

import "github.com/dave/jennifer/jen"

func newFactoryBlock() *jen.Statement {
	return jen.Func().Id("New" + typeName).
		Params(jen.Id("db").Op("*").Qual("database/sql", "DB")).
		Params(jen.Op("*").Id(typeName)).
		Block(jen.Return(jen.Op("&").Id(typeName).Block(
			jen.Id("db").Op(":").Id("db").Op(","),
		)))
}
