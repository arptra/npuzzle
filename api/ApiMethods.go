package api

import (
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func PutState(context *gin.Context) {
	var FirstState map[string][]int
	var arr []int
	key := "firstState"

	if err := context.ShouldBindJSON(&FirstState); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arr = FirstState[key]
	fmt.Println(arr)
	state := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		arr[len(arr)-1],
		arr[len(arr)-1],
		arr,
		nil,
		nil,
	}
	goalState := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		arr[len(arr)-1],
		arr[len(arr)-1],
		board.GetGoalState(int(math.Sqrt(float64(len(arr))))),
		nil,
		nil,
	}
	board.PrintState(state)
	//fmt.Println(state.Size, state.EmptyTile, state.PrevEmptyTilePosition, state.EmptyTilePosition, state.CurrentBoardState, state.From, state.To)
	board.PrintState(goalState)
	//fmt.Println(goalState.Size, goalState.EmptyTile, goalState.PrevEmptyTilePosition, goalState.EmptyTilePosition, goalState.CurrentBoardState, goalState.From, goalState.To)

	algo.AlgoStart(state, goalState)
}
