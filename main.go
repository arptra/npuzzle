package main

import (
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"log"
	"os"
)

func manual() {
	firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//firstState := []int{1, 2, 3, 0, 4, 6, 7, 5, 8}
	State := board.StateOfBoard{
		3,
		board.InitMove,
		4,
		4,
		firstState,
		nil,
		nil,
	}
	//test := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	goalState := board.StateOfBoard{
		3,
		board.InitMove,
		4,
		4,
		board.GetGoalState(3),
		//test,
		nil,
		nil,
	}
	board.PrintState(State)
	board.PrintState(goalState)
	//all := board.GetAllState(State)
	//for _, v := range all {
	//	board.PrintState(v)
	//}
	algo.AlgoStart(State, goalState)
}

func automatic() {
	//var firstState []int
	//apiserver.ApiServerStart(&firstState)
	fmt.Println("auto")
}

func main() {
	if len(os.Args) > 2 || (len(os.Args) == 2 && os.Args[1] != "manual") {
		log.Fatal("usage: ./npuzzle manual | ./npuzzle")
	}
	if len(os.Args) == 1 {
		automatic()
	} else {
		manual()
	}

}
