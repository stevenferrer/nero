package sqlite

import (
	"github.com/dave/jennifer/jen"

	"github.com/sf9v/nero/aggregate"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/sf9v/nero/predicate"
)

const (
	pkgPath = "github.com/sf9v/nero"
	aggPkg  = "github.com/sf9v/nero/aggregate"
	predPkg = "github.com/sf9v/nero/predicate"
	sortPkg = "github.com/sf9v/nero/sort"

	sqPkg  = "github.com/Masterminds/squirrel"
	logPkg = "github.com/rs/zerolog"
	errPkg = "github.com/pkg/errors"

	typeName = "SQLiteRepository"
	rcvrID   = "sl"
)

var (
	rcvrIDC     = jen.Id(rcvrID)
	rcvrParamC  = jen.Add(rcvrIDC).Op("*").Id(typeName)
	ctxC        = jen.Qual("context", "Context")
	ctxIDC      = jen.Id("ctx")
	txC         = jen.Qual(pkgPath, "Tx")
	txCommitC   = jen.Id("tx").Dot("Commit").Call()
	txRollbackC = jen.Id("rollback").Call(jen.Id("tx"), jen.Err())

	predOps = []predicate.Operator{predicate.Eq, predicate.NotEq,
		predicate.Gt, predicate.GtOrEq, predicate.Lt, predicate.LtOrEq}
	aggFns = []aggregate.Function{aggregate.Avg, aggregate.Count,
		aggregate.Max, aggregate.Min, aggregate.Sum}
)

// NewSQLiteRepo generates an sqlite repository implementation
func NewSQLiteRepo(schema *gen.Schema) *jen.Statement {
	ll := jen.Line().Line()
	// // type definition
	return newTypeDefBlock().Add(ll).
		// assert repository
		Add(newTypeAssertBlock()).Add(ll).
		// factory
		Add(newFactoryBlock()).Add(ll).
		// debug
		Add(newDebugBlock()).Add(ll).
		// tx
		Add(newTxBlock()).Add(ll).
		// create
		Add(newCreateBlock(schema)).Add(ll).
		// create many
		Add(newCreateManyBlock()).Add(ll).
		// create tx
		Add(newCreateTxBlock(schema)).Add(ll).
		// create many tx
		Add(newCreateManyTxBlock(schema)).Add(ll).
		// query
		Add(newQueryBlock(schema)).Add(ll).
		// query one
		Add(newQueryOneBlock(schema)).Add(ll).
		// query tx
		Add(newQueryTxBlock(schema)).Add(ll).
		// query one tx
		Add(newQueryOneTxBlock(schema)).Add(ll).
		// select builder
		Add(newSelectBuilderBlock()).Add(ll).
		// update
		Add(newUpdateBlock()).Add(ll).
		// update tx
		Add(newUpdateTxBlock(schema)).Add(ll).
		// delete
		Add(newDeleteBlock()).Add(ll).
		// delete tx
		Add(newDeleteTxBlock()).Add(ll).
		// aggregate
		Add(newAggregateBlock()).Add(ll).
		// aggregate tx
		Add(newAggregateTxBlock()).Add(ll)
}

func newTypeDefBlock() *jen.Statement {
	return jen.Type().Id(typeName).Struct(
		jen.Id("db").Op("*").Qual("database/sql", "DB"),
		jen.Id("log").Op("*").Qual(logPkg, "Logger"),
	)
}

func newTypeAssertBlock() *jen.Statement {
	return jen.Var().Id("_").Op("=").Id("Repository").
		Call(jen.Op("&").Id(typeName).Block())
}

func newDebugLogBlock(operation string) *jen.Statement {
	return jen.If(
		jen.Id("log").Op(":=").Id(rcvrID).Dot("log"),
		jen.Id("log").Op("!=").Nil(),
	).Block(
		jen.List(jen.Id("sql"), jen.Id("args"), jen.Id("err")).
			Op(":=").Id("qb").Dot("ToSql").Call(),
		jen.Id("log").Dot("Debug").Call().
			Dot("Str").Call(jen.Lit("op"), jen.Lit(operation)).
			Dot("Str").Call(jen.Lit("stmnt"), jen.Id("sql")).Op(".").Line().
			Id("Interface").Call(jen.Lit("args"), jen.Id("args")).
			Dot("Err").Call(jen.Id("err")).
			Dot("Msg").Call(jen.Lit("")),
	)
}
