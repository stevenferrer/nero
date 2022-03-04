package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevenferrer/nero/predicate"
)

func TestOperatorStrings(t *testing.T) {
	tests := []struct {
		op predicate.Operator
		wantStr,
		wantDesc string
	}{
		{
			op:       predicate.Eq,
			wantStr:  "Eq",
			wantDesc: "equal",
		},
		{
			op:       predicate.NotEq,
			wantStr:  "NotEq",
			wantDesc: "not equal",
		},
		{
			op:       predicate.Gt,
			wantStr:  "Gt",
			wantDesc: "greater than",
		},
		{
			op:       predicate.GtOrEq,
			wantStr:  "GtOrEq",
			wantDesc: "greater than or equal",
		},
		{
			op:       predicate.Lt,
			wantStr:  "Lt",
			wantDesc: "less than",
		},
		{
			op:       predicate.LtOrEq,
			wantStr:  "LtOrEq",
			wantDesc: "less than or equal",
		},
		{
			op:       predicate.IsNull,
			wantStr:  "IsNull",
			wantDesc: "is null",
		},
		{
			op:       predicate.IsNotNull,
			wantStr:  "IsNotNull",
			wantDesc: "is not null",
		},
		{
			op:       predicate.In,
			wantStr:  "In",
			wantDesc: "in",
		},
		{
			op:       predicate.NotIn,
			wantStr:  "NotIn",
			wantDesc: "not in",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.wantStr, tc.op.String())
		assert.Equal(t, tc.wantDesc, tc.op.Desc())
	}
}
