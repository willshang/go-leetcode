package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(arr))
}

func findShortestSubArray(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	first := make(map[int]int)
	count := make(map[int]int)
	maxCount := 1
	minLen := size
	for i, n := range nums {
		count[n]++
		if count[n] == 1 {
			first[n] = i
		} else {
			length := i - first[n] + 1
			if maxCount < count[n] ||
				(maxCount == count[n] && minLen > length) {
				maxCount = count[n]
				minLen = length
			}
		}
	}
	if len(count) == size {
		return 1
	}
	return minLen
}
