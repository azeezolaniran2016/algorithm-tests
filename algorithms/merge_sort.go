package algorithms

// Merge sort is a sorting technique based on divide and conquer technique.
// With worst-case time complexity being Ο(n log n).
// Merge sort first divides the array into equal halves and then combines them in a sorted manner.

// MergeSortInt uses MergeSort Algorithm to sort an Integer list
func MergeSortInt(list []int) []int {
	// base case
	l := len(list)
	if l == 1 {
		return list
	}

	c := l / 2

	a := MergeSortInt(list[0:c])
	b := MergeSortInt(list[c:l])

	return mergeInt(a, b)
}

// mergeInt merges 2 sorted list into 1 sorted list
func mergeInt(a []int, b []int) (c []int) {
	c = make([]int, 0, len(a)+len(b))
	lenA := len(a)
	lenB := len(b)
	for lenA > 0 && lenB > 0 {
		if a[0] < b[0] {
			c = append(c, a[0])
			a = popFirstElement(a)
		} else if a[0] == b[0] {
			c = append(c, a[0], b[0])
			a = popFirstElement(a)
			b = popFirstElement(b)
		} else {
			c = append(c, b[0])
			b = popFirstElement(b)
		}
		lenA = len(a)
		lenB = len(b)
	}

	if len(a) > 0 {
		c = append(c, a...)
	}

	if len(b) > 0 {
		c = append(c, b...)
	}

	return
}

// popFirstElement removes first element in a list, and returns a list containing
// the remaining elements
func popFirstElement(list []int) []int {
	l := len(list)
	if l < 2 {
		return []int{}
	}
	return list[1:l]
}
