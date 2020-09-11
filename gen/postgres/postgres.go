package postgres

import (
	"github.com/dave/jennifer/jen"

	"github.com/sf9v/nero/aggregate"
	"github.com/sf9v/nero/comparison"
	gen "github.com/sf9v/nero/gen/internal"
)

const (
	pkgPath = "github.com/sf9v/nero"
	aggPkg  = "github.com/sf9v/nero/aggregate"
	sortPkg = "github.com/sf9v/nero/sort"

	sqPkg      = "github.com/Masterminds/squirrel"
	logPkg     = "github.com/rs/zerolog"
	errPkg     = "github.com/pkg/errors"
	reflectPkg = "github.com/goccy/go-reflect"

	typeName = "PostgreSQLRepository"
	rcvrID   = "pg"
)

var (
	rcvrIDC    = jen.Id(rcvrID)
	rcvrParamC = jen.Add(rcvrIDC).Op("*").Id(typeName)
	ctxC       = jen.Qual("context", "Context")
	ctxIDC     = jen.Id("ctx")
	txC        = jen.Qual(pkgPath, "Tx")
	runnerC    = jen.Qual(pkgPath, "SqlRunner")

	predOps = []comparison.Operator{comparison.Eq, comparison.NotEq,
		comparison.Gt, comparison.GtOrEq, comparison.Lt, comparison.LtOrEq,
		comparison.IsNull, comparison.IsNotNull, comparison.In, comparison.NotIn}
	aggFns = []aggregate.Function{aggregate.Avg, aggregate.Count,
		aggregate.Max, aggregate.Min, aggregate.Sum, aggregate.None}
)

// NewPostgreSQLRepo generates a postgresql repository implementation
func NewPostgreSQLRepo(schema *gen.Schema) *jen.Statement {
	ll := jen.Line().Line()
	// // type definition
	return newTypeDefBlock().Add(ll).
		// assert repository
		Add(newInterfaceGuardBlock()).Add(ll).
		// factory
		Add(newFactoryBlock()).Add(ll).
		// debug
		Add(newDebugBlock()).Add(ll).
		// tx
		Add(newTxBlock()).Add(ll).
		// create
		Add(newCreateBlock(schema)).Add(ll).
		// create tx
		Add(newCreateTxBlock(schema)).Add(ll).
		// create runner
		Add(newCreateRunnerBlock(schema)).Add(ll).
		// create many
		Add(newCreateManyBlock()).Add(ll).
		// create many tx
		Add(newCreateManyTxBlock()).Add(ll).
		// create many runner
		Add(newCreateManyRunnerBlock(schema)).Add(ll).
		// query
		Add(newQueryBlock(schema)).Add(ll).
		// query tx
		Add(newQueryTxBlock(schema)).Add(ll).
		// query runner block
		Add(newQueryRunnerBlock(schema)).Add(ll).
		// query one
		Add(newQueryOneBlock(schema)).Add(ll).
		// query one tx
		Add(newQueryOneTxBlock(schema)).Add(ll).
		// query one runner
		Add(newQueryOneRunnerBlock(schema)).Add(ll).
		// select builder
		Add(newBuildSelectBlock(schema)).Add(ll).
		// update
		Add(newUpdateBlock()).Add(ll).
		// update tx
		Add(newUpdateTxBlock()).Add(ll).
		// update runner block
		Add(newUpdateRunnerBlock(schema)).Add(ll).
		// delete
		Add(newDeleteBlock()).Add(ll).
		// delete tx
		Add(newDeleteTxBlock()).Add(ll).
		// delete runner
		Add(newDeleteRunnerBlock(schema)).Add(ll).
		// aggregate
		Add(newAggregateBlock()).Add(ll).
		// aggregate tx
		Add(newAggregateTxBlock()).Add(ll).
		// aggregate runner
		Add(newAggregateRunnerBlock(schema))
}

func newTypeDefBlock() *jen.Statement {
	return jen.Type().Id(typeName).Struct(
		jen.Id("db").Op("*").Qual("database/sql", "DB"),
		jen.Id("log").Op("*").Qual(logPkg, "Logger"),
	)
}

func newInterfaceGuardBlock() *jen.Statement {
	return jen.Var().Id("_").Id("Repository").Op("=").
		Parens(jen.Op("*").Id(typeName)).Parens(jen.Nil())
}

func newDebugLogBlock(operation string) *jen.Statement {
	return jen.If(
		jen.Id("log").Op(":=").Id(rcvrID).Dot("log"),
		jen.Id("log").Op("!=").Nil(),
	).Block(
		jen.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
			Op(":=").Id("qb").Dot("ToSql").Call(),
		jen.Id("log").Dot("Debug").Call().
			Dot("Str").Call(jen.Lit("method"), jen.Lit(operation)).
			Dot("Str").Call(jen.Lit("stmnt"), jen.Id("sql")).Op(".").Line().
			Id("Interface").Call(jen.Lit("args"), jen.Id("args")).
			Dot("Err").Call(jen.Id("err")).
			Dot("Msg").Call(jen.Lit("")),
	)
}
