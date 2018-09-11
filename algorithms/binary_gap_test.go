package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	// ToDo: Add other test cases!
	table := []struct {
		in, expected int
	}{{1041, 5}, {32, 0}}

	for _, test := range table {
		assert.Equal(t, test.expected, BinaryGap(test.in))
	}
}
