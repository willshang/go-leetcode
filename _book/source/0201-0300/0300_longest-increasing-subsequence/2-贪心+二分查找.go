package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	arr := make([]int, len(nums)+1)
	arr[1] = nums[0]
	res := 1
	for i := 1; i < len(nums); i++ {
		if arr[res] < nums[i] {
			res++
			arr[res] = nums[i]
		} else {
			left, right := 1, res
			index := 0
			for left <= right {
				mid := left + (right-left)/2
				if arr[mid] < nums[i] {
					index = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			arr[index+1] = nums[i]
		}
	}
	return res
}
