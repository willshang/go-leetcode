package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	k := 4
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, n := range nums {
		if m[n] != 0 && (i+1)-m[n] <= k {
			return true
		}
		m[n] = i + 1
	}
	return false
}
