package algo

import (
	"N-puzzle-GO/pkg/board"
	"fmt"
)

func printPath(state []board.StateOfBoard) {
	fmt.Printf("[")
	for i, v := range state {
		if i == len(state) - 1{
			fmt.Printf("%s]\n", v.EmptyTile.Move)
		} else {
			fmt.Printf("%s, ", v.EmptyTile.Move)
		}
	}
}

func getRecursivePath(node board.StateOfBoard) []board.StateOfBoard{
	path := make([]board.StateOfBoard, 0)
	for node.From != nil {
		path = append([]board.StateOfBoard{node}, path...)
		node = *node.From
	}
	return path
}