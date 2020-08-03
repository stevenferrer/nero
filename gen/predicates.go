package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/predicate"
)

func newPredicates(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("PredFunc").Func().Params(
		jen.Op("*").Qual(pkgPath+"/predicate", "Predicates"),
	).Line()

	ops := []predicate.Operator{predicate.Eq, predicate.NotEq, predicate.Gt,
		predicate.GtOrEq, predicate.Lt, predicate.LtOrEq}
	predPkg := pkgPath + "/predicate"
	for _, col := range schema.Cols {
		for _, op := range ops {
			opStr := string(op.String())
			structField := col.CamelName()
			if len(col.StructField) > 0 {
				structField = col.StructField
			}
			fn := camel(structField + "_" + opStr)
			stmnt = stmnt.Func().
				Id(fn).
				Params(jen.Id(col.LowerCamelName()).
					Add(gen.GetTypeC(col.Typ))).
				Params(jen.Id("PredFunc")).
				Block(jen.Return(
					jen.Func().
						Params(jen.Id("pb").Op("*").
							Qual(predPkg, "Predicates")).
						Block(
							jen.Id("pb").Dot("Add").Call(
								jen.Op("&").Qual(predPkg, "Predicate").
									Block(
										jen.Id("Col").Op(":").
											Lit(col.Name).Op(","),
										jen.Id("Op").Op(":").
											Qual(predPkg, opStr).Op(","),
										jen.Id("Val").Op(":").
											Id(col.LowerCamelName()).Op(","),
									),
							),
						),
				)).
				Line().Line()
		}
	}

	return stmnt
}
