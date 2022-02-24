package nero

// PostgresTemplate is the template for generating a postgres repository
type PostgresTemplate struct {
	filename string
}

var _ Template = (*PostgresTemplate)(nil)

// NewPostgresTemplate returns a new PostgresTemplate
func NewPostgresTemplate() *PostgresTemplate {
	return &PostgresTemplate{
		filename: "postgres.go",
	}
}

// WithFilename overrides the default filename
func (t *PostgresTemplate) WithFilename(filename string) *PostgresTemplate {
	t.filename = filename
	return t
}

// Filename returns the filename
func (t *PostgresTemplate) Filename() string {
	return t.filename
}

// Content returns the template content
func (t *PostgresTemplate) Content() string {
	return postgresTmpl
}

const postgresTmpl = `
{{- fileHeaders -}}

package {{.PkgName}}

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"io"
	"strings"
	"log"
	"os"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero/aggregate"
	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/sort"
	{{range $import := .Imports -}}
		"{{$import}}"
	{{end -}}
)

{{ $fields := prependToFields .Identity .Fields }}

// PostgresRepository is a repository that uses PostgreSQL as data store
type PostgresRepository struct {
	db  *sql.DB
	logger nero.Logger
	debug bool
}

var _ Repository = (*PostgresRepository)(nil)

// NewPostgresRepository returns a PostgresRepository
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Debug enables debug mode
func (repo *PostgresRepository) Debug() *PostgresRepository {
	l := log.New(os.Stdout, "[nero] ", log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix)
	return &PostgresRepository{
		db:  repo.db,	
		debug: true,
		logger: l,
	}
}

// WithLogger overrides the default logger
func (repo *PostgresRepository) WithLogger(logger nero.Logger) *PostgresRepository {	
	repo.logger = logger
	return repo
}

// BeginTx starts a transaction
func (repo *PostgresRepository) BeginTx(ctx context.Context) (nero.Tx, error) {
	return repo.db.BeginTx(ctx, nil)
}

// Create creates a {{.TypeName}}
func (repo *PostgresRepository) Create(ctx context.Context, c *Creator) ({{rawType .Identity.TypeInfo.V}}, error) {
	return repo.create(ctx, repo.db, c)
}

// CreateInTx creates a {{.TypeName}} in a transaction
func (repo *PostgresRepository) CreateInTx(ctx context.Context, tx nero.Tx, c *Creator) ({{rawType .Identity.TypeInfo.V}}, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return {{zeroValue .Identity.TypeInfo.V}}, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.create(ctx, txx, c)
}

func (repo *PostgresRepository) create(ctx context.Context, runner nero.SQLRunner, c *Creator) ({{rawType .Identity.TypeInfo.V}}, error) {
	if err := c.Validate(); err != nil {
		return {{zeroValue .Identity.TypeInfo.V}}, err
	}

	columns := []string{
		{{range $field := $fields -}}
			{{if and (ne $field.IsOptional true) (ne $field.IsAuto true) -}}
				"\"{{$field.Name}}\"",
			{{end -}}
		{{end -}}
	}

	values := []interface{}{
		{{range $field := $fields -}}
			{{if and (ne $field.IsOptional true) (ne $field.IsAuto true) -}}
				{{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
					pq.Array(c.{{$field.Identifier}}),
				{{else -}}
					c.{{$field.Identifier}},
				{{end -}}
			{{end -}}
		{{end -}}
	}

	{{range $field := $fields -}}
		{{if and ($field.IsOptional) (ne $field.IsAuto true) -}}
			if !isZero(c.{{$field.Identifier}}) {
				columns = append(columns, "{{$field.Name}}")
				values = append(values, c.{{$field.Identifier}})
			}
		{{end -}}
	{{end}}

	qb := squirrel.Insert("\"{{.Table}}\"").
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING \"{{.Identity.Name}}\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(runner)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Create, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	var {{.Identity.Identifier}} {{rawType .Identity.TypeInfo.V}}
	err := qb.QueryRowContext(ctx).Scan(&{{.Identity.Identifier}})
	if err != nil {
		return {{zeroValue .Identity.TypeInfo.V}}, err
	}

	return {{.Identity.Identifier}}, nil
}

// CreateMany batch creates {{.TypeNamePlural}}
func (repo *PostgresRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	return repo.createMany(ctx, repo.db, cs...)
}

// CreateManyInTx batch creates {{.TypeNamePlural}} in a transaction
func (repo *PostgresRepository) CreateManyInTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return repo.createMany(ctx, txx, cs...)
}

func (repo *PostgresRepository) createMany(ctx context.Context, runner nero.SQLRunner, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	columns := []string{
		{{range $field := $fields -}}
			{{if ne $field.IsAuto true -}}
				"\"{{$field.Name}}\"",
			{{end -}}
		{{end -}}
	}

	qb := squirrel.Insert("\"{{.Table}}\"").Columns(columns...)
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}

		qb = qb.Values(
			{{range $field := $fields -}}
				{{if ne $field.IsAuto true -}}
					{{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
						pq.Array(c.{{$field.Identifier}}),
					{{else -}}
						c.{{$field.Identifier}},
					{{end -}}
				{{end -}}
			{{end -}}
		)
	}

	qb = qb.Suffix("RETURNING \"{{.Identity.Name}}\"").
		PlaceholderFormat(squirrel.Dollar)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: CreateMany, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	_, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Query queries {{.TypeNamePlural}}
func (repo *PostgresRepository) Query(ctx context.Context, q *Queryer) ([]{{rawType .TypeInfo.V}}, error) {
	return repo.query(ctx, repo.db, q)
}

// QueryInTx queries {{.TypeNamePlural}} in a transaction
func (repo *PostgresRepository) QueryInTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]{{rawType .TypeInfo.V}}, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.query(ctx, txx, q)
}

func (repo *PostgresRepository) query(ctx context.Context, runner nero.SQLRunner, q *Queryer) ([]{{rawType .TypeInfo.V}}, error) {
	qb := repo.buildSelect(q)	
	if repo.debug  && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Query, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	{{.TypeIdentifierPlural}} := []{{rawType .TypeInfo.V}}{}
	for rows.Next() {
		var {{.TypeIdentifier}} {{type .TypeInfo.V}}
		err = rows.Scan(
			{{range $field := $fields -}}
				{{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
					pq.Array(&{{$.TypeIdentifier}}.{{$field.StructField}}),
				{{else -}}
					&{{$.TypeIdentifier}}.{{$field.StructField}},
				{{end -}}
			{{end -}}
		)
		if err != nil {
			return nil, err
		}

		{{.TypeIdentifierPlural}} = append({{.TypeIdentifierPlural}}, &{{.TypeIdentifier}})
	}

	return {{.TypeIdentifierPlural}}, nil
}

// QueryOne queries a {{.TypeName}}
func (repo *PostgresRepository) QueryOne(ctx context.Context, q *Queryer) ({{rawType .TypeInfo.V}}, error) {
	return repo.queryOne(ctx, repo.db, q)
}

// QueryOneInTx queries a {{.TypeName}} in a transaction
func (repo *PostgresRepository) QueryOneInTx(ctx context.Context, tx nero.Tx, q *Queryer) ({{rawType .TypeInfo.V}}, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.queryOne(ctx, txx, q)
}

func (repo *PostgresRepository) queryOne(ctx context.Context, runner nero.SQLRunner, q *Queryer) ({{rawType .TypeInfo.V}}, error) {
	qb := repo.buildSelect(q)
	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: QueryOne, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	var {{.TypeIdentifier}} {{type .TypeInfo.V}}
	err := qb.RunWith(runner).
		QueryRowContext(ctx).
		Scan(
			{{range $field := $fields -}}
				{{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
					pq.Array(&{{$.TypeIdentifier}}.{{$field.StructField}}),
				{{else -}}
					&{{$.TypeIdentifier}}.{{$field.StructField}},
				{{end -}}	
			{{end -}}
		)
	if err != nil {
		return {{zeroValue .TypeInfo.V}}, err
	}

	return &{{.TypeIdentifier}}, nil
}

func (repo *PostgresRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	columns := []string{
		{{range $field := $fields -}}
			"\"{{$field.Name}}\"",
		{{end -}}
	}
	qb := squirrel.Select(columns...).
		From("\"{{.Table}}\"").
		PlaceholderFormat(squirrel.Dollar)

	preds := []comparison.Predicate{}
	for _, predFunc := range q.predFuncs {
		preds = predFunc(preds)
	}
	qb = squirrel.SelectBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), preds))

	sorts := []sort.Sort{}
	for _, sortFunc := range q.sortFuncs {
		sorts = sortFunc(sorts)
	}
	qb = repo.buildSort(qb, sorts)

	if q.limit > 0 {
		qb = qb.Limit(uint64(q.limit))
	}

	if q.offset > 0 {
		qb = qb.Offset(uint64(q.offset))
	}

	return qb
}

func (repo *PostgresRepository) buildPreds(sb squirrel.StatementBuilderType, 
	preds []comparison.Predicate) squirrel.StatementBuilderType {
	for _, pred := range preds {	
		ph := "?"
		fieldX, arg := pred.Field, pred.Arg

		args := []interface{}{}
		if fieldY, ok := arg.(Field); ok { // a field
			ph = fmt.Sprintf("%q", fieldY)
		} else if vals, ok := arg.([]interface{}); ok  {  // array of values 
			args = append(args, vals...)
		} else { // single value
			args = append(args, arg)
		}
		
		switch pred.Op {
		case comparison.Eq:
			sb = sb.Where(fmt.Sprintf("%q = "+ph, fieldX), args...)
		case comparison.NotEq:
			sb = sb.Where(fmt.Sprintf("%q <> "+ph, fieldX), args...)
		case comparison.Gt:
			sb = sb.Where(fmt.Sprintf("%q > "+ph, fieldX), args...)
		case comparison.GtOrEq:
			sb = sb.Where(fmt.Sprintf("%q >= "+ph, fieldX), args...)
		case comparison.Lt:
			sb = sb.Where(fmt.Sprintf("%q < "+ph, fieldX), args...)
		case comparison.LtOrEq:
			sb = sb.Where(fmt.Sprintf("%q <= "+ph, fieldX), args...)
		case comparison.IsNull, comparison.IsNotNull:
			fmtStr := "%q IS NULL"
			if pred.Op == comparison.IsNotNull {
				fmtStr = "%q IS NOT NULL"
			}
			sb = sb.Where(fmt.Sprintf(fmtStr, fieldX))
		case comparison.In, comparison.NotIn:
			fmtStr := "%q IN (%s)"
			if pred.Op == comparison.NotIn {
				fmtStr = "%q NOT IN (%s)"
			}

			phs := []string{}
			for range args {
				phs = append(phs, "?")
			}

			sb = sb.Where(fmt.Sprintf(fmtStr, fieldX, strings.Join(phs, ",")), args...)
		}
	}

	return sb
}

func (repo *PostgresRepository) buildSort(qb squirrel.SelectBuilder, sorts []sort.Sort) squirrel.SelectBuilder {
	for _, s := range sorts {
		field := fmt.Sprintf("%q", s.Field)
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(field + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(field + " DESC")
		}
	}

	return qb
}

// Update updates a {{.TypeName}} or many {{.TypeNamePlural}}
func (repo *PostgresRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	return repo.update(ctx, repo.db, u)
}

// UpdateInTx updates a {{.TypeName}} many {{.TypeNamePlural}} in a transaction
func (repo *PostgresRepository) UpdateInTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.update(ctx, txx, u)
}

func (repo *PostgresRepository) update(ctx context.Context, runner nero.SQLRunner, u *Updater) (int64, error) {
	qb := squirrel.Update("\"{{.Table}}\"").
		PlaceholderFormat(squirrel.Dollar)

	cnt := 0
	{{range $field := .Fields }}
		{{if ne $field.IsAuto true}}
			if !isZero(u.{{$field.Identifier}}) {
				{{if and ($field.IsArray) (ne $field.IsValueScanner true) -}}
					qb = qb.Set("\"{{$field.Name}}\"", pq.Array(u.{{$field.Identifier}}))
				{{else -}}
					qb = qb.Set("\"{{$field.Name}}\"", u.{{$field.Identifier}})
				{{end -}}
				cnt++
			}
		{{end}}
	{{end}}

	if cnt == 0 {
		return 0, nil
	}

	preds := []comparison.Predicate{}
	for _, predFunc := range u.predFuncs {
		preds = predFunc(preds)
	}
	qb = squirrel.UpdateBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), preds))

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Update, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Delete deletes a {{.TypeName}} or many {{.TypeNamePlural}}
func (repo *PostgresRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	return repo.delete(ctx, repo.db, d)
}

// DeleteInTx deletes a {{.TypeName}} or many {{.TypeNamePlural}} in a transaction
func (repo *PostgresRepository) DeleteInTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	return repo.delete(ctx, txx, d)
}

func (repo *PostgresRepository) delete(ctx context.Context, runner nero.SQLRunner, d *Deleter) (int64, error) {
	qb := squirrel.Delete("\"{{.Table}}\"").
		PlaceholderFormat(squirrel.Dollar)

	preds := []comparison.Predicate{}
	for _, predFunc := range d.predFuncs {
		preds = predFunc(preds)
	}
	qb = squirrel.DeleteBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), preds))

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Delete, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	res, err := qb.RunWith(runner).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Aggregate performs an aggregate query
func (repo *PostgresRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	return repo.aggregate(ctx, repo.db, a)
}

// AggregateInTx performs an aggregate query in a transaction
func (repo *PostgresRepository) AggregateInTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	return repo.aggregate(ctx, txx, a)
}

func (repo *PostgresRepository) aggregate(ctx context.Context, runner nero.SQLRunner, a *Aggregator) error {
	aggs := []aggregate.Aggregate{}
	for _, aggFunc := range a.aggFuncs {
		aggs = aggFunc(aggs)
	}
	columns := []string{}
	for _, agg := range aggs {
		field := agg.Field
		qf := fmt.Sprintf("%q", field)
		switch agg.Op {
		case aggregate.Avg:
			columns = append(columns, "AVG("+qf+") avg_"+field)
		case aggregate.Count:
			columns = append(columns, "COUNT("+qf+") count_"+field)
		case aggregate.Max:
			columns = append(columns, "MAX("+qf+") max_"+field)
		case aggregate.Min:
			columns = append(columns, "MIN("+qf+") min_"+field)
		case aggregate.Sum:
			columns = append(columns, "SUM("+qf+") sum_"+field)
		case aggregate.None:
			columns = append(columns, qf)
		}
	}

	qb := squirrel.Select(columns...).From("\"{{.Table}}\"").
		PlaceholderFormat(squirrel.Dollar)

	groupBys := []string{}
	for _, groupBy := range a.groupBys {
		groupBys = append(groupBys, fmt.Sprintf("%q", groupBy.String()))
	}
	qb = qb.GroupBy(groupBys...)

	preds := []comparison.Predicate{}
	for _, predFunc := range a.predFuncs {
		preds = predFunc(preds)
	}
	qb = squirrel.SelectBuilder(repo.buildPreds(squirrel.StatementBuilderType(qb), preds))

	sorts := []sort.Sort{}
	for _, sortFunc := range a.sortFuncs {
		sorts = sortFunc(sorts)
	}	
	qb = repo.buildSort(qb, sorts)

	if repo.debug && repo.logger != nil {
		sql, args, err := qb.ToSql()
		repo.logger.Printf("method: Aggregate, stmt: %q, args: %v, error: %v", sql, args, err)
	}

	rows, err := qb.RunWith(runner).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	v := reflect.ValueOf(a.v).Elem()
	t := reflect.TypeOf(v.Interface()).Elem()
	if len(columns) != t.NumField() {
		return errors.Errorf("column count (%v) and destination struct field count (%v) doesn't match",  len(columns), t.NumField(),)
	}

	for rows.Next() {
		ve := reflect.New(t).Elem()
		dest := make([]interface{}, ve.NumField())
		for i := 0; i < ve.NumField(); i++ {
			dest[i] = ve.Field(i).Addr().Interface()
		}

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		v.Set(reflect.Append(v, ve))
	}

	return nil
}
`
