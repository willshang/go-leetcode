package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200}, 150))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}

// leetcode948_令牌放置
func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	res := 0
	maxValue := 0
	left, right := 0, len(tokens)-1
	for left <= right {
		if tokens[left] <= P { // 分够：消耗分数，增加能量
			P = P - tokens[left]
			left++
			maxValue++
		} else if maxValue > 0 { // 分不够：消耗能量，增加分数
			P = P + tokens[right]
			right--
			maxValue--
		} else {
			break
		}
		if maxValue > res {
			res = maxValue
		}
	}
	return res
}
