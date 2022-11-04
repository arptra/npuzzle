package api

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func GetPath(context *gin.Context) {
	if globalvars.ALGO_END == true {
		context.JSON(200, globalvars.SuccessPath)
	} else {
		context.AbortWithStatus(202)
	}
}

func PutState(context *gin.Context) {
	var arr []int

	if err := context.ShouldBindJSON(&globalvars.InputState); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	globalvars.ALGO_END = false // if previous GET request not 200

	arr = globalvars.InputState[globalvars.InputKey]
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
