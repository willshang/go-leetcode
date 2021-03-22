package main

import "sort"

func main() {

}

// leetcode1402_做菜顺序
func maxSatisfaction(satisfaction []int) int {
	res := 0
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	sum := 0
	for i := 0; i < len(satisfaction); i++ {
		if sum+satisfaction[i] <= 0 {
			break
		}
		sum = sum + satisfaction[i]
		res = res + sum // 每多遍历1次，前后相邻的2个数，较大数相对较小数多+1次
	}
	return res
}
