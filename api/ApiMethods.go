package api

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
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
	var emptyTileIndex int

	if err := context.ShouldBindJSON(&globalvars.InputState); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	globalvars.ALGO_END = false // if previous GET request not 200

	arr = globalvars.InputState[globalvars.InputStateKey]
	emptyTile := globalvars.InputState[globalvars.EmptyTileKey][0]
	for i := range arr {
		if emptyTile == arr[i] {
			emptyTileIndex = i
		}
	}

	state := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		emptyTileIndex,
		emptyTileIndex,
		arr,
		nil,
		nil,
	}
	goalState := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		emptyTileIndex,
		emptyTileIndex,
		board.GetGoalState(int(math.Sqrt(float64(len(arr))))),
		nil,
		nil,
	}
	//board.PrintState(state)
	//board.PrintState(goalState)

	algo.AlgoStart(state, goalState)
}
