package postgres

import (
	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/comparison"
)

func newPredicatesBlock() *jen.Statement {
	// predicates
	return jen.Id("pb").Op(":=").Op("&").
		Qual(pkgPath+"/comparison", "Predicates").Block().Line().
		For(jen.List(jen.Id("_"), jen.Id("pf")).
			Op(":=").Range().Id("pfs")).
		Block(jen.Id("pf").Call(jen.Id("pb"))).Line().
		For(jen.List(jen.Id("_"), jen.Id("p").Op(":=").
			Range().Id("pb").Dot("All").Call())).
		Block(
			// switch block
			jen.Switch(jen.Id("p").Dot("Op")).
				BlockFunc(func(g *jen.Group) {
					inBlock := false
					for _, op := range predOps {
						var fmtExpr string
						if op == comparison.In || op == comparison.NotIn {
							if inBlock {
								continue
							}

							g.Case(
								jen.Qual(pkgPath+"/comparison", comparison.In.String()),
								jen.Qual(pkgPath+"/comparison", comparison.NotIn.String()),
							).BlockFunc(func(g *jen.Group) {
								g.Id("args").Op(":=").Id("p").Dot("Val").Assert(jen.Op("[]").Interface())
								g.If(jen.Len(jen.Id("args")).Op("==").Lit(0)).Block(jen.Continue())
								g.Id("qms").Op(":=").Op("[]").String().Block()
								g.For(jen.Range().Id("args")).Block(
									jen.Id("qms").Op("=").Append(jen.Id("qms"), jen.Lit("?")),
								)
								g.Id("fmtStr").Op(":=").Lit("%q IN (%s)")
								g.If(jen.Id("p").Dot("Op").Op("==").Qual(pkgPath+"/comparison", "NotIn")).Block(
									jen.Id("fmtStr").Op("=").Lit("%q NOT IN (%s)"),
								)
								g.Id("plchldr").Op(":=").Qual("strings", "Join").Call(jen.Id("qms"), jen.Lit(","))
								g.Id("qb").Op("=").Id("qb").Dot("Where").Call(
									jen.Qual("fmt", "Sprintf").Call(
										jen.Id("fmtStr"),
										jen.Id("p").Dot("Col"),
										jen.Id("plchldr"),
									),
									jen.Id("args").Op("..."),
								)
							})

							inBlock = true
							continue
						}

						switch op {
						case comparison.Eq:
							fmtExpr = "%q = ?"
						case comparison.NotEq:
							fmtExpr = "%q <> ?"
						case comparison.Gt:
							fmtExpr = "%q > ?"
						case comparison.GtOrEq:
							fmtExpr = "%q >= ?"
						case comparison.Lt:
							fmtExpr = "%q < ?"
						case comparison.LtOrEq:
							fmtExpr = "%q <= ?"
						case comparison.IsNull:
							fmtExpr = "%q IS NULL"
						case comparison.IsNotNull:
							fmtExpr = "%q IS NOT NULL"
						}

						val := jen.Id("p").Dot("Val")
						if op == comparison.IsNull ||
							op == comparison.IsNotNull {
							val = nil
						}

						g.Case(jen.Qual(pkgPath+"/comparison", op.String())).
							Block(jen.Id("qb").Op("=").Id("qb").Dot("Where").
								Call(
									jen.Qual("fmt", "Sprintf").Call(
										jen.Lit(fmtExpr),
										jen.Id("p").Dot("Col"),
									),
									val,
								))
					}
				}))
}
