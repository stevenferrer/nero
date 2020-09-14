package gen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/jinzhu/inflection"

	gen "github.com/sf9v/nero/gen/internal"
)

func newRepositoryFile(schema *gen.Schema) (*bytes.Buffer, error) {
	tmpl, err := template.New("repository.tmpl").Funcs(template.FuncMap{
		"plural": inflection.Plural,
		"type": func(v interface{}) string {
			return fmt.Sprintf("%T", v)
		},
	}).Parse(repositoryTmpl)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, schema)
	return buf, err
}

const repositoryTmpl = `
package {{.Pkg}}

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sf9v/nero"

	{{range $import := .SchemaImports -}}
		"{{$import}}"
	{{end -}}
	{{range $import := .ColumnImports -}}
		"{{$import}}"
	{{end -}}
)

type Repository interface {
	Tx(context.Context) (nero.Tx, error)
	Create(context.Context, *Creator) (id {{type .Ident.Type.V}}, err error)
	CreateTx(context.Context, nero.Tx, *Creator) (id {{type .Ident.Type.V}}, err error)
	CreateMany(context.Context, ...*Creator) error
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	Query(context.Context, *Queryer) ([]{{type .Type.V}}, error)
	QueryTx(context.Context, nero.Tx, *Queryer) ([]{{type .Type.V}}, error)
	QueryOne(context.Context, *Queryer) ({{type .Type.V}}, error)
	QueryOneTx(context.Context, nero.Tx, *Queryer) ({{type .Type.V}}, error)
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	Aggregate(context.Context, *Aggregator) error
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}

type Creator struct {
	{{range $col := .Cols -}}
		{{if ne $col.Auto true -}}
		{{$col.Identifier}} {{type $col.Type.V}}
		{{end -}}
	{{end -}}
}

func NewCreator() *Creator {
	return &Creator{}
}

{{range $col := .Cols}}
	{{if ne $col.Auto true -}}
		func (c *Creator) {{$col.Field}}({{$col.Identifier}} {{type $col.Type.V}}) *Creator {
			c.{{$col.Identifier}} = {{$col.Identifier}}
			return c
		}
	{{end -}}
{{end -}}

type Queryer struct {
	limit  uint
	offset uint
	pfs    []PredFunc
	sfs    []SortFunc
}

func NewQueryer() *Queryer {
	return &Queryer{}
}

func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

type Updater struct {
	{{range $col := .Cols -}}
		{{if ne $col.Auto true -}}
		{{$col.Identifier}} {{type $col.Type.V}}
		{{end -}}
	{{end -}}
	pfs []PredFunc
}

func NewUpdater() *Updater {
	return &Updater{}
}

{{range $col := .Cols}}
	{{if ne $col.Auto true -}}
		func (c *Updater) {{$col.Field}}({{$col.Identifier}} {{type $col.Type.V}}) *Updater {
			c.{{$col.Identifier}} = {{$col.Identifier}}
			return c
		}
	{{end -}}
{{end -}}

func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}

type Deleter struct {
	pfs []PredFunc
}

func NewDeleter() *Deleter {
	return &Deleter{}
}

func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}

type Aggregator struct {
	v      interface{}
	aggfs  []AggFunc
	pfs    []PredFunc
	sfs    []SortFunc
	groups []Column
}

func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{
		v: v,
	}
}

func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
	return a
}

func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}
`
