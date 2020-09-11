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

// IDEq returns a/an equal predicate on id
func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

// IDNotEq returns a/an not equal predicate on id
func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

// IDGt returns a/an greater than predicate on id
func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

// IDGtOrEq returns a/an greater than or equal predicate on id
func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

// IDLt returns a/an less than predicate on id
func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

// IDLtOrEq returns a/an less than or equal predicate on id
func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

func IDIn(ids ...int64) PredFunc {
	vals := []interface{}{}
	for _, v := range ids {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func IDNotIn(ids ...int64) PredFunc {
	vals := []interface{}{}
	for _, v := range ids {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

// NameEq returns a/an equal predicate on name
func NameEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Eq,
			Val: name,
		})
	}
}

// NameNotEq returns a/an not equal predicate on name
func NameNotEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotEq,
			Val: name,
		})
	}
}

// NameGt returns a/an greater than predicate on name
func NameGt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Gt,
			Val: name,
		})
	}
}

// NameGtOrEq returns a/an greater than or equal predicate on name
func NameGtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.GtOrEq,
			Val: name,
		})
	}
}

// NameLt returns a/an less than predicate on name
func NameLt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Lt,
			Val: name,
		})
	}
}

// NameLtOrEq returns a/an less than or equal predicate on name
func NameLtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.LtOrEq,
			Val: name,
		})
	}
}

func NameIn(names ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range names {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func NameNotIn(names ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range names {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

// GroupEq returns a/an equal predicate on group
func GroupEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Eq,
			Val: group,
		})
	}
}

// GroupNotEq returns a/an not equal predicate on group
func GroupNotEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.NotEq,
			Val: group,
		})
	}
}

// GroupGt returns a/an greater than predicate on group
func GroupGt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Gt,
			Val: group,
		})
	}
}

// GroupGtOrEq returns a/an greater than or equal predicate on group
func GroupGtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.GtOrEq,
			Val: group,
		})
	}
}

// GroupLt returns a/an less than predicate on group
func GroupLt(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.Lt,
			Val: group,
		})
	}
}

// GroupLtOrEq returns a/an less than or equal predicate on group
func GroupLtOrEq(group string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.LtOrEq,
			Val: group,
		})
	}
}

func GroupIn(groups ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range groups {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func GroupNotIn(groups ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range groups {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group_res",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

// UpdatedAtEq returns a/an equal predicate on updatedAt
func UpdatedAtEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Eq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtNotEq returns a/an not equal predicate on updatedAt
func UpdatedAtNotEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotEq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtGt returns a/an greater than predicate on updatedAt
func UpdatedAtGt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Gt,
			Val: updatedAt,
		})
	}
}

// UpdatedAtGtOrEq returns a/an greater than or equal predicate on updatedAt
func UpdatedAtGtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.GtOrEq,
			Val: updatedAt,
		})
	}
}

// UpdatedAtLt returns a/an less than predicate on updatedAt
func UpdatedAtLt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Lt,
			Val: updatedAt,
		})
	}
}

// UpdatedAtLtOrEq returns a/an less than or equal predicate on updatedAt
func UpdatedAtLtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.LtOrEq,
			Val: updatedAt,
		})
	}
}

func UpdatedAtIn(updatedAts ...*time.Time) PredFunc {
	vals := []interface{}{}
	for _, v := range updatedAts {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func UpdatedAtNotIn(updatedAts ...*time.Time) PredFunc {
	vals := []interface{}{}
	for _, v := range updatedAts {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

// CreatedAtEq returns a/an equal predicate on createdAt
func CreatedAtEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Eq,
			Val: createdAt,
		})
	}
}

// CreatedAtNotEq returns a/an not equal predicate on createdAt
func CreatedAtNotEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotEq,
			Val: createdAt,
		})
	}
}

// CreatedAtGt returns a/an greater than predicate on createdAt
func CreatedAtGt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Gt,
			Val: createdAt,
		})
	}
}

// CreatedAtGtOrEq returns a/an greater than or equal predicate on createdAt
func CreatedAtGtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.GtOrEq,
			Val: createdAt,
		})
	}
}

// CreatedAtLt returns a/an less than predicate on createdAt
func CreatedAtLt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Lt,
			Val: createdAt,
		})
	}
}

// CreatedAtLtOrEq returns a/an less than or equal predicate on createdAt
func CreatedAtLtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.LtOrEq,
			Val: createdAt,
		})
	}
}

func CreatedAtIn(createdAts ...*time.Time) PredFunc {
	vals := []interface{}{}
	for _, v := range createdAts {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func CreatedAtNotIn(createdAts ...*time.Time) PredFunc {
	vals := []interface{}{}
	for _, v := range createdAts {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotIn,
			Val: vals,
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

// IDEq returns a/an equal predicate on id
func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

// IDNotEq returns a/an not equal predicate on id
func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

// IDGt returns a/an greater than predicate on id
func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

// IDGtOrEq returns a/an greater than or equal predicate on id
func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

// IDLt returns a/an less than predicate on id
func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

// IDLtOrEq returns a/an less than or equal predicate on id
func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

func IDIn(ids ...int64) PredFunc {
	vals := []interface{}{}
	for _, v := range ids {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func IDNotIn(ids ...int64) PredFunc {
	vals := []interface{}{}
	for _, v := range ids {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

// AEq returns a/an equal predicate on a
func AEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Eq,
			Val: a,
		})
	}
}

// ANotEq returns a/an not equal predicate on a
func ANotEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.NotEq,
			Val: a,
		})
	}
}

// AGt returns a/an greater than predicate on a
func AGt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Gt,
			Val: a,
		})
	}
}

// AGtOrEq returns a/an greater than or equal predicate on a
func AGtOrEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.GtOrEq,
			Val: a,
		})
	}
}

// ALt returns a/an less than predicate on a
func ALt(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.Lt,
			Val: a,
		})
	}
}

// ALtOrEq returns a/an less than or equal predicate on a
func ALtOrEq(a [1]string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.LtOrEq,
			Val: a,
		})
	}
}

func AIn(as ...[1]string) PredFunc {
	vals := []interface{}{}
	for _, v := range as {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func ANotIn(as ...[1]string) PredFunc {
	vals := []interface{}{}
	for _, v := range as {
		vals = append(vals, v)
	}
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "a",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}
`)

		got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
		assert.Equal(t, expect, got)
	})
}
