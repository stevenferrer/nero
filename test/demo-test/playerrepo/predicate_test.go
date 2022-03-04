package playerrepo_test

import (
	"testing"
	"time"

	comparison "github.com/stevenferrer/nero/predicate"
	"github.com/stevenferrer/nero/test/demo-test/playerpkg"
	"github.com/stevenferrer/nero/test/demo-test/playerrepo"
	"github.com/stretchr/testify/assert"
)

func TestPredicate(t *testing.T) {
	now := time.Now()
	tests := []struct {
		predFunc comparison.Func
		want     comparison.Predicate
	}{
		// id
		{
			predFunc: playerrepo.IDEq("1"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldID.String(),
				Operator: comparison.Eq,
				Argument: "1",
			},
		},
		{
			predFunc: playerrepo.IDNotEq("1"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldID.String(),
				Operator: comparison.NotEq,
				Argument: "1",
			},
		},
		{
			predFunc: playerrepo.IDIn("1"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldID.String(),
				Operator: comparison.In,
				Argument: []interface{}{"1"},
			},
		},
		{
			predFunc: playerrepo.IDNotIn("1"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldID.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{"1"},
			},
		},

		// email
		{
			predFunc: playerrepo.EmailEq("me@me.io"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldEmail.String(),
				Operator: comparison.Eq,
				Argument: "me@me.io",
			},
		},
		{
			predFunc: playerrepo.EmailNotEq("me@me.io"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldEmail.String(),
				Operator: comparison.NotEq,
				Argument: "me@me.io",
			},
		},
		{
			predFunc: playerrepo.EmailIn("me@me.io"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldEmail.String(),
				Operator: comparison.In,
				Argument: []interface{}{"me@me.io"},
			},
		},
		{
			predFunc: playerrepo.EmailNotIn("me@me.io"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldEmail.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{"me@me.io"},
			},
		},
		// name
		{
			predFunc: playerrepo.NameEq("me"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldName.String(),
				Operator: comparison.Eq,
				Argument: "me",
			},
		},
		{
			predFunc: playerrepo.NameNotEq("me"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldName.String(),
				Operator: comparison.NotEq,
				Argument: "me",
			},
		},
		{
			predFunc: playerrepo.NameIn("me"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldName.String(),
				Operator: comparison.In,
				Argument: []interface{}{"me"},
			},
		},
		{
			predFunc: playerrepo.NameNotIn("me"),
			want: comparison.Predicate{
				Field:    playerrepo.FieldName.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{"me"},
			},
		},

		// age
		{
			predFunc: playerrepo.AgeEq(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.Eq,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeNotEq(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.NotEq,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeGt(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.Gt,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeGtOrEq(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.GtOrEq,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeLt(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.Lt,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeLtOrEq(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.LtOrEq,
				Argument: 18,
			},
		},
		{
			predFunc: playerrepo.AgeNotEq(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.NotEq,
				Argument: 18,
			},
		},

		{
			predFunc: playerrepo.AgeIn(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.In,
				Argument: []interface{}{18},
			},
		},
		{
			predFunc: playerrepo.AgeNotIn(18),
			want: comparison.Predicate{
				Field:    playerrepo.FieldAge.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{18},
			},
		},

		// race
		{
			predFunc: playerrepo.RaceEq(playerpkg.RaceHuman),
			want: comparison.Predicate{
				Field:    playerrepo.FieldRace.String(),
				Operator: comparison.Eq,
				Argument: playerpkg.RaceHuman,
			},
		},
		{
			predFunc: playerrepo.RaceNotEq(playerpkg.RaceHuman),
			want: comparison.Predicate{
				Field:    playerrepo.FieldRace.String(),
				Operator: comparison.NotEq,
				Argument: playerpkg.RaceHuman,
			},
		},
		{
			predFunc: playerrepo.RaceIn(playerpkg.RaceHuman),
			want: comparison.Predicate{
				Field:    playerrepo.FieldRace.String(),
				Operator: comparison.In,
				Argument: []interface{}{playerpkg.RaceHuman},
			},
		},
		{
			predFunc: playerrepo.RaceNotIn(playerpkg.RaceHuman),
			want: comparison.Predicate{
				Field:    playerrepo.FieldRace.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{playerpkg.RaceHuman},
			},
		},

		// updated at
		{
			predFunc: playerrepo.UpdatedAtEq(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.Eq,
				Argument: &now,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtNotEq(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.NotEq,
				Argument: &now,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIsNull(),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.IsNull,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIsNotNull(),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.IsNotNull,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIn(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.In,
				Argument: []interface{}{&now},
			},
		},
		{
			predFunc: playerrepo.UpdatedAtNotIn(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{&now},
			},
		},

		// created at
		{
			predFunc: playerrepo.CreatedAtEq(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.Eq,
				Argument: &now,
			},
		},
		{
			predFunc: playerrepo.CreatedAtNotEq(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.NotEq,
				Argument: &now,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIsNull(),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.IsNull,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIsNotNull(),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.IsNotNull,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIn(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.In,
				Argument: []interface{}{&now},
			},
		},
		{
			predFunc: playerrepo.CreatedAtNotIn(&now),
			want: comparison.Predicate{
				Field:    playerrepo.FieldCreatedAt.String(),
				Operator: comparison.NotIn,
				Argument: []interface{}{&now},
			},
		},

		// field to field comparison
		{
			predFunc: playerrepo.FieldXEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.Eq,
				Argument: playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXNotEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.NotEq,
				Argument: playerrepo.FieldCreatedAt,
			},
		},

		{
			predFunc: playerrepo.FieldXGtFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.Gt,
				Argument: playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXGtOrEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.GtOrEq,
				Argument: playerrepo.FieldCreatedAt,
			},
		},

		{
			predFunc: playerrepo.FieldXLtFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.Lt,
				Argument: playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXLtOrEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),

			want: comparison.Predicate{
				Field:    playerrepo.FieldUpdatedAt.String(),
				Operator: comparison.LtOrEq,
				Argument: playerrepo.FieldCreatedAt,
			},
		},
	}

	for _, tc := range tests {
		got := tc.predFunc([]comparison.Predicate{})[0]
		assert.Equal(t, tc.want.Field, got.Field)
		assert.Equal(t, tc.want.Argument, got.Argument)
		assert.Equal(t, tc.want.Operator, got.Operator)
	}
}
