package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newRepository(schema *gen.Schema) *jen.Statement {
	typC := jen.Op("*").Qual(schema.Typ.PkgPath, schema.Typ.Name)
	comment := fmt.Sprintf("Repository is the contract for storing %s",
		schema.Typ.Name)
	ctxC := jen.Qual("context", "Context")
	txC := jen.Id("Tx")
	return jen.Comment(comment).Line().
		Type().Id("Repository").Interface(
		jen.Id("Tx").
			Params(ctxC).Params(jen.Id("Tx"), jen.Error()),
		jen.Id("Create").
			Params(ctxC, jen.Op("*").Id("Creator")).
			Params(
				jen.Id("id").Add(gen.GetTypeC(schema.Ident.Typ)),
				jen.Err().Error(),
			),
		jen.Id("Query").
			Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(typC), jen.Error()),
		jen.Id("Update").
			Params(ctxC, jen.Op("*").Id("Updater")).
			Params(jen.Id("rowsAffected").Int64(),
				jen.Id("err").Error()),
		jen.Id("Delete").
			Params(ctxC, jen.Op("*").Id("Deleter")).
			Params(jen.Id("rowsAffected").Int64(),
				jen.Id("err").Error()),
		jen.Id("CreateTx").
			Params(ctxC, txC, jen.Op("*").Id("Creator")).
			Params(
				jen.Id("id").Add(gen.GetTypeC(schema.Ident.Typ)),
				jen.Err().Error(),
			),
		jen.Id("QueryTx").
			Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(typC), jen.Error()),
		jen.Id("UpdateTx").
			Params(ctxC, txC, jen.Op("*").Id("Updater")).
			Params(jen.Id("rowsAffected").Int64(),
				jen.Id("err").Error()),
		jen.Id("DeleteTx").
			Params(ctxC, txC, jen.Op("*").Id("Deleter")).
			Params(jen.Id("rowsAffected").Int64(),
				jen.Id("err").Error()),
	)
}
