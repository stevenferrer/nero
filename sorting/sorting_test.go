package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevenferrer/nero/sorting"
)

func TestDirectionStrings(t *testing.T) {
	tests := []struct {
		direction sorting.Direction
		wantStr,
		wantDesc string
	}{
		{
			direction: sorting.Asc,
			wantStr:   "Asc",
			wantDesc:  "ascending",
		},
		{
			direction: sorting.Desc,
			wantStr:   "Desc",
			wantDesc:  "descending",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.wantStr, tc.direction.String())
		assert.Equal(t, tc.wantDesc, tc.direction.Desc())
	}
}
