package gen

import (
	"bytes"
	"text/template"

	"github.com/stevenferrer/nero"
)

func newRepositoryFile(schema nero.Schema) (*File, error) {
	tmpl, err := template.New("repository.tmpl").
		Funcs(nero.NewFuncMap()).Parse(repositoryTmpl)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, schema)
	if err != nil {
		return nil, err
	}

	return &File{name: "repository.go", buf: buf.Bytes()}, nil
}

const repositoryTmpl = `
{{- fileHeaders -}}

package {{.PkgName}}

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/sort"
	"github.com/stevenferrer/nero/aggregate"
	multierror "github.com/hashicorp/go-multierror"
	{{range $import := .Imports -}}
		"{{$import}}"
	{{end -}}
)

// Repository is an interface that provides the methods 
// for interacting with a {{.TypeInfo.Name}} repository
type Repository interface {
	// BeginTx starts a transaction
	BeginTx(context.Context) (nero.Tx, error)
	// Create creates a {{.TypeName}}
	Create(context.Context, *Creator) (id {{rawType .Identity.TypeInfo.V}}, err error)
	// CreateInTx creates a {{.TypeName}} in a transaction
	CreateInTx(context.Context, nero.Tx, *Creator) (id {{rawType .Identity.TypeInfo.V}}, err error)
	// CreateMany batch creates {{.TypeNamePlural}}
	CreateMany(context.Context, ...*Creator) error
	// CreateManyInTx batch creates {{.TypeNamePlural}} in a transaction
	CreateManyInTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries {{.TypeNamePlural}}
	Query(context.Context, *Queryer) ([]{{rawType .TypeInfo.V}}, error)
	// QueryTx queries {{.TypeNamePlural}} in a transaction
	QueryInTx(context.Context, nero.Tx, *Queryer) ([]{{rawType .TypeInfo.V}}, error)
	// QueryOne queries a {{.TypeName}}
	QueryOne(context.Context, *Queryer) ({{rawType .TypeInfo.V}}, error)
	// QueryOneTx queries a {{.TypeName}} in a transaction
	QueryOneInTx(context.Context, nero.Tx, *Queryer) ({{rawType .TypeInfo.V}}, error)
	// Update updates a {{.TypeName}} or many {{.TypeNamePlural}}
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates a {{.TypeName}} many {{.TypeNamePlural}} in a transaction
	UpdateInTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes a {{.TypeName}} or many {{.TypeNamePlural}}
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes a {{.TypeName}} or many {{.TypeNamePlural}} in a transaction
	DeleteInTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate performs an aggregate query
	Aggregate(context.Context, *Aggregator) error
	// Aggregate performs an aggregate query in a transaction
	AggregateInTx(context.Context, nero.Tx, *Aggregator) error
}


{{ $fields := prependToFields .Identity .Fields }}

// Creator is a create builder
type Creator struct {
	{{range $field := $fields -}}
		{{if ne $field.IsAuto true -}}
		{{$field.Identifier}} {{rawType $field.TypeInfo.V}}
		{{end -}}
	{{end -}}
}

// NewCreator returns a Creator
func NewCreator() *Creator {
	return &Creator{}
}

{{range $field := $fields }}
	{{if ne $field.IsAuto true -}}
		// {{$field.StructField}} sets the {{$field.StructField}} field
		func (c *Creator) {{$field.StructField}}({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) *Creator {
			c.{{$field.Identifier}} = {{$field.Identifier}}
			return c
		}
	{{end -}}
{{end -}}

// Validate validates the fields
func (c *Creator) Validate() error {
	var err error
	{{range $field := .Fields -}}
		{{if and (ne $field.IsOptional true) (ne $field.IsAuto true) -}}
			if isZero(c.{{$field.Identifier}}) {
				err = multierror.Append(err, nero.NewErrRequiredField("{{$field.Name}}"))
			}
		{{end}} 
	{{end}}

	return err
}

// Queryer is a query builder
type Queryer struct {
	limit  uint
	offset uint
	predFuncs []comparison.PredFunc
	sortFuncs []sort.SortFunc
}

// NewQueryer returns a Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where applies predicates
func (q *Queryer) Where(predFuncs ...comparison.PredFunc) *Queryer {
	q.predFuncs = append(q.predFuncs, predFuncs...)
	return q
}

// Sort applies sorting expressions
func (q *Queryer) Sort(sortFuncs ...sort.SortFunc) *Queryer {
	q.sortFuncs = append(q.sortFuncs, sortFuncs...)
	return q
}

// Limit applies limit
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset applies offset
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder
type Updater struct {
	{{range $field := .Fields -}}
		{{if ne $field.IsAuto true -}}
			{{$field.Identifier}} {{rawType $field.TypeInfo.V}}
		{{end -}}
	{{end -}}
	predFuncs []comparison.PredFunc
}

// NewUpdater returns an Updater
func NewUpdater() *Updater {
	return &Updater{}
}

{{range $field := .Fields}}
	{{if ne $field.IsAuto true -}}
		// {{$field.StructField}} sets the {{$field.StructField}} field
		func (c *Updater) {{$field.StructField}}({{$field.Identifier}} {{rawType $field.TypeInfo.V}}) *Updater {
			c.{{$field.Identifier}} = {{$field.Identifier}}
			return c
		}
	{{end -}}
{{end -}}

// Where applies predicates
func (u *Updater) Where(predFuncs ...comparison.PredFunc) *Updater {
	u.predFuncs = append(u.predFuncs, predFuncs...)
	return u
}

// Deleter is a delete builder
type Deleter struct {
	predFuncs []comparison.PredFunc
}

// NewDeleter returns a Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where applies predicates
func (d *Deleter) Where(predFuncs ...comparison.PredFunc) *Deleter {
	d.predFuncs = append(d.predFuncs, predFuncs...)
	return d
}

// Aggregator is an aggregate query builder
type Aggregator struct {
	v      interface{}
	aggFuncs	[]aggregate.AggFunc
	predFuncs	[]comparison.PredFunc
	sortFuncs	[]sort.SortFunc
	groupBys []Field
}

// NewAggregator expects a v and returns an Aggregator 
// where 'v' argument must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{v: v}
}

// Aggregate applies aggregate functions
func (a *Aggregator) Aggregate(aggFuncs ...aggregate.AggFunc) *Aggregator {
	a.aggFuncs = append(a.aggFuncs, aggFuncs...)
	return a
}

// Where applies predicates
func (a *Aggregator) Where(predFuncs ...comparison.PredFunc) *Aggregator {
	a.predFuncs = append(a.predFuncs, predFuncs...)
	return a
}

// Sort applies sorting expressions
func (a *Aggregator) Sort(sortFuncs ...sort.SortFunc) *Aggregator {
	a.sortFuncs = append(a.sortFuncs, sortFuncs...)
	return a
}

// Group applies group clauses
func (a *Aggregator) GroupBy(fields ...Field) *Aggregator {
	a.groupBys = append(a.groupBys, fields...)
	return a
}

// rollback performs a rollback
func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}

// isZero checks if v is a zero-value
func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
`
