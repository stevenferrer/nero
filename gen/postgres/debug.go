package postgres

import "github.com/dave/jennifer/jen"

func newDebugBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Debug").
		Params(jen.Id("out").Qual("io", "Writer")).
		Params(jen.Op("*").Id(typeName)).
		Block(
			jen.Id("lg").Op(":=").Qual(logPkg, "New").
				Call(jen.Id("out")).Dot("With").Call().
				Dot("Timestamp").Call().Dot("Logger").Call(),
			jen.Return(jen.Op("&").Id(typeName).Block(
				jen.Id("db").Op(":").Add(rcvrIDC).Dot("db").Op(","),
				jen.Id("log").Op(":").Op("&").Id("lg").Op(","),
			)),
		)
}
