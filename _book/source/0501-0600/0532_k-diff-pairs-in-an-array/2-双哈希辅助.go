package main

import "fmt"

func main() {
	//fmt.Println(findPairs([]int{1,3,1,5,4},1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 2))
}

// leetcode532_数组中的K-diff数对
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m := make(map[int]bool)
	res := make(map[int]bool)
	for _, value := range nums {
		if m[value-k] {
			res[value-k] = true
		}
		if m[value+k] {
			res[value] = true
		}
		m[value] = true
	}
	return len(res)
}
