package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findBestValue([]int{4, 9, 3}, 10))
}

// leetcode1300_转变数组后最接近目标值的数组和
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	temp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		temp[i] = temp[i-1] + arr[i-1]
	}
	right := arr[n-1]
	res := 0
	diff := target
	for i := 1; i <= right; i++ {
		index := sort.SearchInts(arr, i)
		if index < 0 {
			index = abs(index) - 1
		}
		total := temp[index] + (n-index)*i
		if abs(total-target) < diff {
			diff = abs(total - target)
			res = i
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
