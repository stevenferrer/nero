package comparison_test

import (
	"testing"

	"github.com/stevenferrer/nero/comparison"
	"github.com/stretchr/testify/assert"
)

func TestOperatorStrings(t *testing.T) {
	tests := []struct {
		op comparison.Operator
		wantStr,
		wantDesc string
	}{
		{
			op:       comparison.Eq,
			wantStr:  "Eq",
			wantDesc: "equal",
		},
		{
			op:       comparison.NotEq,
			wantStr:  "NotEq",
			wantDesc: "not equal",
		},
		{
			op:       comparison.Gt,
			wantStr:  "Gt",
			wantDesc: "greater than",
		},
		{
			op:       comparison.GtOrEq,
			wantStr:  "GtOrEq",
			wantDesc: "greater than or equal",
		},
		{
			op:       comparison.Lt,
			wantStr:  "Lt",
			wantDesc: "less than",
		},
		{
			op:       comparison.LtOrEq,
			wantStr:  "LtOrEq",
			wantDesc: "less than or equal",
		},
		{
			op:       comparison.IsNull,
			wantStr:  "IsNull",
			wantDesc: "is null",
		},
		{
			op:       comparison.IsNotNull,
			wantStr:  "IsNotNull",
			wantDesc: "is not null",
		},
		{
			op:       comparison.In,
			wantStr:  "In",
			wantDesc: "in",
		},
		{
			op:       comparison.NotIn,
			wantStr:  "NotIn",
			wantDesc: "not in",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.wantStr, tc.op.String())
		assert.Equal(t, tc.wantDesc, tc.op.Desc())
	}
}
