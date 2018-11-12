package algorithms

/*
Quick sort works by choosing an element as pivot and partitions the given array around the picked pivot.
The role of the pivot value is to assist with splitting the list.
The actual position where the pivot value belongs in the final sorted list will be used to divide the list
for subsequent calls to the quick sort.
*/

// QuickSortInt uses QuickSort algorithm to sort an integrer list
func QuickSortInt(list []int) []int {
	l := len(list)

	// base case. Return
	if l < 2 {
		return list
	}

	left, right := 0, l-1

	// partitioning:
	// All element lesser > pivot > All elements greater
	// assume element at right is the pivot.
	for i := range list {
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}
	// move pivot element to it's appropriate postion
	list[left], list[right] = list[right], list[left]

	// repeat recursively for right side and left side of partitioned array. Pivot remains in it's original position
	QuickSortInt(list[:left])
	QuickSortInt(list[left+1:])

	return list
}
