package main

import (
	"fmt"

	. "github.com/azeezolaniran2016/algorithm-tests/algorithms"
)

func main() {
	tcTrees := []Tree{
		Tree{
			Nodes: map[int]Node{
				1: Node{
					Vertex:      1,
					AdjVertices: []int{2, 3, 4, 5, 6},
				},
				2: Node{
					Vertex:      2,
					AdjVertices: []int{1},
				},
				3: Node{
					Vertex:      3,
					AdjVertices: []int{1, 6},
				},
				4: Node{
					Vertex:      4,
					AdjVertices: []int{1, 5, 6},
				},
				5: Node{
					Vertex:      5,
					AdjVertices: []int{1, 4},
				},
				6: Node{
					Vertex:      6,
					AdjVertices: []int{1, 3, 4},
				},
			},
		},
	}
	result := DFS(6, &tcTrees[0])
	fmt.Printf("result: %+v", result)
}
