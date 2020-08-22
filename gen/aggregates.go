package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/aggregate"
)

func newAggregates() *jen.Statement {
	aggPkg := pkgPath + "/aggregate"
	aggDoc := "AggFunc is the aggregate function type"
	stmnt := jen.Comment(aggDoc).Line().
		Type().Id("AggFunc").Func().Params(
		jen.Op("*").Qual(aggPkg, "Aggregates"),
	).Line()

	// supported aggregate functions
	aggFns := []aggregate.Function{
		aggregate.Avg, aggregate.Count,
		aggregate.Max, aggregate.Min,
		aggregate.Sum, aggregate.None,
	}

	for _, aggFn := range aggFns {
		stmnt = stmnt.Comment(aggFn.Description()).Line().
			Func().Id(aggFn.String()).
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
