package main

import "math"

func main() {

}

// leetcode822_翻转卡片游戏
func flipgame(fronts []int, backs []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] { // 前后相同，不能选
			m[fronts[i]] = true
		}
	}
	res := math.MaxInt32
	for i := 0; i < len(fronts); i++ {
		if m[fronts[i]] == false { // 不相同
			if fronts[i] < res {
				res = fronts[i]
			}
		}
		if m[backs[i]] == false { // 不相同
			if backs[i] < res {
				res = backs[i]
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
