package algorithms

// InsertionSort is an in-place comparison-based sorting algorithm.
// Here, a sub-list is maintained which is always sorted.
// For example, the lower part of an array is maintained to be sorted.
// An element which is to be inserted in this sorted sub-list, has to find its appropriate place and then it has to be inserted there.
// Hence the name, insertion sort.
// The array is searched sequentially and unsorted items are moved and inserted  into the sorted
// sub-list (in the same array).
// This algorithm is not suitable for large data sets as its average and worst case complexity are of Ο(n2)

// InsertionSortInt sorts a list of integers
func InsertionSortInt(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}

	for index, value := range list {
		if index > length-2 {
			break
		}
		if value > list[index+1] {
			list[index] = list[index+1]
			list[index+1] = value

			// sort sublist
			for i := index; i > 0; i-- {
				if list[i] < list[i-1] {
					v := list[i]
					list[i] = list[i-1]
					list[i-1] = v
				}
			}
		}
	}
	return list
}
