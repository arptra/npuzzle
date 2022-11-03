package main

import (
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/apiserver"
	"N-puzzle-GO/pkg/board"
	"log"
	"math"
	"os"
)

//var firstState map[string][]int

func manual() {
	manualFirstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//manualFirstState := []int{1, 2, 3, 0, 4, 6, 7, 5, 8}
	State := board.StateOfBoard{
		3,
		board.InitMove,
		4,
		4,
		manualFirstState,
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
	//firstState := make(map[string][]int)
	apiserver.ApiServerStart(&firstState)
	key := "firstState"
	State := board.StateOfBoard{
		int(math.Sqrt(float64(len(firstState[key])))),
		board.InitMove,
		firstState[key][len(firstState[key])],
		firstState[key][len(firstState)],
		firstState[key],
		nil,
		nil,
	}
	goalState := board.StateOfBoard{
		int(math.Sqrt(float64(len(firstState[key])))),
		board.InitMove,
		firstState[key][len(firstState[key])],
		firstState[key][len(firstState)],
		board.GetGoalState(int(math.Sqrt(float64(len(firstState[key]))))),
		nil,
		nil,
	}
	algo.AlgoStart(State, goalState)
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
