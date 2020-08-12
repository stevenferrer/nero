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
// PredFunc is the predicate function type
type PredFunc func(*comparison.Predicates)

// IDEq applies a equal operator to id
func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

// IDNotEq applies a not equal operator to id
func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

// IDGt applies a greater than operator to id
func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

// IDGtOrEq applies a greater than or equal operator to id
func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

// IDLt applies a less than operator to id
func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

// IDLtOrEq applies a less than or equal operator to id
func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

// NameEq applies a equal operator to name
func NameEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Eq,
			Val: name,
		})
	}
}

// NameNotEq applies a not equal operator to name
func NameNotEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotEq,
			Val: name,
		})
	}
}

// NameGt applies a greater than operator to name
func NameGt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Gt,
			Val: name,
		})
	}
}

// NameGtOrEq applies a greater than or equal operator to name
func NameGtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.GtOrEq,
			Val: name,
		})
	}
}

// NameLt applies a less than operator to name
func NameLt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Lt,
			Val: name,
		})
	}
}

// NameLtOrEq applies a less than or equal operator to name
func NameLtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.LtOrEq,
			Val: name,
		})
	}
}

// GroupEq applies a equal operator to group
func GroupEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Eq,
			Val: group,
		})
	}
}

// GroupNotEq applies a not equal operator to group
func GroupNotEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.NotEq,
			Val: group,
		})
	}
}

// GroupGt applies a greater than operator to group
func GroupGt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Gt,
			Val: group,
		})
	}
}

// GroupGtOrEq applies a greater than or equal operator to group
func GroupGtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.GtOrEq,
			Val: group,
		})
	}
}

// GroupLt applies a less than operator to group
func GroupLt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Lt,
			Val: group,
		})
	}
}

// GroupLtOrEq applies a less than or equal operator to group
func GroupLtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.LtOrEq,
			Val: group,
		})
	}
}

// UpdatedAtEq applies a equal operator to updatedAt
func UpdatedAtEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Eq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtNotEq applies a not equal operator to updatedAt
func UpdatedAtNotEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotEq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtGt applies a greater than operator to updatedAt
func UpdatedAtGt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Gt,
			Val: updatedAt,
		})
	}
}

// UpdatedAtGtOrEq applies a greater than or equal operator to updatedAt
func UpdatedAtGtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.GtOrEq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtLt applies a less than operator to updatedAt
func UpdatedAtLt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Lt,
			Val: updatedAt,
		})
	}
}

// UpdatedAtLtOrEq applies a less than or equal operator to updatedAt
func UpdatedAtLtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.LtOrEq,
			Val: updatedAt,
		})
	}
}

// CreatedAtEq applies a equal operator to createdAt
func CreatedAtEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Eq,
			Val: createdAt,
		})
	}
}

// CreatedAtNotEq applies a not equal operator to createdAt
func CreatedAtNotEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotEq,
			Val: createdAt,
		})
	}
}

// CreatedAtGt applies a greater than operator to createdAt
func CreatedAtGt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Gt,
			Val: createdAt,
		})
	}
}

// CreatedAtGtOrEq applies a greater than or equal operator to createdAt
func CreatedAtGtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.GtOrEq,
			Val: createdAt,
		})
	}
}

// CreatedAtLt applies a less than operator to createdAt
func CreatedAtLt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Lt,
			Val: createdAt,
		})
	}
}

// CreatedAtLtOrEq applies a less than or equal operator to createdAt
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
// PredFunc is the predicate function type
type PredFunc func(*comparison.Predicates)

// IDEq applies a equal operator to id
func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

// IDNotEq applies a not equal operator to id
func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

// IDGt applies a greater than operator to id
func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

// IDGtOrEq applies a greater than or equal operator to id
func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

// IDLt applies a less than operator to id
func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

// IDLtOrEq applies a less than or equal operator to id
func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

// AEq applies a equal operator to a
func AEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Eq,
			Val: a,
		})
	}
}

// ANotEq applies a not equal operator to a
func ANotEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.NotEq,
			Val: a,
		})
	}
}

// AGt applies a greater than operator to a
func AGt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Gt,
			Val: a,
		})
	}
}

// AGtOrEq applies a greater than or equal operator to a
func AGtOrEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.GtOrEq,
			Val: a,
		})
	}
}

// ALt applies a less than operator to a
func ALt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Lt,
			Val: a,
		})
	}
}

// ALtOrEq applies a less than or equal operator to a
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
