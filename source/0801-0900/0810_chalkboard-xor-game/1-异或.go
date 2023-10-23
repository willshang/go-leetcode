package main

import "fmt"

func main() {
	fmt.Println(1 ^ 2 ^ 3)
}

// leetcode810_黑板异或游戏
func xorGame(nums []int) bool {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		res = res ^ nums[i]
	}
	return res == 0 || n%2 == 0
}
