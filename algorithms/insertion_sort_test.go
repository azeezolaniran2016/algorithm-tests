package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortInt(t *testing.T) {
	table := []struct {
		in       []int
		expected []int
	}{
		{
			in:       []int{1, 5, 7, 8, 9, 3, 2, 5},
			expected: []int{1, 2, 3, 5, 5, 7, 8, 9},
		},
		{
			in:       []int{1},
			expected: []int{1},
		},
		{
			in:       []int{2, 1},
			expected: []int{1, 2},
		},
		{
			in:       []int{1, 2},
			expected: []int{1, 2},
		},
		{
			in:       []int{0, 0, -1, 10, 23, 100, -1000, 50, 34, 86, 78},
			expected: []int{-1000, -1, 0, 0, 10, 23, 34, 50, 78, 86, 100},
		},
	}

	for _, val := range table {
		assert.Equal(t, val.expected, InsertionSortInt(val.in))
	}
}
