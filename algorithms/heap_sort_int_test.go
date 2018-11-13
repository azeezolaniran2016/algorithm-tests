package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSortInt(t *testing.T) {
	for _, tc := range sortTestCases {
		assert.Equal(t, tc.expected, HeapSortInt(tc.in))
	}
}
