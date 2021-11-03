package playerrepo_test

import (
	"testing"
	"time"

	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/test/integration/player"
	"github.com/stevenferrer/nero/test/integration/playerrepo"
	"github.com/stretchr/testify/assert"
)

func TestPredicate(t *testing.T) {
	now := time.Now()
	tests := []struct {
		predFunc comparison.PredFunc
		want     *comparison.Predicate
	}{
		// id
		{
			predFunc: playerrepo.IDEq("1"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldID.String(),
				Op:    comparison.Eq,
				Arg:   "1",
			},
		},
		{
			predFunc: playerrepo.IDNotEq("1"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldID.String(),
				Op:    comparison.NotEq,
				Arg:   "1",
			},
		},
		{
			predFunc: playerrepo.IDIn("1"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldID.String(),
				Op:    comparison.In,
				Arg:   []interface{}{"1"},
			},
		},
		{
			predFunc: playerrepo.IDNotIn("1"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldID.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{"1"},
			},
		},

		// email
		{
			predFunc: playerrepo.EmailEq("me@me.io"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldEmail.String(),
				Op:    comparison.Eq,
				Arg:   "me@me.io",
			},
		},
		{
			predFunc: playerrepo.EmailNotEq("me@me.io"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldEmail.String(),
				Op:    comparison.NotEq,
				Arg:   "me@me.io",
			},
		},
		{
			predFunc: playerrepo.EmailIn("me@me.io"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldEmail.String(),
				Op:    comparison.In,
				Arg:   []interface{}{"me@me.io"},
			},
		},
		{
			predFunc: playerrepo.EmailNotIn("me@me.io"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldEmail.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{"me@me.io"},
			},
		},
		// name
		{
			predFunc: playerrepo.NameEq("me"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldName.String(),
				Op:    comparison.Eq,
				Arg:   "me",
			},
		},
		{
			predFunc: playerrepo.NameNotEq("me"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldName.String(),
				Op:    comparison.NotEq,
				Arg:   "me",
			},
		},
		{
			predFunc: playerrepo.NameIn("me"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldName.String(),
				Op:    comparison.In,
				Arg:   []interface{}{"me"},
			},
		},
		{
			predFunc: playerrepo.NameNotIn("me"),
			want: &comparison.Predicate{
				Field: playerrepo.FieldName.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{"me"},
			},
		},

		// age
		{
			predFunc: playerrepo.AgeEq(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.Eq,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeNotEq(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.NotEq,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeGt(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.Gt,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeGtOrEq(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.GtOrEq,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeLt(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.Lt,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeLtOrEq(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.LtOrEq,
				Arg:   18,
			},
		},
		{
			predFunc: playerrepo.AgeNotEq(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.NotEq,
				Arg:   18,
			},
		},

		{
			predFunc: playerrepo.AgeIn(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.In,
				Arg:   []interface{}{18},
			},
		},
		{
			predFunc: playerrepo.AgeNotIn(18),
			want: &comparison.Predicate{
				Field: playerrepo.FieldAge.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{18},
			},
		},

		// race
		{
			predFunc: playerrepo.RaceEq(player.RaceHuman),
			want: &comparison.Predicate{
				Field: playerrepo.FieldRace.String(),
				Op:    comparison.Eq,
				Arg:   player.RaceHuman,
			},
		},
		{
			predFunc: playerrepo.RaceNotEq(player.RaceHuman),
			want: &comparison.Predicate{
				Field: playerrepo.FieldRace.String(),
				Op:    comparison.NotEq,
				Arg:   player.RaceHuman,
			},
		},
		{
			predFunc: playerrepo.RaceIn(player.RaceHuman),
			want: &comparison.Predicate{
				Field: playerrepo.FieldRace.String(),
				Op:    comparison.In,
				Arg:   []interface{}{player.RaceHuman},
			},
		},
		{
			predFunc: playerrepo.RaceNotIn(player.RaceHuman),
			want: &comparison.Predicate{
				Field: playerrepo.FieldRace.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{player.RaceHuman},
			},
		},

		// updated at
		{
			predFunc: playerrepo.UpdatedAtEq(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.Eq,
				Arg:   &now,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtNotEq(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.NotEq,
				Arg:   &now,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIsNull(),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.IsNull,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIsNotNull(),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.IsNotNull,
			},
		},
		{
			predFunc: playerrepo.UpdatedAtIn(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.In,
				Arg:   []interface{}{&now},
			},
		},
		{
			predFunc: playerrepo.UpdatedAtNotIn(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{&now},
			},
		},

		// created at
		{
			predFunc: playerrepo.CreatedAtEq(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.Eq,
				Arg:   &now,
			},
		},
		{
			predFunc: playerrepo.CreatedAtNotEq(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.NotEq,
				Arg:   &now,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIsNull(),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.IsNull,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIsNotNull(),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.IsNotNull,
			},
		},
		{
			predFunc: playerrepo.CreatedAtIn(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.In,
				Arg:   []interface{}{&now},
			},
		},
		{
			predFunc: playerrepo.CreatedAtNotIn(&now),
			want: &comparison.Predicate{
				Field: playerrepo.FieldCreatedAt.String(),
				Op:    comparison.NotIn,
				Arg:   []interface{}{&now},
			},
		},

		// field to field comparison
		{
			predFunc: playerrepo.FieldXEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.Eq,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXNotEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.NotEq,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},

		{
			predFunc: playerrepo.FieldXGtFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.Gt,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXGtOrEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.GtOrEq,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},

		{
			predFunc: playerrepo.FieldXLtFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),
			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.Lt,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},
		{
			predFunc: playerrepo.FieldXLtOrEqFieldY(
				playerrepo.FieldUpdatedAt,
				playerrepo.FieldCreatedAt,
			),

			want: &comparison.Predicate{
				Field: playerrepo.FieldUpdatedAt.String(),
				Op:    comparison.LtOrEq,
				Arg:   playerrepo.FieldCreatedAt,
			},
		},
	}

	for _, tc := range tests {
		got := tc.predFunc([]*comparison.Predicate{})[0]
		assert.Equal(t, tc.want.Field, got.Field)
		assert.Equal(t, tc.want.Arg, got.Arg)
		assert.Equal(t, tc.want.Op, got.Op)
	}
}
