package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLeftChildNode(t *testing.T) {
	td := []struct {
		in, expected int
	}{
		{
			in:       0,
			expected: 1,
		},
		{
			in:       1,
			expected: 3,
		},
		{
			in:       3,
			expected: 7,
		},
	}

	for _, tc := range td {
		assert.Equal(t, tc.expected, getLeftChildNode(tc.in))
	}
}

func TestGetRightChildNode(t *testing.T) {
	td := []struct {
		in, expected int
	}{
		{
			in:       0,
			expected: 2,
		},
		{
			in:       1,
			expected: 4,
		},
		{
			in:       3,
			expected: 8,
		},
	}

	for _, tc := range td {
		assert.Equal(t, tc.expected, getRightChildNode(tc.in))
	}
}

func TestMakeMinHeap(t *testing.T) {
	td := []struct {
		in, expected []int
	}{
		{
			in:       []int{},
			expected: []int{},
		},
		{
			in:       []int{10},
			expected: []int{10},
		},
		{
			in:       []int{10, 1},
			expected: []int{1, 10},
		},
		{
			in:       []int{1, 10},
			expected: []int{1, 10},
		},
		{
			in:       []int{1, 10, 11, 45, 111, 203},
			expected: []int{1, 10, 11, 45, 111, 203},
		},
		{
			in:       []int{203, 111, 45, 11, 10, 1},
			expected: []int{1, 10, 45, 11, 111, 203},
		},
	}
	for _, tc := range td {
		assert.Equal(t, tc.expected, makeMinHeap(tc.in))
	}
}

func TestHeapSortInt(t *testing.T) {
	for _, tc := range sortTestCases {
		assert.Equal(t, tc.expected, HeapSortInt(tc.in))
	}
}
