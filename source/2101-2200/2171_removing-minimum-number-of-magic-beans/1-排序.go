package main

import "sort"

func main() {

}

// leetcode2171_拿出最少数目的魔法豆
func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sum := int64(0)
	for i := 0; i < n; i++ {
		sum = sum + int64(beans[i])
	}
	res := sum
	sort.Ints(beans)
	for i := 0; i < n; i++ {
		res = min(res, sum-int64(beans[i])*int64(n-i)) // 把较大的n-i个数都变为beans[i]
	}
	return res
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
