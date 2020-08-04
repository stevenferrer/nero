package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/aggregate"
	gen "github.com/sf9v/nero/gen/internal"
)

func newAggregates(schema *gen.Schema) *jen.Statement {
	aggPkg := pkgPath + "/aggregate"
	stmnt := jen.Type().Id("AggFunc").Func().Params(
		jen.Op("*").Qual(aggPkg, "Aggregates"),
	).Line()

	aggFns := []aggregate.Function{
		aggregate.Avg, aggregate.Count,
		aggregate.Max, aggregate.Min,
		aggregate.Sum,
	}

	for _, aggFn := range aggFns {
		stmnt = stmnt.Func().Id(aggFn.String()).
			Params(jen.Id("col").Id("Column")).
			Params(jen.Id("AggFunc")).
			Block(jen.Return(jen.Func().
				Params(jen.Id("aggs").Op("*").Qual(aggPkg, "Aggregates")).
				Params().Block(jen.Id("aggs").Dot("Add").
				Call(
					jen.Op("&").Qual(aggPkg, "Aggregate").Block(
						jen.Id("Col").Op(":").Id("col").Dot("String").Call().Op(","),
						jen.Id("Fn").Op(":").Qual(aggPkg, aggFn.String()).Op(","),
					),
				)))).Line().Line()
	}

	return stmnt
}
