package repository

import (
	"github.com/sf9v/nero/comparison"

	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/test/integration/user"
	"time"
)

type PredFunc func(*comparison.Predicates)

func IDEq(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Val: id,
		})
	}
}

func IDNotEq(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Val: id,
		})
	}
}

func IDGt(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Val: id,
		})
	}
}

func IDGtOrEq(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Val: id,
		})
	}
}

func IDLt(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Val: id,
		})
	}
}

func IDLtOrEq(id string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Val: id,
		})
	}
}

func IDIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.IsNull,
		})
	}
}

func IDIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.IsNotNull,
		})
	}
}

func IDIn(ids ...string) PredFunc {
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

func IDNotIn(ids ...string) PredFunc {
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

func UIDEq(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.Eq,
			Val: uid,
		})
	}
}

func UIDNotEq(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.NotEq,
			Val: uid,
		})
	}
}

func UIDGt(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.Gt,
			Val: uid,
		})
	}
}

func UIDGtOrEq(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.GtOrEq,
			Val: uid,
		})
	}
}

func UIDLt(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.Lt,
			Val: uid,
		})
	}
}

func UIDLtOrEq(uid ksuid.KSUID) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.LtOrEq,
			Val: uid,
		})
	}
}

func UIDIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.IsNull,
		})
	}
}

func UIDIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.IsNotNull,
		})
	}
}

func UIDIn(uids ...ksuid.KSUID) PredFunc {
	vals := []interface{}{}
	for _, v := range uids {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func UIDNotIn(uids ...ksuid.KSUID) PredFunc {
	vals := []interface{}{}
	for _, v := range uids {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "uid",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

func EmailEq(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.Eq,
			Val: email,
		})
	}
}

func EmailNotEq(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.NotEq,
			Val: email,
		})
	}
}

func EmailGt(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.Gt,
			Val: email,
		})
	}
}

func EmailGtOrEq(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.GtOrEq,
			Val: email,
		})
	}
}

func EmailLt(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.Lt,
			Val: email,
		})
	}
}

func EmailLtOrEq(email string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.LtOrEq,
			Val: email,
		})
	}
}

func EmailIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.IsNull,
		})
	}
}

func EmailIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.IsNotNull,
		})
	}
}

func EmailIn(emails ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range emails {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func EmailNotIn(emails ...string) PredFunc {
	vals := []interface{}{}
	for _, v := range emails {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "email",
			Op:  comparison.NotIn,
			Val: vals,
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

func NameIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.IsNull,
		})
	}
}

func NameIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.IsNotNull,
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

func AgeEq(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.Eq,
			Val: age,
		})
	}
}

func AgeNotEq(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.NotEq,
			Val: age,
		})
	}
}

func AgeGt(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.Gt,
			Val: age,
		})
	}
}

func AgeGtOrEq(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.GtOrEq,
			Val: age,
		})
	}
}

func AgeLt(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.Lt,
			Val: age,
		})
	}
}

func AgeLtOrEq(age int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.LtOrEq,
			Val: age,
		})
	}
}

func AgeIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.IsNull,
		})
	}
}

func AgeIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.IsNotNull,
		})
	}
}

func AgeIn(ages ...int) PredFunc {
	vals := []interface{}{}
	for _, v := range ages {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func AgeNotIn(ages ...int) PredFunc {
	vals := []interface{}{}
	for _, v := range ages {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "age",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}

func GroupEq(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.Eq,
			Val: group,
		})
	}
}

func GroupNotEq(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.NotEq,
			Val: group,
		})
	}
}

func GroupGt(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.Gt,
			Val: group,
		})
	}
}

func GroupGtOrEq(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.GtOrEq,
			Val: group,
		})
	}
}

func GroupLt(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.Lt,
			Val: group,
		})
	}
}

func GroupLtOrEq(group user.Group) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.LtOrEq,
			Val: group,
		})
	}
}

func GroupIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.IsNull,
		})
	}
}

func GroupIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.IsNotNull,
		})
	}
}

func GroupIn(groups ...user.Group) PredFunc {
	vals := []interface{}{}
	for _, v := range groups {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.In,
			Val: vals,
		})
	}
}

func GroupNotIn(groups ...user.Group) PredFunc {
	vals := []interface{}{}
	for _, v := range groups {
		vals = append(vals, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "group",
			Op:  comparison.NotIn,
			Val: vals,
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

func UpdatedAtIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.IsNull,
		})
	}
}

func UpdatedAtIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.IsNotNull,
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
			Col: "updatedAt",
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
			Col: "updatedAt",
			Op:  comparison.NotIn,
			Val: vals,
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

func CreatedAtIsNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.IsNull,
		})
	}
}

func CreatedAtIsNotNull() PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.IsNotNull,
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
			Col: "createdAt",
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
			Col: "createdAt",
			Op:  comparison.NotIn,
			Val: vals,
		})
	}
}
