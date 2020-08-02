package gen

import "github.com/dave/jennifer/jen"

func newTx() *jen.Statement {
	// rollback function
	stmnt := jen.Func().Id("rollback").
		Params(jen.Id("tx").Qual(pkgPath, "Tx"), jen.Err().Error()).
		Params(jen.Error()).BlockFunc(func(g *jen.Group) {
		g.Id("rerr").Op(":=").Id("tx").Dot("Rollback").Call()
		g.If(jen.Id("rerr").Op("!=").Nil()).Block(
			jen.Id("err").Op("=").Qual(errPkg, "Wrapf").Call(
				jen.Err(),
				jen.Lit("rollback error: %v"),
				jen.Id("rerr"),
			),
		)
		g.Return(jen.Id("err"))
	})

	return stmnt
}
