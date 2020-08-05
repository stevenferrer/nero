package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
)

func Test_newPredicates(t *testing.T) {
	schema, err := gen.BuildSchema(new(example.User))
	require.NoError(t, err)
	require.NotNil(t, schema)

	predicates := newPredicates(schema)
	expect := strings.TrimSpace(`
type PredFunc func(*predicate.Predicates)

func IDEq(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.Eq,
			Val: id,
		})
	}
}

func IDNotEq(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.NotEq,
			Val: id,
		})
	}
}

func IDGt(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.Gt,
			Val: id,
		})
	}
}

func IDGtOrEq(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.GtOrEq,
			Val: id,
		})
	}
}

func IDLt(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.Lt,
			Val: id,
		})
	}
}

func IDLtOrEq(id int64) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "id",
			Op:  predicate.LtOrEq,
			Val: id,
		})
	}
}

func NameEq(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.Eq,
			Val: name,
		})
	}
}

func NameNotEq(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.NotEq,
			Val: name,
		})
	}
}

func NameGt(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.Gt,
			Val: name,
		})
	}
}

func NameGtOrEq(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.GtOrEq,
			Val: name,
		})
	}
}

func NameLt(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.Lt,
			Val: name,
		})
	}
}

func NameLtOrEq(name string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "name",
			Op:  predicate.LtOrEq,
			Val: name,
		})
	}
}

func GroupEq(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.Eq,
			Val: group,
		})
	}
}

func GroupNotEq(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.NotEq,
			Val: group,
		})
	}
}

func GroupGt(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.Gt,
			Val: group,
		})
	}
}

func GroupGtOrEq(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.GtOrEq,
			Val: group,
		})
	}
}

func GroupLt(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.Lt,
			Val: group,
		})
	}
}

func GroupLtOrEq(group string) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "group_res",
			Op:  predicate.LtOrEq,
			Val: group,
		})
	}
}

func UpdatedAtEq(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.Eq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtNotEq(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.NotEq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtGt(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.Gt,
			Val: updatedAt,
		})
	}
}

func UpdatedAtGtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.GtOrEq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtLt(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.Lt,
			Val: updatedAt,
		})
	}
}

func UpdatedAtLtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "updated_at",
			Op:  predicate.LtOrEq,
			Val: updatedAt,
		})
	}
}

func CreatedAtEq(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.Eq,
			Val: createdAt,
		})
	}
}

func CreatedAtNotEq(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.NotEq,
			Val: createdAt,
		})
	}
}

func CreatedAtGt(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.Gt,
			Val: createdAt,
		})
	}
}

func CreatedAtGtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.GtOrEq,
			Val: createdAt,
		})
	}
}

func CreatedAtLt(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.Lt,
			Val: createdAt,
		})
	}
}

func CreatedAtLtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *predicate.Predicates) {
		pb.Add(&predicate.Predicate{
			Col: "created_at",
			Op:  predicate.LtOrEq,
			Val: createdAt,
		})
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
	assert.Equal(t, expect, got)
}
