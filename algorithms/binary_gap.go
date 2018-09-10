package algorithms

import (
	"fmt"

	"github.com/azeezolaniran2016/algorithm-tests/utils"
)

/*
	A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.
	For example, number 9 has binary representation 1001 and contains a binary gap of length 2.
	The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3.
	The number 20 has binary representation 10100 and contains one binary gap of length 1.
	The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

	Write a function:
	func Solution(N int) int
	that, given a positive integer N, returns the length of its longest binary gap.
	The function should return 0 if N doesn't contain a binary gap.

	For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.
	Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

	Write an efficient algorithm for the following assumptions:
	-	N is an integer within the range [1..2,147,483,647].
*/

// BinaryGap calculates binary gap following the requirements in the comments above
// ToDo: I got too lazy to comment what I did here, but should do it later ;)
func BinaryGap(n int) int {
	binaryGap := 0

	binary := utils.DecimalToBaseX(n, 2)
	tempCount := 0

	for idx := range binary {
		if idx+1 < len(binary) {
			if fmt.Sprintf("%c", binary[idx]) == "0" && fmt.Sprintf("%c", binary[idx+1]) == "1" {
				if tempCount > binaryGap {
					binaryGap = tempCount
					tempCount = 0
				}
			}
			if fmt.Sprintf("%c", binary[idx]) == "1" && fmt.Sprintf("%c", binary[idx+1]) == "0" {
				tempCount = 1
			}
		}
		if tempCount > 0 && fmt.Sprintf("%c", binary[idx]) == "0" {
			tempCount++
		}

		if binaryGap > len(binary)/2 {
			break
		}
	}

	return binaryGap
}
