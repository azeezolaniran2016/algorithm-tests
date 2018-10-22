package algorithms

import (
	"math/rand"
)

/*
Quick sort works by choosing an element as pivot and partitions the given array around the picked pivot.
The role of the pivot value is to assist with splitting the list.
The actual position where the pivot value belongs in the final sorted list will be used to divide the list
for subsequent calls to the quick sort.
*/

// QuickSortInt uses QuickSort algorithm to sort an integrer list
func QuickSortInt(list []int) []int {
	// base case. Return
	if len(list) < 2 {
		return list
	}

	left, right := 0, len(list)-1

	// pick a random index as the pivot
	pivot := rand.Int() % len(list)

	// swap pivot with right most element
	list[pivot], list[right] = list[right], list[pivot]

	// partitioning: insert the pivot in it's appropriate position - All element lesser > pivot > All elements greater
	for i := range list {
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}
	list[left], list[right] = list[right], list[left]

	// repeat recursively for right side and left side of partitioned array. Pivot remains in it's original position
	QuickSortInt(list[:left])
	QuickSortInt(list[left+1:])

	return list
}
