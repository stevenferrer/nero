package postgres

import (
	"github.com/dave/jennifer/jen"
	"github.com/sf9v/nero/sort"
)

func newSortsBlock() *jen.Statement {
	return jen.Id("sorts").Op(":=").Op("&").
		Qual(sortPkg, "Sorts").Block().Line().
		For(jen.List(jen.Id("_"), jen.Id("sf")).
			Op(":=").Range().Id("sfs")).
		Block(jen.Id("sf").Call(jen.Id("sorts"))).Line().
		For(jen.List(jen.Id("_"), jen.Id("s").Op(":=").
			Range().Id("sorts").Dot("All").Call())).
		Block(
			jen.Id("col").Op(":=").Qual("fmt", "Sprintf").
				Call(jen.Lit("%q"), jen.Id("s").Dot("Col")),
			jen.Switch(jen.Id("s").Dot("Direction")).
				BlockFunc(func(g *jen.Group) {
					// ascending
					g.Case(jen.Qual(sortPkg, sort.Asc.String())).
						Block(jen.Id("qb").Op("=").Id("qb").Dot("OrderBy").
							Call(jen.Id("col").Op("+").Lit(" ASC")))
					// descending
					g.Case(jen.Qual(sortPkg, sort.Desc.String())).
						Block(jen.Id("qb").Op("=").Id("qb").Dot("OrderBy").
							Call(jen.Id("col").Op("+").Lit(" DESC")))
				}))
}
