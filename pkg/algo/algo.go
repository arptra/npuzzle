package algo

import (
	"N-puzzle-GO/pkg/board"
	"fmt"
	"math"
)


var FOUND = false


func AlgoStart(start, goal  board.StateOfBoard) int {
	fmt.Println("Algo start")
	threshold := ManhattanDist(start, goal)
	for {
		fmt.Println(threshold)
		temp := Search(start, goal, 0, threshold) //function search(node,g score,threshold)
		if FOUND {                //if goal FOUND
			return 1
		}
		if temp == math.MaxInt32 {         //Threshold larger than maximum possible f value
			return 0
		}							       //or set Time limit exceeded
		threshold = temp
	}
}

func Search(node, goal board.StateOfBoard, g, threshold float64) float64 {
	f := g + ManhattanDist(node, goal)
	if f > threshold {
		 return f
	}
	if board.BoardEquals(node, goal) {
		FOUND = true
		path := getRecursivePath(node)
		printPath(path)
		return f
	}
	min := math.MaxFloat64
	for _, newNode := range board.GetAllState(node) {
		node.To = &newNode
		newNode.From = &node
		temp := Search(newNode, goal, g+1, threshold)
		if board.BoardEquals(newNode, goal) {
			return f
		}
		if temp < min {
			min = temp
		}
	}
	return min
}
