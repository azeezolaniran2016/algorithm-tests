package algorithms

/*
	Heap sort is a comparison based sorting technique based on Binary Heap data structure.
	It is similar to selection sort where we first find the maximum element and place the maximum element at the end.
	We repeat the same process for remaining element.
	What is Binary Heap?
	Let us first define a Complete Binary Tree. A complete binary tree is a binary tree in which every level,
	except possibly the last, is completely filled, and all nodes are as far left as possible
	A Binary Heap is a Complete Binary Tree where items are stored in a special order such that value in a parent node is greater(or smaller) than the values in its two children nodes.
	The former is called as max heap and the latter is called min heap. The heap can be represented by binary tree or array.
												a
						b													c
			d						e							f							g
	h				i		j				k			l				m			n				o

	The algorithm presented below does the following:
	1. Takes an input array, and rearranges it to satisfy the Min Heap rules (Heapfiy). Heapification goes from Bottom -> Top and Right most -> Left most nodes
	2. Picks the item at the root node (index 0 of the heapified array) and appends it into a new sorted array.
	3. Repeats step 1 and 2 only this time the input array is less by 1 element (element picked at step 2).
*/

// HeapSortInt uses HeapSort Algorithm to sort a slice (list) of integers
func HeapSortInt(list []int) (r []int) {

	// heapify the list into a max heap

	// create a max heap from the input list
	// we start from the right most and deepest node and move upwards to build the max heap
	l := len(list)
	for i := l / 2; i >= 0; i-- {
		heapifyMax(list, i, l)
	}

	// start popping the max element, swapping it with the right most and deepest node and re-heapifying the resulting sub array
	for i := l - 1; i > 0; i-- {
		// pop and swap, but we still use same list :)
		list[0], list[i] = list[i], list[0]

		// heapify resulting sub-array
		heapifyMax(list, 0, i)
	}

	return list
}

func heapifyMax(list []int, root, size int) {
	largest := root
	left := root*2 + 1
	right := root*2 + 2

	// compare left child to root node, and update our largest node
	if left < size && list[left] > list[largest] {
		largest = left
	}

	// compare right child to root node, and update our largest node
	if right < size && list[right] > list[largest] {
		largest = right
	}

	// if the root node isn't the largest node, swap root with larger child, and heapify the sub-tree downward
	if largest != root {
		list[root], list[largest] = list[largest], list[root]
		heapifyMax(list, largest, size)
	}
}
