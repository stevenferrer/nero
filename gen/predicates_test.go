package gen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sf9v/nero/example"
	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newPredicates(t *testing.T) {
	t.Run("all columns with predicates", func(t *testing.T) {
		schema, err := gen.BuildSchema(new(example.User))
		require.NoError(t, err)
		require.NotNil(t, schema)

		predicates := newPredicates(schema)
		expect := strings.TrimSpace(`
type PredFunc func(*comparison.Predicates)

func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

func NameEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Eq,
			Val: name,
		})
	}
}

func NameNotEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotEq,
			Val: name,
		})
	}
}

func NameGt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Gt,
			Val: name,
		})
	}
}

func NameGtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.GtOrEq,
			Val: name,
		})
	}
}

func NameLt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Lt,
			Val: name,
		})
	}
}

func NameLtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.LtOrEq,
			Val: name,
		})
	}
}

func GroupEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Eq,
			Val: group,
		})
	}
}

func GroupNotEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.NotEq,
			Val: group,
		})
	}
}

func GroupGt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Gt,
			Val: group,
		})
	}
}

func GroupGtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.GtOrEq,
			Val: group,
		})
	}
}

func GroupLt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Lt,
			Val: group,
		})
	}
}

func GroupLtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.LtOrEq,
			Val: group,
		})
	}
}

func UpdatedAtEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Eq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtNotEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotEq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtGt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Gt,
			Val: updatedAt,
		})
	}
}

func UpdatedAtGtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.GtOrEq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtLt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Lt,
			Val: updatedAt,
		})
	}
}

func UpdatedAtLtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.LtOrEq,
			Val: updatedAt,
		})
	}
}

func CreatedAtEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Eq,
			Val: createdAt,
		})
	}
}

func CreatedAtNotEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotEq,
			Val: createdAt,
		})
	}
}

func CreatedAtGt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Gt,
			Val: createdAt,
		})
	}
}

func CreatedAtGtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.GtOrEq,
			Val: createdAt,
		})
	}
}

func CreatedAtLt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Lt,
			Val: createdAt,
		})
	}
}

func CreatedAtLtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.LtOrEq,
			Val: createdAt,
		})
	}
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
		assert.Equal(t, expect, got)
	})

	t.Run("some columns without predicates", func(t *testing.T) {
		schema, err := gen.BuildSchema(new(example.NoPreds))
		require.NoError(t, err)
		require.NotNil(t, schema)

		predicates := newPredicates(schema)
		expect := strings.TrimSpace(`
type PredFunc func(*comparison.Predicates)

func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

func AEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Eq,
			Val: a,
		})
	}
}

func ANotEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.NotEq,
			Val: a,
		})
	}
}

func AGt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Gt,
			Val: a,
		})
	}
}

func AGtOrEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.GtOrEq,
			Val: a,
		})
	}
}

func ALt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Lt,
			Val: a,
		})
	}
}

func ALtOrEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.LtOrEq,
			Val: a,
		})
	}
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
		assert.Equal(t, expect, got)
	})
}
