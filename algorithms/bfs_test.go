package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	trees = []Tree{
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
)

func TestBFS(t *testing.T) {
	for _, tree := range trees {
		l := len(tree.Nodes)
		for x := l; x > 0; x-- {
			result := BFS(x, &tree)
			assert.Equal(t, len(tree.Nodes), len(result))
			fmt.Printf("start: %v, result: %+v\n", x, result) // ToDo: assert that the result shows that the Tree was traveresed in the right order. Right now I do this manually
		}
	}
}
