package main

import "fmt"

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}

// leetcode1535_找出数组游戏的赢家
func getWinner(arr []int, k int) int {
	res := arr[0]
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
			count = 1
		} else {
			count++
		}
		if count == k {
			break
		}
	}
	return res
}
