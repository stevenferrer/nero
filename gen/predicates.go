package gen

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/jenx"

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
		if !hasPreds(col.Type.T()) {
			continue
		}

		for _, op := range ops {
			opStr := op.String()
			field := col.CamelName()
			if len(col.StructField) > 0 {
				field = col.StructField
			}
			fnName := camel(field + "_" + opStr)
			paramID := lowCamel(field)
			stmnt = stmnt.Func().Id(fnName).
				Params(jen.Id(paramID).
					Add(jenx.Type(col.Type.V()))).
				Params(jen.Id("PredFunc")).
				Block(jen.Return(jen.Func().Params(jen.Id("pb").Op("*").
					Qual(predPkg, "Predicates")).
					Block(jen.Id("pb").Dot("Add").Call(
						jen.Op("&").Qual(predPkg, "Predicate").
							Block(
								jen.Id("Col").Op(":").
									Lit(col.Name).Op(","),
								jen.Id("Op").Op(":").
									Qual(predPkg, opStr).Op(","),
								jen.Id("Val").Op(":").
									Id(paramID).Op(","),
							),
					)),
				)).
				Line().Line()
		}
	}

	return stmnt
}

func hasPreds(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Map, reflect.Slice:
		return false
	}

	return true
}
