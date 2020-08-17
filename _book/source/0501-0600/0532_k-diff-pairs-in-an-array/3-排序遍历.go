package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(findPairs([]int{1,3,1,5,4},1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 2))
}

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	sort.Ints(nums)
	count := 0
	prev := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == prev {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] > k {
				break
			}
			if nums[i]+k == nums[j] {
				count++
				break
			}
		}
		prev = nums[i]
	}
	return count
}
