package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	jenx "github.com/sf9v/nero/x/jen"
)

func newRepository(schema *gen.Schema) *jen.Statement {
	schemaTypeC := jen.Op("*").Qual(schema.Type.PkgPath(), schema.Type.Name())
	comment := fmt.Sprintf("Repository is a %s repository", schema.Type.Name())
	ctxC := jen.Qual("context", "Context")
	txC := jen.Qual(pkgPath, "Tx")
	identv := schema.Ident.Type.V()
	identParamC := jen.Id("id").Add(jenx.Type(identv))
	rowsAffectedC := jen.Id("rowsAffected").Int64()
	name := schema.Type.Name()
	return jen.Comment(comment).Line().
		Type().Id("Repository").Interface(
		jen.Comment("Tx returns a new transaction").
			Line().Id("Tx").Params(ctxC).Params(txC, jen.Error()),
		jen.Comment(fmt.Sprintf("Create creates a %s", name)).
			Line().Id("Create").Params(ctxC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("CreateTx creates a %s inside transaction", name)).
			Line().Id("CreateTx").Params(ctxC, txC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("CreateMany is a batch-create for %s", name)).
			Line().Id("CreateMany").Params(ctxC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Comment(fmt.Sprintf("CreateManyTx is a batch-create for %s inside transaction", name)).
			Line().Id("CreateManyTx").
			Params(ctxC, txC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Comment(fmt.Sprintf("Query is used for querying many %s", name)).
			Line().Id("Query").Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryTx is used for querying many %s inside transaction", name)).
			Line().Id("QueryTx").Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryOne is used for querying a single %s", name)).
			Line().Id("QueryOne").Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryOneTx is used for querying a single %s inside transaction", name)).
			Line().Id("QueryOneTx").Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("Update updates %s", name)).
			Line().Id("Update").Params(ctxC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("UpdateTx updates %s inside transaction", name)).Line().
			Id("UpdateTx").Params(ctxC, txC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("Delete deletes %s", name)).Line().
			Id("Delete").Params(ctxC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("Delete deletes %s inside transaction", name)).
			Line().Id("DeleteTx").Params(ctxC, txC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment("Aggregate is used for doing aggregation").Line().
			Id("Aggregate").Params(ctxC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
		jen.Comment("Aggregate is used for doing aggregation inside transaction").
			Line().Id("AggregateTx").Params(ctxC, txC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
	).Line().Line()
}
