package gen

import (
	"fmt"
	"strings"
	"testing"

	gen "github.com/sf9v/nero/gen/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newPredicates(t *testing.T) {
	schema, err := buildSchema(new(gen.Example))
	require.NoError(t, err)
	require.NotNil(t, schema)

	predicates := newPredicates(schema)
	expect := strings.TrimSpace(`
type PredicateFunc func(*predicate.Builder)

func IDEq(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.Eq,
			Val:   id,
		})
	}
}

func IDNotEq(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.NotEq,
			Val:   id,
		})
	}
}

func IDGt(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.Gt,
			Val:   id,
		})
	}
}

func IDGtOrEq(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.GtOrEq,
			Val:   id,
		})
	}
}

func IDLt(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.Lt,
			Val:   id,
		})
	}
}

func IDLtOrEq(id int64) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "id",
			Op:    predicate.LtOrEq,
			Val:   id,
		})
	}
}

func NameEq(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.Eq,
			Val:   name,
		})
	}
}

func NameNotEq(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.NotEq,
			Val:   name,
		})
	}
}

func NameGt(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.Gt,
			Val:   name,
		})
	}
}

func NameGtOrEq(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.GtOrEq,
			Val:   name,
		})
	}
}

func NameLt(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.Lt,
			Val:   name,
		})
	}
}

func NameLtOrEq(name string) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "name",
			Op:    predicate.LtOrEq,
			Val:   name,
		})
	}
}

func UpdatedAtEq(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.Eq,
			Val:   updatedAt,
		})
	}
}

func UpdatedAtNotEq(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.NotEq,
			Val:   updatedAt,
		})
	}
}

func UpdatedAtGt(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.Gt,
			Val:   updatedAt,
		})
	}
}

func UpdatedAtGtOrEq(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.GtOrEq,
			Val:   updatedAt,
		})
	}
}

func UpdatedAtLt(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.Lt,
			Val:   updatedAt,
		})
	}
}

func UpdatedAtLtOrEq(updatedAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "updated_at",
			Op:    predicate.LtOrEq,
			Val:   updatedAt,
		})
	}
}

func CreatedAtEq(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.Eq,
			Val:   createdAt,
		})
	}
}

func CreatedAtNotEq(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.NotEq,
			Val:   createdAt,
		})
	}
}

func CreatedAtGt(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.Gt,
			Val:   createdAt,
		})
	}
}

func CreatedAtGtOrEq(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.GtOrEq,
			Val:   createdAt,
		})
	}
}

func CreatedAtLt(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.Lt,
			Val:   createdAt,
		})
	}
}

func CreatedAtLtOrEq(createdAt *time.Time) PredicateFunc {
	return func(pb *predicate.Builder) {
		pb.Append(&predicate.Predicate{
			Field: "created_at",
			Op:    predicate.LtOrEq,
			Val:   createdAt,
		})
	}
}
`)

	got := strings.TrimSpace(fmt.Sprintf("%#v", predicates))
	assert.Equal(t, expect, got)
}
