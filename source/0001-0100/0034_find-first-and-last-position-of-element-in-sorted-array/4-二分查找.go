package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1}, 0))
}

// leetcode34_在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left := -1
	right := -1
	for i, j := 0, len(nums)-1; i <= j; {
		mid := i + (j-i)/2
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			for temp := mid; temp >= 0; temp-- {
				if target == nums[temp] {
					left = temp
				} else {
					break
				}
			}
			for temp := mid; temp < len(nums); temp++ {
				if target == nums[temp] {
					right = temp
				} else {
					break
				}
			}
			break
		}
	}
	return []int{left, right}
}
