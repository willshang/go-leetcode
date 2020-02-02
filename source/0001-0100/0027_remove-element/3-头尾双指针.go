package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for {
		// 从左向右找到等于 val 的位置
		for i < len(nums) && nums[i] != val {
			i++
		}
		// 从右向左找到不等于 val 的位置
		for j >= 0 && nums[j] == val {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}
