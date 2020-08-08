package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newTypeDefBlock(t *testing.T) {
	block := newTypeDefBlock()
	expect := strings.TrimSpace(`
type PostgreSQLRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newTypeAssertBlock(t *testing.T) {
	block := newTypeAssertBlock()
	expect := `var _ = Repository(&PostgreSQLRepository{})`
	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func Test_newDebugLogBlock(t *testing.T) {
	block := newDebugLogBlock("Query")
	expect := strings.TrimSpace(`
if log := pg.log; log != nil {
	sql, args, err := qb.ToSql()
	log.Debug().Str("method", "Query").Str("stmnt", sql).
		Interface("args", args).Err(err).Msg("")
}
`)
	got := strings.TrimSpace(fmt.Sprintf("%#v", block))
	assert.Equal(t, expect, got)
}

func TestNewPostgreSQLRepo(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	stmnt := NewPostgreSQLRepo(schema)
	expect := strings.TrimSpace(`
type PostgreSQLRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}

var _ = Repository(&PostgreSQLRepository{})

func NewPostgreSQLRepository(db *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (pg *PostgreSQLRepository) Debug(out io.Writer) *PostgreSQLRepository {
	lg := zerolog.New(out).With().Timestamp().Logger()
	return &PostgreSQLRepository{
		db:  pg.db,
		log: &lg,
	}
}

func (pg *PostgreSQLRepository) Tx(ctx context.Context) (nero.Tx, error) {
	return pg.db.BeginTx(ctx, nil)
}

func (pg *PostgreSQLRepository) Create(ctx context.Context, c *Creator) (int64, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return 0, err
	}

	id, err := pg.CreateTx(ctx, tx, c)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return id, tx.Commit()
}

func (pg *PostgreSQLRepository) CreateMany(ctx context.Context, cs ...*Creator) error {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return err
	}

	err = pg.CreateManyTx(ctx, tx, cs...)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func (pg *PostgreSQLRepository) CreateTx(ctx context.Context, tx nero.Tx, c *Creator) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	columns := []string{}
	values := []interface{}{}
	if c.name != "" {
		columns = append(columns, "\"name\"")
		values = append(values, c.name)
	}
	if c.group != "" {
		columns = append(columns, "\"group_res\"")
		values = append(values, c.group)
	}
	if c.updatedAt != nil {
		columns = append(columns, "\"updated_at\"")
		values = append(values, c.updatedAt)
	}

	qb := squirrel.Insert("\"users\"").
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(txx)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Create").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	var id int64
	err := qb.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pg *PostgreSQLRepository) CreateManyTx(ctx context.Context, tx nero.Tx, cs ...*Creator) error {
	if len(cs) == 0 {
		return nil
	}

	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	columns := []string{"\"name\"", "\"group_res\"", "\"updated_at\""}
	qb := squirrel.Insert("\"users\"").Columns(columns...)
	for _, c := range cs {
		qb = qb.Values(c.name, c.group, c.updatedAt)
	}

	qb = qb.Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "CreateMany").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	_, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PostgreSQLRepository) Query(ctx context.Context, q *Queryer) ([]*example.User, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return nil, err
	}

	list, err := pg.QueryTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return list, tx.Commit()
}

func (pg *PostgreSQLRepository) QueryOne(ctx context.Context, q *Queryer) (*example.User, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return nil, err
	}

	item, err := pg.QueryOneTx(ctx, tx, q)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return item, tx.Commit()
}

func (pg *PostgreSQLRepository) QueryTx(ctx context.Context, tx nero.Tx, q *Queryer) ([]*example.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := pg.buildSelect(q)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Query").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(txx).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*example.User{}
	for rows.Next() {
		var item example.User
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Group,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &item)
	}

	return list, nil
}

func (pg *PostgreSQLRepository) QueryOneTx(ctx context.Context, tx nero.Tx, q *Queryer) (*example.User, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, errors.New("expecting tx to be *sql.Tx")
	}

	qb := pg.buildSelect(q)
	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "QueryOne").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	var item example.User
	err := qb.RunWith(txx).
		QueryRowContext(ctx).
		Scan(
			&item.ID,
			&item.Name,
			&item.Group,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (pg *PostgreSQLRepository) buildSelect(q *Queryer) squirrel.SelectBuilder {
	columns := []string{"\"id\"", "\"name\"", "\"group_res\"", "\"updated_at\"", "\"created_at\""}
	qb := squirrel.Select(columns...).
		From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := q.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case comparison.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case comparison.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case comparison.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case comparison.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case comparison.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		}
	}

	sfs := q.sfs
	sorts := &sort.Sorts{}
	for _, sf := range sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := fmt.Sprintf("%q", s.Col)
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if q.limit > 0 {
		qb = qb.Limit(q.limit)
	}

	if q.offset > 0 {
		qb = qb.Offset(q.offset)
	}

	return qb
}

func (pg *PostgreSQLRepository) Update(ctx context.Context, u *Updater) (int64, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := pg.UpdateTx(ctx, tx, u)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (pg *PostgreSQLRepository) UpdateTx(ctx context.Context, tx nero.Tx, u *Updater) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Update("\"users\"").PlaceholderFormat(squirrel.Dollar)
	if u.name != "" {
		qb = qb.Set("\"name\"", u.name)
	}
	if u.group != "" {
		qb = qb.Set("\"group_res\"", u.group)
	}
	if u.updatedAt != nil {
		qb = qb.Set("\"updated_at\"", u.updatedAt)
	}

	pfs := u.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case comparison.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case comparison.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case comparison.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case comparison.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case comparison.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		}
	}

	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Update").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (pg *PostgreSQLRepository) Delete(ctx context.Context, d *Deleter) (int64, error) {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := pg.DeleteTx(ctx, tx, d)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return rowsAffected, tx.Commit()
}

func (pg *PostgreSQLRepository) DeleteTx(ctx context.Context, tx nero.Tx, d *Deleter) (int64, error) {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return 0, errors.New("expecting tx to be *sql.Tx")
	}

	qb := squirrel.Delete("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	pfs := d.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case comparison.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case comparison.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case comparison.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case comparison.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case comparison.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		}
	}

	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Delete").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	res, err := qb.RunWith(txx).ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (pg *PostgreSQLRepository) Aggregate(ctx context.Context, a *Aggregator) error {
	tx, err := pg.Tx(ctx)
	if err != nil {
		return err
	}

	err = pg.AggregateTx(ctx, tx, a)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func (pg *PostgreSQLRepository) AggregateTx(ctx context.Context, tx nero.Tx, a *Aggregator) error {
	txx, ok := tx.(*sql.Tx)
	if !ok {
		return errors.New("expecting tx to be *sql.Tx")
	}

	aggs := &aggregate.Aggregates{}
	for _, aggf := range a.aggfs {
		aggf(aggs)
	}
	cols := []string{}
	for _, agg := range aggs.All() {
		col := agg.Col
		qcol := fmt.Sprintf("%q", col)
		switch agg.Fn {
		case aggregate.Avg:
			cols = append(cols, "AVG("+qcol+") avg_"+col)
		case aggregate.Count:
			cols = append(cols, "COUNT("+qcol+") count_"+col)
		case aggregate.Max:
			cols = append(cols, "MAX("+qcol+") max_"+col)
		case aggregate.Min:
			cols = append(cols, "MIN("+qcol+") min_"+col)
		case aggregate.Sum:
			cols = append(cols, "SUM("+qcol+") sum_"+col)
		case aggregate.None:
			cols = append(cols, qcol)
		}
	}

	qb := squirrel.Select(cols...).From("\"users\"").
		PlaceholderFormat(squirrel.Dollar)

	groups := []string{}
	for _, group := range a.groups {
		groups = append(groups, fmt.Sprintf("%q", group.String()))
	}
	qb = qb.GroupBy(groups...)

	pfs := a.pfs
	pb := &comparison.Predicates{}
	for _, pf := range pfs {
		pf(pb)
	}
	for _, p := range pb.All() {
		switch p.Op {
		case comparison.Eq:
			qb = qb.Where(fmt.Sprintf("%q = ?", p.Col), p.Val)
		case comparison.NotEq:
			qb = qb.Where(fmt.Sprintf("%q <> ?", p.Col), p.Val)
		case comparison.Gt:
			qb = qb.Where(fmt.Sprintf("%q > ?", p.Col), p.Val)
		case comparison.GtOrEq:
			qb = qb.Where(fmt.Sprintf("%q >= ?", p.Col), p.Val)
		case comparison.Lt:
			qb = qb.Where(fmt.Sprintf("%q < ?", p.Col), p.Val)
		case comparison.LtOrEq:
			qb = qb.Where(fmt.Sprintf("%q <= ?", p.Col), p.Val)
		case comparison.IsNull:
			qb = qb.Where(fmt.Sprintf("%q IS NULL", p.Col))
		case comparison.IsNotNull:
			qb = qb.Where(fmt.Sprintf("%q IS NOT NULL", p.Col))
		}
	}

	sfs := a.sfs
	sorts := &sort.Sorts{}
	for _, sf := range sfs {
		sf(sorts)
	}
	for _, s := range sorts.All() {
		col := fmt.Sprintf("%q", s.Col)
		switch s.Direction {
		case sort.Asc:
			qb = qb.OrderBy(col + " ASC")
		case sort.Desc:
			qb = qb.OrderBy(col + " DESC")
		}
	}

	if log := pg.log; log != nil {
		sql, args, err := qb.ToSql()
		log.Debug().Str("method", "Aggregate").Str("stmnt", sql).
			Interface("args", args).Err(err).Msg("")
	}

	rows, err := qb.RunWith(txx).QueryContext(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	v := reflect.ValueOf(a.v).Elem()
	t := reflect.TypeOf(v.Interface()).Elem()
	if t.NumField() != len(cols) {
		return errors.New("aggregate columns and destination struct field count should match")
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
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", stmnt))
	assert.Equal(t, expect, got)
}
