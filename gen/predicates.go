package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
)

var ops = []string{
	"Eq", "NotEq", "Gt", "GtOrEq", "Lt", "LtOrEq",
}

func newPredicates(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("PredicateFunc").Func().Params(
		jen.Op("*").Qual(pkgPath+"/predicate", "Builder"),
	).Line()

	predPkg := pkgPath + "/predicate"
	for _, col := range schema.Cols {
		for _, op := range ops {
			field := col.CamelName()
			if len(col.FieldName) > 0 {
				field = col.FieldName
			}
			fn := toCamel(field + "_" + op)
			stmnt = stmnt.Func().
				Id(fn).
				Params(jen.Id(col.LowerCamelName()).
					Add(gen.GetTypeC(col.Typ))).
				Params(jen.Id("PredicateFunc")).
				Block(jen.Return(
					jen.Func().
						Params(jen.Id("pb").Op("*").
							Qual(predPkg, "Builder")).
						Block(
							jen.Id("pb").Dot("Append").Call(
								jen.Op("&").Qual(predPkg, "Predicate").
									Block(
										jen.Id("Field").Op(":").
											Lit(col.Name).Op(","),
										jen.Id("Op").Op(":").
											Qual(predPkg, op).Op(","),
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
