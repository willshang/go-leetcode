package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
	fmt.Println(tictactoe([][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}}))
}

func tictactoe(moves [][]int) string {
	arrA := make([]int, 0)
	arrB := make([]int, 0)
	for i := 0; i < len(moves); i++ {
		value := moves[i][0]*3 + moves[i][1] + 1
		if i%2 == 0 {
			arrA = append(arrA, value)
			if check(arrA) {
				return "A"
			}
		} else {
			arrB = append(arrB, value)
			if check(arrB) {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

var state = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
	{1, 5, 9},
	{3, 5, 7},
}

func check(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				temp := []int{arr[i], arr[j], arr[k]}
				sort.Ints(temp)
				for n := 0; n < len(state); n++ {
					if temp[0] == state[n][0] &&
						temp[1] == state[n][1] &&
						temp[2] == state[n][2] {
						return true
					}
				}
			}
		}
	}
	return false
}
