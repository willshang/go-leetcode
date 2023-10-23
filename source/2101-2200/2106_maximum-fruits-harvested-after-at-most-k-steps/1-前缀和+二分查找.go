package main

import (
	"sort"
)

func main() {

}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	res := 0
	n := len(fruits)
	sum := make([]int, n+1)
	position := make([]int, n)
	for i := 0; i < n; i++ {
		position[i] = fruits[i][0]
		sum[i+1] = sum[i] + fruits[i][1]
	}
	for i := 0; i <= k/2; i++ { // 尝试所有组合，先i后j
		j := k - 2*i
		res = max(res, getValue(position, sum, startPos-i, startPos+j)) // i往左边、j往右边
		res = max(res, getValue(position, sum, startPos-j, startPos+i)) // i往右边、j往左边
	}
	return res
}

func getValue(arr []int, sum []int, left, right int) int {
	l := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= left // 第一个大于等于left的下标
	})
	r := sort.Search(len(arr), func(i int) bool {
		return arr[i] > right // 第一个大于right的下标
	})
	return sum[r] - sum[l]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
