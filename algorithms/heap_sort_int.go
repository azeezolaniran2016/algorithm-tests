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
	// heapify the list into a min binary heap
	// after every heapify, the item at the root node (list[0]) == smallest item
	// append that item to our result, and re-heapify the remaining items
	r = make([]int, 0, len(list))
	for len(list) > 0 {
		list = makeMinHeap(list)
		r = append(r, list[0])
		list = list[1:]
	}
	return r
}

// makeListHeap converts an int array in to a Min Heap array
func makeMinHeap(list []int) []int {
	for i := len(list) - 1; i >= 0; i-- {
		list = heapifyMin(list, i)
	}
	return list
}

func heapifyMin(list []int, idx int) []int {
	leftChild := getLeftChildNode(idx)
	rightChild := getRightChildNode(idx)

	if rightChild < len(list) && list[idx] > list[rightChild] {
		// swap position and heapify min downwards
		list[idx], list[rightChild] = list[rightChild], list[idx]
		list = heapifyMin(list, rightChild)
	} else if leftChild < len(list) && list[idx] > list[leftChild] {
		// swap positions and heapifyMin downards
		list[idx], list[leftChild] = list[leftChild], list[idx]
		list = heapifyMin(list, leftChild)
	}
	return list
}

func getLeftChildNode(idx int) int {
	return idx*2 + 1
}

func getRightChildNode(idx int) int {
	return idx*2 + 2
}
