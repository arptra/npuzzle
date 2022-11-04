package algo

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/board"
	"fmt"
)

func savePath(state []board.StateOfBoard) {
	globalvars.SuccessPath = make(map[string][]string)
	//globalvars.SuccessPath[globalvars.PathKey] = make([]string, len(state)-1)
	fmt.Printf("[")
	for i, v := range state {
		if i == len(state)-1 {
			fmt.Printf("%s]\n", v.EmptyTile.Move)
			//break
		} else {
			//globalvars.SuccessPath[globalvars.PathKey][i] = strings.Clone(v.EmptyTile.Move)
			//fmt.Printf("%s", v.EmptyTile.Move)
			globalvars.SuccessPath[globalvars.PathKey] = append(globalvars.SuccessPath[globalvars.PathKey], v.EmptyTile.Move)
			fmt.Printf("%s, ", globalvars.SuccessPath[globalvars.PathKey][i])
			//fmt.Print(" = ")
			//test[i] = strings.Clone(v.EmptyTile.Move)
			//fmt.Println(i)
		}
	}
	//
	////fmt.Printf("%i = %s, %i = %s", 0, test[0], 1, test[1])
	globalvars.ALGO_END = true
}

func printPath(state []board.StateOfBoard) {
	fmt.Printf("[")
	for i, v := range state {
		if i == len(state)-1 {
			fmt.Printf("%s]\n", v.EmptyTile.Move)
			//fmt.Println(i)
		} else {
			fmt.Printf("%s, ", v.EmptyTile.Move)
			//fmt.Println(i)
		}
	}
}

func getRecursivePath(node board.StateOfBoard) []board.StateOfBoard {
	path := make([]board.StateOfBoard, 0)
	for node.From != nil {
		path = append([]board.StateOfBoard{node}, path...)
		node = *node.From
	}
	return path
}
