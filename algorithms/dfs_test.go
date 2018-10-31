package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	assert.Equal(t, 2, 2)
	for _, tree := range trees {
		l := len(tree.Nodes)
		for x := l; x > 0; x-- {
			result := DFS(x, &tree)
			assert.Equal(t, len(tree.Nodes), len(result))
			switch x {

			// checks that the node is truly traversed using DFS
			case 6:
				{
					assert.Equal(t, result, []int{6, 1, 2, 3, 4, 5})
				}
			case 5:
				{
					assert.Equal(t, result, []int{5, 1, 2, 3, 6, 4})
				}
			case 4:
				{
					assert.Equal(t, result, []int{4, 1, 2, 3, 6, 5})
				}
			case 3:
				{
					assert.Equal(t, result, []int{3, 1, 2, 4, 5, 6})
				}
			case 2:
				{
					assert.Equal(t, result, []int{2, 1, 3, 6, 4, 5})
				}
			case 1:
				{
					assert.Equal(t, result, []int{1, 2, 3, 6, 4, 5})
				}
			}
			fmt.Printf("start: %v, result: %+v\n", x, result) // manual validation of traversal
		}
	}
}
