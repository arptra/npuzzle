package algo

import (
	"N-puzzle-GO/pkg/board"
	"fmt"
	"math"
)

func ManhattanDist(state, goal board.StateOfBoard) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var total float64
	total = 0
	for ii := range state.CurrentBoardState {
		if state.CurrentBoardState[ii] != 9 {
			n, err := goal.FindN(state.CurrentBoardState[ii])
			if err != nil {
				panic(err)
			}
			total += math.Abs(float64(ii%state.Size-n%state.Size)) +
				math.Abs(float64(ii/state.Size-n/state.Size))
		}
	}
	return total
}

func EuclideanDistance(tab []int, result []int) int {
	var dist int
	var i int
	var destIndex int
	var distRow float64
	var distCol float64

	inverseGoal := invert(result)
	size := int(math.Sqrt(float64(len(tab))))
	for i = 0; i < int(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			distRow = math.Pow(float64(i/size-destIndex/size), 2)
			distCol = math.Pow(float64(i%size-destIndex%size), 2)
			dist += int(math.Sqrt(distRow + distCol))
		}
	}
	return dist
}

func HammingDistance(tab []int, result []int) int {
	inverseGoal := invert(result)
	var i int
	var misplaced int
	var destIndex int
	for i = 0; i < len(tab); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			if destIndex != i {
				misplaced++
			}
		}
	}
	return misplaced
}

func invert(tab []int) []int {
	result := append(tab[:0:0], tab...)
	var i = 0
	for i = 0; i < len(result); i++ {
		result[tab[i]] = i
	}
	return result
}

func getHeuristic(heuristic string) int {

	return
}
