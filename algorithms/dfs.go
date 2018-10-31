package algorithms

// DFS traverses a Tree using the Breadth First Search algorithm
// It returns an integer slice containing all the nodes visited specifically for testing
// No Node should be visited more than once
// All Nodes must be visited
func DFS(init int, tree *Tree) (visitedVertices []int) {
	// mark initial vertex as visited
	visited := map[int]struct{}{
		init: struct{}{},
	}

	// add initial vertex to our visited list for validation
	visitedVertices = []int{init}

	// add initial vertex to our stack: push
	stack := []Node{tree.Nodes[init]}

	l := len(stack)
	for l > 0 {
		// grab next on stack i.e pop
		node := stack[l-1]     // grab next on stack
		stack = stack[0 : l-1] // pop stack

		visitedVertices = traverseNode(tree, &node, stack, visited, visitedVertices)
		l = len(stack)
	}
	return visitedVertices
}

// traverseNode traverses a node depth first.
func traverseNode(tree *Tree, node *Node, stack []Node, visited map[int]struct{}, visitedVertices []int) []int {
	for _, adjVertex := range node.AdjVertices {
		_, ok := visited[adjVertex]
		if ok {
			continue
		}
		n := tree.Nodes[adjVertex]
		stack = append(stack, n)
		visited[adjVertex] = struct{}{}
		visitedVertices = append(visitedVertices, adjVertex)
		visitedVertices = traverseNode(tree, &n, stack, visited, visitedVertices)
	}
	return visitedVertices
}
