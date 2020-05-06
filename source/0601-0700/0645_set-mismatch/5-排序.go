package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 4, 2}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	repeatNum := nums[0]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			repeatNum = nums[i]
		}
	}
	return []int{repeatNum, n*(n+1)/2 - sum + repeatNum}
}
