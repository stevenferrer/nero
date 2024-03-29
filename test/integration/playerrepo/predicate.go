// Code generated by nero, DO NOT EDIT.
package playerrepo

import (
	"time"

	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/test/integration/player"
)

// IDEq equal operator on ID field
func IDEq(id string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "id",
			Op:    comparison.Eq,
			Arg:   id,
		})
	}
}

// IDNotEq not equal operator on ID field
func IDNotEq(id string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "id",
			Op:    comparison.NotEq,
			Arg:   id,
		})
	}
}

// IDIn in operator on ID field
func IDIn(ids ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "id",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// IDNotIn not in operator on ID field
func IDNotIn(ids ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "id",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// EmailEq equal operator on Email field
func EmailEq(email string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "email",
			Op:    comparison.Eq,
			Arg:   email,
		})
	}
}

// EmailNotEq not equal operator on Email field
func EmailNotEq(email string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "email",
			Op:    comparison.NotEq,
			Arg:   email,
		})
	}
}

// EmailIn in operator on Email field
func EmailIn(emails ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range emails {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "email",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// EmailNotIn not in operator on Email field
func EmailNotIn(emails ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range emails {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "email",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// NameEq equal operator on Name field
func NameEq(name string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "name",
			Op:    comparison.Eq,
			Arg:   name,
		})
	}
}

// NameNotEq not equal operator on Name field
func NameNotEq(name string) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "name",
			Op:    comparison.NotEq,
			Arg:   name,
		})
	}
}

// NameIn in operator on Name field
func NameIn(names ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "name",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// NameNotIn not in operator on Name field
func NameNotIn(names ...string) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "name",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// AgeEq equal operator on Age field
func AgeEq(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.Eq,
			Arg:   age,
		})
	}
}

// AgeNotEq not equal operator on Age field
func AgeNotEq(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.NotEq,
			Arg:   age,
		})
	}
}

// AgeGt greater than operator on Age field
func AgeGt(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.Gt,
			Arg:   age,
		})
	}
}

// AgeGtOrEq greater than or equal operator on Age field
func AgeGtOrEq(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.GtOrEq,
			Arg:   age,
		})
	}
}

// AgeLt less than operator on Age field
func AgeLt(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.Lt,
			Arg:   age,
		})
	}
}

// AgeLtOrEq less than or equal operator on Age field
func AgeLtOrEq(age int) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.LtOrEq,
			Arg:   age,
		})
	}
}

// AgeIn in operator on Age field
func AgeIn(ages ...int) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range ages {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// AgeNotIn not in operator on Age field
func AgeNotIn(ages ...int) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range ages {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "age",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// RaceEq equal operator on Race field
func RaceEq(race player.Race) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "race",
			Op:    comparison.Eq,
			Arg:   race,
		})
	}
}

// RaceNotEq not equal operator on Race field
func RaceNotEq(race player.Race) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "race",
			Op:    comparison.NotEq,
			Arg:   race,
		})
	}
}

// RaceIn in operator on Race field
func RaceIn(races ...player.Race) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range races {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "race",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// RaceNotIn not in operator on Race field
func RaceNotIn(races ...player.Race) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range races {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "race",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// UpdatedAtEq equal operator on UpdatedAt field
func UpdatedAtEq(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.Eq,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtNotEq not equal operator on UpdatedAt field
func UpdatedAtNotEq(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.NotEq,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtGt greater than operator on UpdatedAt field
func UpdatedAtGt(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.Gt,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtGtOrEq greater than or equal operator on UpdatedAt field
func UpdatedAtGtOrEq(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.GtOrEq,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtLt less than operator on UpdatedAt field
func UpdatedAtLt(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.Lt,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtLtOrEq less than or equal operator on UpdatedAt field
func UpdatedAtLtOrEq(updatedAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.LtOrEq,
			Arg:   updatedAt,
		})
	}
}

// UpdatedAtIsNull is null operator on UpdatedAt field
func UpdatedAtIsNull() comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.IsNull,
		})
	}
}

// UpdatedAtIsNotNull is not null operator on UpdatedAt field
func UpdatedAtIsNotNull() comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.IsNotNull,
		})
	}
}

// UpdatedAtIn in operator on UpdatedAt field
func UpdatedAtIn(updatedAts ...*time.Time) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range updatedAts {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// UpdatedAtNotIn not in operator on UpdatedAt field
func UpdatedAtNotIn(updatedAts ...*time.Time) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range updatedAts {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "updated_at",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// CreatedAtEq equal operator on CreatedAt field
func CreatedAtEq(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.Eq,
			Arg:   createdAt,
		})
	}
}

// CreatedAtNotEq not equal operator on CreatedAt field
func CreatedAtNotEq(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.NotEq,
			Arg:   createdAt,
		})
	}
}

// CreatedAtGt greater than operator on CreatedAt field
func CreatedAtGt(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.Gt,
			Arg:   createdAt,
		})
	}
}

// CreatedAtGtOrEq greater than or equal operator on CreatedAt field
func CreatedAtGtOrEq(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.GtOrEq,
			Arg:   createdAt,
		})
	}
}

// CreatedAtLt less than operator on CreatedAt field
func CreatedAtLt(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.Lt,
			Arg:   createdAt,
		})
	}
}

// CreatedAtLtOrEq less than or equal operator on CreatedAt field
func CreatedAtLtOrEq(createdAt *time.Time) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.LtOrEq,
			Arg:   createdAt,
		})
	}
}

// CreatedAtIsNull is null operator on CreatedAt field
func CreatedAtIsNull() comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.IsNull,
		})
	}
}

// CreatedAtIsNotNull is not null operator on CreatedAt field
func CreatedAtIsNotNull() comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.IsNotNull,
		})
	}
}

// CreatedAtIn in operator on CreatedAt field
func CreatedAtIn(createdAts ...*time.Time) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range createdAts {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.In,
			Arg:   args,
		})
	}
}

// CreatedAtNotIn not in operator on CreatedAt field
func CreatedAtNotIn(createdAts ...*time.Time) comparison.PredFunc {
	args := []interface{}{}
	for _, v := range createdAts {
		args = append(args, v)
	}

	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: "created_at",
			Op:    comparison.NotIn,
			Arg:   args,
		})
	}
}

// FieldXEqFieldY fieldX equal fieldY
//
// fieldX and fieldY must be of the same type
func FieldXEqFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.Eq,
			Arg:   fieldY,
		})
	}
}

// FieldXNotEqFieldY fieldX not equal fieldY
//
// fieldX and fieldY must be of the same type
func FieldXNotEqFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.NotEq,
			Arg:   fieldY,
		})
	}
}

// FieldXGtFieldY fieldX greater than fieldY
//
// fieldX and fieldY must be of the same type
func FieldXGtFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.Gt,
			Arg:   fieldY,
		})
	}
}

// FieldXGtOrEqFieldY fieldX greater than or equal fieldY
//
// fieldX and fieldY must be of the same type
func FieldXGtOrEqFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.GtOrEq,
			Arg:   fieldY,
		})
	}
}

// FieldXLtFieldY fieldX less than fieldY
//
// fieldX and fieldY must be of the same type
func FieldXLtFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.Lt,
			Arg:   fieldY,
		})
	}
}

// FieldXLtOrEqFieldY fieldX less than or equal fieldY
//
// fieldX and fieldY must be of the same type
func FieldXLtOrEqFieldY(fieldX, fieldY Field) comparison.PredFunc {
	return func(preds []*comparison.Predicate) []*comparison.Predicate {
		return append(preds, &comparison.Predicate{
			Field: fieldX.String(),
			Op:    comparison.LtOrEq,
			Arg:   fieldY,
		})
	}
}
