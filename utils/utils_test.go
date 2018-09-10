package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalToBaseX(t *testing.T) {
	// table tests
	table := []struct {
		decimal  int
		base     int
		expected string
	}{
		{0, 2, "0"}, {1, 2, "1"}, {2, 2, "10"}, {3, 2, "11"}, {100, 2, "1100100"},
	}

	for _, test := range table {
		assert.Equal(t, test.expected, DecimalToBaseX(test.decimal, test.base))
	}
}
