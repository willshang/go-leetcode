package main

import "fmt"

func main() {
	//fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("xxxxx"))
}

// leetcode830_较大分组的位置
func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0, len(S)/3)
	left := 0
	for right := 0; right < len(S); right++ {
		if right == len(S)-1 || S[right] != S[right+1] {
			if right-left+1 >= 3 {
				res = append(res, []int{left, right})
			}
			left = right + 1
		}
	}
	return res
}
