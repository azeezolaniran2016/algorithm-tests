package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortInt(t *testing.T) {
	for _, tc := range sortTestCases {
		assert.Equal(t, tc.expected, QuickSortInt(tc.in))
	}
}
