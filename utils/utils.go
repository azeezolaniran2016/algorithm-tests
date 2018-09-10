package utils

import (
	"strconv"
	"strings"
)

// DecimalToBaseX converts a number from base 10 to a specified base via recursion
func DecimalToBaseX(decimal, base int) string {
	// base case
	if decimal == 0 {
		return "0"
	}

	// get base x value
	baseXVal := DecimalToBaseX(decimal/base, base)

	// trim trailing "0" before appending
	leftTrimmed := strings.TrimLeft(baseXVal, "0")

	// append
	return leftTrimmed + strconv.Itoa(decimal%base)
}
