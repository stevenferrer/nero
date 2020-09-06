package gen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/goccy/go-reflect"

	"github.com/sf9v/nero/comparison"
	gen "github.com/sf9v/nero/gen/internal"
	jenx "github.com/sf9v/nero/x/jen"
)

func newPredicates(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Comment("PredFunc is the predicate function type").
		Line().Type().Id("PredFunc").Func().Params(
		jen.Op("*").Qual(pkgPath+"/comparison", "Predicates"),
	).Line()

	ops := []comparison.Operator{comparison.Eq, comparison.NotEq,
		comparison.Gt, comparison.GtOrEq, comparison.Lt, comparison.LtOrEq,
		comparison.IsNull, comparison.IsNotNull}
	compPkg := pkgPath + "/comparison"
	for _, col := range schema.Cols {
		kind := col.Type.T().Kind()
		if kind == reflect.Map ||
			kind == reflect.Slice {
			continue
		}

		for _, op := range ops {
			field := col.CamelName()
			if len(col.StructField) > 0 {
				field = col.StructField
			}
			opStr := op.String()
			fnName := camel(field + "_" + opStr)

			if op == comparison.IsNull ||
				op == comparison.IsNotNull {
				if !col.Nullable {
					continue
				}

				stmnt = stmnt.Func().Id(fnName).
					Params().
					Params(jen.Id("PredFunc")).
					Block(jen.Return(jen.Func().Params(jen.Id("pb").Op("*").
						Qual(compPkg, "Predicates")).
						Block(jen.Id("pb").Dot("Add").Call(
							jen.Op("&").Qual(compPkg, "Predicate").
								Block(
									jen.Id("Col").Op(":").
										Lit(col.Name).Op(","),
									jen.Id("Op").Op(":").
										Qual(compPkg, opStr).Op(","),
								),
						)),
					)).Line().Line()

				continue
			}

			paramID := lowCamel(field)
			fnDoc := fmt.Sprintf("%s returns a/an %s predicate on %s", fnName, strings.ToLower(op.Description()), paramID)
			stmnt = stmnt.Comment(fnDoc).Line().
				Func().Id(fnName).
				Params(jen.Id(paramID).
					Add(jenx.Type(col.Type.V()))).
				Params(jen.Id("PredFunc")).
				Block(jen.Return(jen.Func().Params(jen.Id("pb").Op("*").
					Qual(compPkg, "Predicates")).
					Block(jen.Id("pb").Dot("Add").Call(
						jen.Op("&").Qual(compPkg, "Predicate").
							Block(
								jen.Id("Col").Op(":").
									Lit(col.Name).Op(","),
								jen.Id("Op").Op(":").
									Qual(compPkg, opStr).Op(","),
								jen.Id("Val").Op(":").
									Id(paramID).Op(","),
							),
					)),
				)).Line().Line()
		}
	}

	return stmnt
}
