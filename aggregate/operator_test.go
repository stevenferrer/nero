package aggregate_test

import (
	"testing"

	"github.com/stevenferrer/nero/aggregate"
	"github.com/stretchr/testify/assert"
)

func TestOperatorStrings(t *testing.T) {
	tests := []struct {
		op aggregate.Operator
		wantStr,
		wantDesc string
	}{
		{
			op:       aggregate.Avg,
			wantStr:  "Avg",
			wantDesc: "average",
		},
		{
			op:       aggregate.Count,
			wantStr:  "Count",
			wantDesc: "count",
		},
		{
			op:       aggregate.Max,
			wantStr:  "Max",
			wantDesc: "max",
		},
		{
			op:       aggregate.Min,
			wantStr:  "Min",
			wantDesc: "min",
		},
		{
			op:       aggregate.Sum,
			wantStr:  "Sum",
			wantDesc: "sum",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.wantStr, tc.op.String())
		assert.Equal(t, tc.wantDesc, tc.op.Desc())
	}
}
