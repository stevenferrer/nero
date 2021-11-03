package sort_test

import (
	"testing"

	"github.com/stevenferrer/nero/sort"
	"github.com/stretchr/testify/assert"
)

func TestDirectionStrings(t *testing.T) {
	tests := []struct {
		direction sort.Direction
		wantStr,
		wantDesc string
	}{
		{
			direction: sort.Asc,
			wantStr:   "Asc",
			wantDesc:  "ascending",
		},
		{
			direction: sort.Desc,
			wantStr:   "Desc",
			wantDesc:  "descending",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.wantStr, tc.direction.String())
		assert.Equal(t, tc.wantDesc, tc.direction.Desc())
	}
}
