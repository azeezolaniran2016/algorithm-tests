package algorithms

// Node defines a node which contains a Vert
type Node struct {
	Vertex      int
	AdjVertices []int
}

// Tree defines a Tree containing all it's Nodes
type Tree struct {
	Nodes map[int]Node
}

// BFS traverses a Tree using the Breadth First Search algorithm
// It returns an integer slice containing all the nodes visited specifically for testing
// No Node should be visited more than once
// All Nodes must be visited
func BFS(init int, tree *Tree) (visitedVertices []int) {
	// mark initial vertex as visited
	visited := map[int]struct{}{
		init: struct{}{},
	}

	// add initial vertex to our visited list for validation
	visitedVertices = []int{init}

	// add initial vertex to our queue
	queue := []Node{tree.Nodes[init]}

	l := len(queue)
	for l > 0 {
		// grab next on queue i.e dequeue
		node := queue[l-1]  // grab next on queue
		queue = queue[:l-1] // dequeue

		for _, v := range node.AdjVertices {
			_, ok := visited[v]
			if !ok {
				// append vertex to queue
				queue = append(queue, tree.Nodes[v])
				// mark it as visited
				visited[v] = struct{}{}
				// add it to our visited list for validation
				visitedVertices = append(visitedVertices, v)
			}
		}
		// evaluate length of our queue
		l = len(queue)
	}
	return
}
