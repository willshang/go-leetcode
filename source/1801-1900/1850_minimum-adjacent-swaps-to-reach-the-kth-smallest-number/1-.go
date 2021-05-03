package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getMinSwaps("5489355142", 4))
}

func getMinSwaps(num string, k int) int {
	arr := []byte(num)
	n := len(arr)
	res := 0
	for i := 1; i <= k; i++ {
		fmt.Println(string(arr))
		nextPermutation(arr)
	}
	fmt.Println(string(arr), n)
	return res
}

func nextPermutation(nums []byte) {
	left := len(nums) - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left == -1 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}
	right := len(nums) - 1
	for right >= 0 && nums[right] <= nums[left] {
		right--
	}
	nums[left], nums[right] = nums[right], nums[left]
	// sort.Ints(nums[left+1:])
	count := 0
	for i := left + 1; i < (left+1+len(nums)-1)/2; i++ {
		nums[i], nums[len(nums)-1-count] = nums[len(nums)-1-count], nums[i]
		count++
	}
}
