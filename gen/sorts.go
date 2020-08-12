package gen

import (
	"github.com/dave/jennifer/jen"

	"github.com/sf9v/nero/sort"
)

func newSorts() *jen.Statement {
	sortPkg := pkgPath + "/sort"
	stmnt := jen.Comment("SortFunc is the sort function type").Line().
		Type().Id("SortFunc").Func().Params(
		jen.Op("*").Qual(sortPkg, "Sorts"),
	).Line()

	directns := []sort.Direction{sort.Asc, sort.Desc}
	for _, directn := range directns {
		stmnt = stmnt.Comment(directn.Description()).Line().
			Func().Id(directn.String()).
			Params(jen.Id("col").Id("Column")).
			Params(jen.Id("SortFunc")).
			Block(jen.Return(jen.Func().
				Params(jen.Id("srts").Op("*").Qual(sortPkg, "Sorts")).
				Params().Block(jen.Id("srts").Dot("Add").
				Call(jen.Op("&").Qual(sortPkg, "Sort").Block(
					jen.Id("Col").Op(":").Id("col").Dot("String").Call().Op(","),
					jen.Id("Direction").Op(":").Qual(sortPkg, directn.String()).Op(","),
				))))).Line().Line()
	}

	return stmnt
}
