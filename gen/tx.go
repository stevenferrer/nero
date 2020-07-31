package gen

import "github.com/dave/jennifer/jen"

func newTx() *jen.Statement {
	comment := "Tx is a transaction type"
	return jen.Comment(comment).Line().
		Type().Id("Tx").Interface(
		jen.Id("Commit").Params().Params(jen.Error()),
		jen.Id("Rollback").Params().Params(jen.Error()),
	)
}
