package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortInt(t *testing.T) {
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
		assert.Equal(t, val.expected, MergeSortInt(val.in))
	}
}

func TestPopFirstElement(t *testing.T) {
	table := []struct {
		in       []int
		expected []int
	}{
		{
			in:       []int{},
			expected: []int{},
		},
		{
			in:       []int{1},
			expected: []int{},
		},
		{
			in:       []int{1, 2},
			expected: []int{2},
		},
		{
			in:       []int{0, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			in:       []int{0, 1, 1, 2, 3},
			expected: []int{1, 1, 2, 3},
		},
	}

	for _, tb := range table {
		assert.Equal(t, tb.expected, popFirstElement(tb.in), "input: %+v", tb.in)
	}
}

func TestMergeInt(t *testing.T) {
	table := []struct {
		aIn      []int
		bIn      []int
		expected []int
	}{
		{
			aIn:      []int{},
			bIn:      []int{},
			expected: []int{},
		},
		{
			aIn:      []int{1},
			bIn:      []int{2},
			expected: []int{1, 2},
		},
		{
			aIn:      []int{1, 5},
			bIn:      []int{2, 3},
			expected: []int{1, 2, 3, 5},
		},
		{
			aIn:      []int{},
			bIn:      []int{1, 4, 5},
			expected: []int{1, 4, 5},
		},
		{
			aIn:      []int{1, 1, 4},
			bIn:      []int{1, 1, 4, 5},
			expected: []int{1, 1, 1, 1, 4, 4, 5},
		},
	}

	for _, tb := range table {
		assert.Equal(t, tb.expected, mergeInt(tb.aIn, tb.bIn), "input; a: %+v, b: %+v", tb.aIn, tb.bIn)
	}
}
