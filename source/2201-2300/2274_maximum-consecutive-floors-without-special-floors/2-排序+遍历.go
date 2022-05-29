package main

import "sort"

func main() {

}

// leetcode2274_不含特殊楼层的最大连续楼层数
func maxConsecutive(bottom int, top int, special []int) int {
	res := 0
	special = append(special, bottom-1, top+1) // 补齐
	sort.Ints(special)
	for i := 0; i < len(special)-1; i++ {
		res = max(res, special[i+1]-special[i]-1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
