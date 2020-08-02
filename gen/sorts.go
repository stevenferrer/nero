package gen

import (
	"github.com/dave/jennifer/jen"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/sort"
)

func newSorts(schema *gen.Schema) *jen.Statement {
	stmnt := jen.Type().Id("SortFunc").Func().Params(
		jen.Op("*").Qual(pkgPath+"/sort", "Sorts"),
	).Line()

	sortPkg := pkgPath + "/sort"
	dirs := []sort.Direction{sort.Asc, sort.Desc}
	for _, col := range schema.Cols {
		for _, dir := range dirs {
			dirStr := string(dir)
			field := col.CamelName()
			if len(col.Field) > 0 {
				field = col.Field
			}
			fn := camel(field + "_" + dirStr)
			stmnt = stmnt.Func().
				Id(fn).Params().
				Params(jen.Id("SortFunc")).
				Block(jen.Return(jen.Func().Params(jen.Id("srt").Op("*").
					Qual(sortPkg, "Sorts")).Block(jen.Id("srt").Dot("Add").
					Call(
						jen.Op("&").Qual(sortPkg, "Sort").Block(
							jen.Id("Field").Op(":").
								Lit(col.Name).Op(","),
							jen.Id("Direction").Op(":").
								Qual(sortPkg, dirStr).Op(","),
						),
					),
				),
				)).
				Line().Line()
		}
	}

	return stmnt
}
