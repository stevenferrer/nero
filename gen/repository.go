package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"

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
	plural := inflection.Plural(name)
	return jen.Comment(comment).Line().
		Type().Id("Repository").Interface(
		jen.Comment("Tx returns a transaction").
			Line().Id("Tx").Params(ctxC).Params(txC, jen.Error()),
		jen.Comment(fmt.Sprintf("Create creates a %s", name)).
			Line().Id("Create").Params(ctxC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("CreateTx creates a %s in a transaction", name)).
			Line().Id("CreateTx").Params(ctxC, txC, jen.Op("*").Id("Creator")).
			Params(identParamC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("CreateMany creates %s", plural)).
			Line().Id("CreateMany").Params(ctxC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Comment(fmt.Sprintf("CreateManyTx creates %s in a transaction", plural)).
			Line().Id("CreateManyTx").
			Params(ctxC, txC, jen.Op("...").Op("*").Id("Creator")).
			Params(jen.Error()),
		jen.Comment(fmt.Sprintf("Query queries %s", plural)).
			Line().Id("Query").Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryTx queries %s in a transaction", plural)).
			Line().Id("QueryTx").Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Op("[]").Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryOne queries a %s", name)).
			Line().Id("QueryOne").Params(ctxC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("QueryOneTx queries a %s in a transaction", name)).
			Line().Id("QueryOneTx").Params(ctxC, txC, jen.Op("*").Id("Queryer")).
			Params(jen.Add(schemaTypeC), jen.Error()),
		jen.Comment(fmt.Sprintf("Update updates %s", plural)).
			Line().Id("Update").Params(ctxC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("UpdateTx updates %s in a transaction", plural)).Line().
			Id("UpdateTx").Params(ctxC, txC, jen.Op("*").Id("Updater")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("Delete deletes %s", plural)).Line().
			Id("Delete").Params(ctxC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("Delete deletes %s in a transaction", name)).
			Line().Id("DeleteTx").Params(ctxC, txC, jen.Op("*").Id("Deleter")).
			Params(rowsAffectedC, jen.Err().Error()),
		jen.Comment(fmt.Sprintf("Aggregate aggregates %s", plural)).Line().
			Id("Aggregate").Params(ctxC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
		jen.Comment(fmt.Sprintf("Aggregate aggregates %s in a transaction", plural)).
			Line().Id("AggregateTx").Params(ctxC, txC, jen.Op("*").Id("Aggregator")).
			Params(jen.Error()),
	).Line().Line()
}
