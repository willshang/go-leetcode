package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{0, 0}, 0))
	fmt.Println(fourSum([]int{0, 1, 5, 0, 1, 5, 5, -4}, 11))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			temp := target - nums[i] - nums[j]
			left := j + 1
			right := len(nums) - 1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < len(nums)-2 && nums[right] == nums[right+1] {
					right--
					continue
				}
				if nums[left]+nums[right] > temp {
					right--
				} else if nums[left]+nums[right] < temp {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				}
			}
		}
	}
	return res
}
