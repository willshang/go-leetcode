package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(arr))
}
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := 0, -1
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++{
		if max <= nums[i]{
			max = nums[i]
		}else {
			right = i
		}

		j := n - i - 1
		if min >= nums[j]{
			min = nums[j]
		}else {
			left = j
		}
	}

	return right - left + 1
}

/*func findUnsortedSubarray(nums []int) int {
	temp := make([]int,len(nums))
	copy(temp,nums)
	sort.Ints(temp)

	i, j := 0, len(nums)-1
	for i < len(nums) && nums[i] == temp[i]{
		i++
	}
	for i+1 < j && nums[j] == temp[j]{
		j--
	}

	return j-i+1
}*/