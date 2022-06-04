package main

import "sort"

func main() {

}

// leetcode2285_道路的最大总重要性
func maximumImportance(n int, roads [][]int) int64 {
	arr := make([]int, n)
	for i := 0; i < len(roads); i++ {
		a, b := roads[i][0], roads[i][1]
		arr[a]++
		arr[b]++
	}
	sort.Ints(arr)
	res := int64(0)
	for i := 1; i <= n; i++ {
		res = res + int64(arr[i-1])*int64(i)
	}
	return res
}
