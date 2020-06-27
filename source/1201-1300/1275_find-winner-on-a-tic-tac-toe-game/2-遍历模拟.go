package main

import (
	"fmt"
)

func main() {
	//fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
	fmt.Println(tictactoe([][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}}))
}

// leetcode1275_找出井字棋的获胜者
func tictactoe(moves [][]int) string {
	arr := make([][]int, 4)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 4)
	}
	for i := 0; i < len(moves); i++ {
		x := moves[i][0] + 1
		y := moves[i][1] + 1
		if i%2 == 0 {
			arr[x][y]++
		} else {
			arr[x][y]--
		}
	}
	sum1 := arr[1][1] + arr[2][2] + arr[3][3]
	sum2 := arr[1][3] + arr[2][2] + arr[3][1]
	if sum1 == 3 || sum2 == 3 {
		return "A"
	} else if sum1 == -3 || sum2 == -3 {
		return "B"
	}
	for i := 1; i <= 3; i++ {
		sum1 := arr[i][1] + arr[i][2] + arr[i][3]
		sum2 := arr[1][i] + arr[2][i] + arr[3][i]
		if sum1 == 3 || sum2 == 3 {
			return "A"
		} else if sum1 == -3 || sum2 == -3 {
			return "B"
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
