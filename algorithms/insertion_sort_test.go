package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortInt(t *testing.T) {
	for _, tc := range sortTestCases {
		assert.Equal(t, tc.expected, InsertionSortInt(tc.in))
	}
}
