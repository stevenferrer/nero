package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

func newRepository(schema *gen.Schema) *jen.Statement {
	schemaTypeC := jen.Op("*").Qual(schema.Type.PkgPath(), schema.Type.Name())
	comment := fmt.Sprintf("Repository is the contract for storing %s",
		schema.Type.Name())
	ctxC := jen.Qual("context", "Context")
	txC := jen.Qual(pkgPath, "Tx")
	identParamC := jen.Id("id").Add(gen.GetTypeC(schema.Ident.Type))
	rowsAffectedC := jen.Id("rowsAffected").Int64()
	return jen.Comment(comment).Line().
		Type().Id("Repository").Interface(
		jen.Id("Tx").Params(ctxC).Params(txC, jen.Error()),
		jen.Id("Create").Params(ctxC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Id("CreateMany").
			Params(ctxC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Id("CreateTx").
			Params(ctxC, txC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Id("CreateManyTx").
			Params(ctxC, txC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Id("Query").
			Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Id("QueryOne").
			Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Id("QueryTx").
			Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Id("QueryOneTx").
			Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Id("Update").
			Params(ctxC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Id("UpdateTx").
			Params(ctxC, txC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Id("Delete").
			Params(ctxC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Id("DeleteTx").
			Params(ctxC, txC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Id("Aggregate").
			Params(ctxC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
		jen.Id("AggregateTx").
			Params(ctxC, txC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
	)
}
