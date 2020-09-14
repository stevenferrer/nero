package gen

import (
	"bytes"
	"text/template"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/sort"
)

func newSortsFile(schema *gen.Schema) (*bytes.Buffer, error) {
	v := struct {
		Directions []sort.Direction
		Schema     *gen.Schema
	}{
		Directions: []sort.Direction{
			sort.Asc, sort.Desc,
		},
		Schema: schema,
	}

	tmpl, err := template.New("sorts.tmpl").Parse(sortTmpl)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

const sortTmpl = `
package {{.Schema.Pkg}}

import (
	"github.com/sf9v/nero/sort"
)

type SortFunc func(*sort.Sorts)

{{range $direction := .Directions}}
func {{$direction.String}}(col Column) SortFunc {
	return func(s *sort.Sorts) {
		s.Add(&sort.Sort{
			Col: col.String(),
			Direction: sort.{{$direction.String}},
		})
	}
}
{{end}}
`
