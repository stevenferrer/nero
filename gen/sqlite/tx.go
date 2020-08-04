package sqlite

import "github.com/dave/jennifer/jen"

func newTxBlock() *jen.Statement {
	return jen.Func().Params(rcvrParamC).Id("Tx").
		Params(jen.Id("ctx").Add(ctxC)).
		Params(txC, jen.Error()).Block(
		jen.Return(jen.Add(rcvrIDC).Dot("db").Dot("BeginTx").
			Call(ctxIDC, jen.Nil()))).
		Line().Line()
}
