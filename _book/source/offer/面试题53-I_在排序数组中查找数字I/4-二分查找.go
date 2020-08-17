package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	count := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			count++
			for i := mid + 1; i < len(nums); i++ {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}
			for i := mid - 1; i >= 0; i-- {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}
			return count
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return count
}
