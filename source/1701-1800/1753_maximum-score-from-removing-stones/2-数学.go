package main

import "sort"

func main() {

}

// leetcode1753_移除石子的最大得分
func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	if arr[0]+arr[1] <= arr[2] {
		return arr[0] + arr[1]
	}
	return (arr[0] + arr[1] + arr[2]) / 2
}
